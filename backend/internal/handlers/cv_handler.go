package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"digital-cv-backend/internal/services"
)

type CVHandler struct {
	service *services.CVService
}

func NewCVHandler(service *services.CVService) *CVHandler {
	return &CVHandler{service: service}
}

func (h *CVHandler) GetCV(w http.ResponseWriter, r *http.Request) {
	cv, err := h.service.GetCV()
	if err != nil {
		log.Println("GET CV ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cv)
}

