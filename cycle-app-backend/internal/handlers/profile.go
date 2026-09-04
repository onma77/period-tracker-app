package handlers

import (
	"encoding/json"
	"net/http"

	"cycle-app-backend/internal/middleware"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.UserContextKey()).(map[string]interface{})

	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"message": "Profile accessed successfully",
		"user_id": user["user_id"],
		"email":   user["email"],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
