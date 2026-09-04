package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/middleware"
	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/services"
)

type PeriodHandler struct {
	PeriodService *services.PeriodService
}

func NewPeriodHandler(periodService *services.PeriodService) *PeriodHandler {
	return &PeriodHandler{
		PeriodService: periodService,
	}
}

func (h *PeriodHandler) CreatePeriod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value(middleware.UserContextKey()).(map[string]interface{})

	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	userID, ok := user["user_id"].(int)

	if !ok {
		http.Error(w, "Invalid user information", http.StatusUnauthorized)
		return
	}

	var request models.CreatePeriodRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	period, err := h.PeriodService.CreatePeriod(userID, request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(period)
}

func (h *PeriodHandler) GetPeriods(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	user, ok := r.Context().Value(middleware.UserContextKey()).(map[string]interface{})

	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	userID, ok := user["user_id"].(int)

	if !ok {
		http.Error(w, "Invalid user information", http.StatusUnauthorized)
		return
	}

	periods := h.PeriodService.GetUserPeriods(userID)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(periods)
}
