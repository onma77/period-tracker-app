package main

import (
	"log"
	"net/http"

	"cycle-app-backend/internal/routes"
)

func main() {

	mux := routes.RegisterRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println(" ONMA_CYCLE🩸 API started on http://localhost:8080")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}