package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/models"
)

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	message, token, err := h.AuthService.Login(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"message": message,
		"token":   token,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// package handlers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"cycle-app-backend/internal/models"
// )

// func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var request models.LoginRequest

// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	message, err := h.AuthService.Login(request)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	response := map[string]string{
// 		"message": message,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }
