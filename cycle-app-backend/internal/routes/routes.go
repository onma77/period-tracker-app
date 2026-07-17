package routes

import (
	"net/http"

	"cycle-app-backend/internal/config"
	"cycle-app-backend/internal/handlers"
)

func RegisterRoutes(app *config.App) *http.ServeMux {

	mux := http.NewServeMux()

	authHandler := handlers.NewAuthHandler(app.AuthService)

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/register", authHandler.RegisterUser)

	return mux
}