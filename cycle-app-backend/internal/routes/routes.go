package routes

import (
	"net/http"

	"cycle-app-backend/internal/handlers"
)

func RegisterRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/register", handlers.RegisterUser)

	return mux
}