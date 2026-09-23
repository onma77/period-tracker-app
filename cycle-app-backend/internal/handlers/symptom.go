package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/middleware"
	"cycle-app-backend/internal/models"
	"cycle-app-backend/internal/services"
)

type SymptomHandler struct {
	SymptomService *services.SymptomService
}

func NewSymptomHandler(symptomService *services.SymptomService) *SymptomHandler {
	return &SymptomHandler{
		SymptomService: symptomService,
	}
}

func (h *SymptomHandler) CreateSymptom(w http.ResponseWriter, r *http.Request) {
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

	var request models.CreateSymptomRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	symptom, err := h.SymptomService.CreateSymptom(userID, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(symptom)
}

func (h *SymptomHandler) GetSymptoms(w http.ResponseWriter, r *http.Request) {
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

	symptoms := h.SymptomService.GetUserSymptoms(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(symptoms)
}
