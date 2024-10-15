package interfaces

import "Backend-go/app/models"

type IUser interface {
	CreateUser(user models.Usuario) error
	GetAllUsers() ([]models.Usuario, error)
	GetUserByID(id string) (*models.Usuario, error)
	GetUserByUsername(username string) (*models.Usuario, error)
	GetUserByRol(rol string) ([]models.Usuario, error)
	UpdateUser(id string, user models.Usuario) error
	DeleteUser(id string) error
}
