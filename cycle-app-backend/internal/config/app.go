package config

import (
	"cycle-app-backend/internal/repository"
	"cycle-app-backend/internal/services"
)

type App struct {
	AuthService    *services.AuthService
	PeriodService  *services.PeriodService
	CycleService   *services.CycleService
	SymptomService *services.SymptomService
}

func NewApp() *App {
	userRepo := repository.NewUserRepository()
	authService := services.NewAuthService(userRepo)

	periodRepo := repository.NewPeriodRepository()
	periodService := services.NewPeriodService(periodRepo)

	cycleService := services.NewCycleService()

	symptomRepo := repository.NewSymptomRepository()
	symptomService := services.NewSymptomService(symptomRepo)
	return &App{
		AuthService:    authService,
		PeriodService:  periodService,
		CycleService:   cycleService,
		SymptomService: symptomService,
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
