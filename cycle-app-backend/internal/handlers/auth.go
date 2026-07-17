package handlers

import (
	"encoding/json"
	//"log"
	"net/http"

	"cycle-app-backend/internal/models"
)

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request){
		//log.Println("Request Method:", r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.RegisterRequest 


	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	message, err := h.AuthService.Register(request)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// response := map[string]string{
	// 	"message": "User registered successfully",
	// }

	response := map[string]string{
	"message": message,
}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
