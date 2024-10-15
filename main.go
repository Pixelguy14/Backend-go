package main

import (
	"backend-go/app/routes"
	"backend-go/config"
	"log"
	"net/http"
	"github/gorilla/mux"
)

func main () {
	config.InitializeFirebaseApp()
	router := mux.NewRouter()
	routes.InitializeRoutes(router)
	log.Println("Servidor trabajando en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}

var FirebaseApp *Firebase.App