package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/services"
)

type HealthGuidanceHandler struct {
	HealthGuidanceService *services.HealthGuidanceService
}

func NewHealthGuidanceHandler(
	service *services.HealthGuidanceService,
) *HealthGuidanceHandler {
	return &HealthGuidanceHandler{
		HealthGuidanceService: service,
	}
}

func (h *HealthGuidanceHandler) GetGuidance(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	symptom := r.URL.Query().Get("symptom")
	if symptom == "" {
		http.Error(w, "Please provide a symptom", http.StatusBadRequest)
		return
	}

	guidance := h.HealthGuidanceService.GetGuidance(symptom)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guidance)
}
