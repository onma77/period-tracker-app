package routes

import "net/http"

func RegisterRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)

	mux.HandleFunc("/register", handlers.RegisterUser)
	return mux
}