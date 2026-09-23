package routes

import (
	"net/http"

	"cycle-app-backend/internal/config"
	"cycle-app-backend/internal/handlers"
	"cycle-app-backend/internal/middleware"
)

func RegisterRoutes(app *config.App) *http.ServeMux {
	mux := http.NewServeMux()

	authHandler := handlers.NewAuthHandler(app.AuthService)

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	mux.HandleFunc("/register", authHandler.RegisterUser)
	mux.HandleFunc("/login", authHandler.LoginUser)

	mux.Handle(
		"/profile",
		middleware.AuthMiddleware(
			http.HandlerFunc(handlers.ProfileHandler),
		),
	)

	periodHandler := handlers.NewPeriodHandler(app.PeriodService)

	// periodHandler := handlers.NewPeriodHandler(app.PeriodService)

	cycleHandler := handlers.NewCycleHandler(
		app.PeriodService,
		app.CycleService,
	)

	mux.Handle(
		"/periods",
		middleware.AuthMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					periodHandler.CreatePeriod(w, r)
					return
				}

				if r.Method == http.MethodGet {
					periodHandler.GetPeriods(w, r)
					return
				}

				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}),
		),
	)

	mux.Handle(
		"/cycle",
		middleware.AuthMiddleware(
			http.HandlerFunc(cycleHandler.GetCycle),
		),
	)

	mux.Handle("/prediction", middleware.AuthMiddleware(http.HandlerFunc(cycleHandler.GetPrediction)))
	mux.Handle("/ovulation", middleware.AuthMiddleware(http.HandlerFunc(cycleHandler.GetOvulationPrediction)))

	return mux
}

// func RegisterRoutes(app *config.App) *http.ServeMux {

// 	mux := http.NewServeMux()

// 	authHandler := handlers.NewAuthHandler(app.AuthService)

// 	mux.HandleFunc("/", handlers.HomeHandler)
// 	mux.HandleFunc("/health", handlers.HealthHandler)

// 	mux.HandleFunc("/register", authHandler.RegisterUser)
// 	mux.HandleFunc("/login", authHandler.LoginUser)

// 	mux.Handle(
// 		"/profile",
// 		middleware.AuthMiddleware(
// 			http.HandlerFunc(handlers.ProfileHandler),
// 		),
// 	)

// 	return mux
// }

// periodHandler := handlers.NewPeriodHandler(app.PeriodService)

// mux.Handle(
// 	"/periods",
// 	middleware.AuthMiddleware(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			if r.Method == http.MethodPost {
// 				periodHandler.CreatePeriod(w, r)
// 				return
// 			}

// 			if r.Method == http.MethodGet {
// 				periodHandler.GetPeriods(w, r)
// 				return
// 			}

// 			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		}),
// 	),
// )

// package routes

// import (
// 	"net/http"

// 	"cycle-app-backend/internal/config"
// 	"cycle-app-backend/internal/handlers"
// )

// func RegisterRoutes(app *config.App) *http.ServeMux {

// 	mux := http.NewServeMux()

// 	authHandler := handlers.NewAuthHandler(app.AuthService)

// 	mux.HandleFunc("/", handlers.HomeHandler)
// 	mux.HandleFunc("/health", handlers.HealthHandler)
// 	mux.HandleFunc("/register", authHandler.RegisterUser)
// 	mux.HandleFunc("/login", authHandler.LoginUser)

// 	return mux
// }
