package main

import (
	"log"
	"net/http"

	"cycle-app-backend/internal/config"
	"cycle-app-backend/internal/routes"
)

func main() {

	app := config.NewApp()

	router := routes.RegisterRoutes(app)

	log.Println("🩸 ONMA_CYCLE API started on http://localhost:8080")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}



// package main

// import (
// 	"log"
// 	"net/http"

// 	"cycle-app-backend/internal/routes"
// )

// func main() {

// 	mux := routes.RegisterRoutes()

// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}

// 	log.Println(" ONMA_CYCLE🩸 API started on http://localhost:8080")

// 	err := server.ListenAndServe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }