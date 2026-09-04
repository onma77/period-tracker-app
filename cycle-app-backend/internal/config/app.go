package config

import (
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/services"
)

type App struct {
	AuthService   *services.AuthService
	PeriodService *services.PeriodService
	CycleService  *services.CycleService
}

func NewApp() *App {
	userRepo := repository.NewUserRepository()
	authService := services.NewAuthService(userRepo)

	periodRepo := repository.NewPeriodRepository()
	periodService := services.NewPeriodService(periodRepo)

	cycleService := services.NewCycleService()

	return &App{
		AuthService:   authService,
		PeriodService: periodService,
		CycleService:  cycleService,
	}
}

// package config

// import (
// 	"cycle-app-backend/internal/repository"
// 	"cycle-app-backend/internal/services"
// )

// type App struct {
// 	AuthService *services.AuthService
// }

// func NewApp() *App {

// 	userRepo := repository.NewUserRepository()

// 	authService := services.NewAuthService(userRepo)

// 	return &App{
// 		AuthService: authService,
// 	}
// }
