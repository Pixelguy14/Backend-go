package controller

import (
	"net/http"
	"backend-go/app/models"
	"backend-go/app/services"
	"encoding/json"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.Usuario
	json.NewDecoder(r.Body).Decode(&user)
	err := c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}