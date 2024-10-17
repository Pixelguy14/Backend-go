package routes

import (
	"Backend-go/app/controllers"
	"Backend-go/app/repositories"
	"Backend-go/app/services"
	"net/http"
)

func InitializeRoutes(router *mux.router) {
	repo := &repositories.UserRepository{}
	service := &services.UserService{Repo: repo}
	controller := &controllers.UserController{Service: service}

	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users{id}", controller.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users/rol/{rol}", controller.GetUserByRol).Methods(http.MethodGet)
	router.HandleFunc("/users{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users{id}", controller.DeleteUser).Methods(http.MethodDelete)
}
