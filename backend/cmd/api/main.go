package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"digital-cv-backend/internal/db"
	"digital-cv-backend/internal/handlers"
	"digital-cv-backend/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// config
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// db
	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// wiring
	repo := db.NewPostgresCVRepository(pool)
	cvService := services.NewCVService(repo)
	cvHandler := handlers.NewCVHandler(cvService)

	// router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://localhost:3000"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
    AllowCredentials: false,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/cv", cvHandler.GetCV)
	})

	// server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// graceful shutdown
	go func() {
		log.Printf("API running on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
