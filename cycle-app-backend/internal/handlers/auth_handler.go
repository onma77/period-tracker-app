package handlers

import "cycle-app-backend/internal/services"

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(auth *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: auth,
	}
}