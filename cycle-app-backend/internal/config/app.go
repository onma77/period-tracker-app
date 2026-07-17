package config

import (
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/services"
)

type App struct {
	AuthService *services.AuthService
}

func NewApp() *App {

	userRepo := repository.NewUserRepository()

	authService := services.NewAuthService(userRepo)

	return &App{
		AuthService: authService,
	}
}