package main
import (
	"log"
	"net/http"

        "digital-cv-backend/internal/db"
		"digital-cv-backend/internal/handlers"
)
func main() {
	pool, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/api/cv", handlers.GetCV(pool))

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

