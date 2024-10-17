package main

import (
	"Backend-go/app/routes"
	"Backend-go/config"

	// "Backend-go/app/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitializeFirebaseApp()
	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	// routes.Use(middleware.AuthMiddleware)
	log.Println("Servidor trabajando en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
