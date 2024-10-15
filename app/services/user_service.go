package services

import (
	"Backend-go/app/models"
	"Backend-go/app/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

//las funciones en go, primero van los valores que recibira, luego el prototipo de funcion y despues el valor de retorno.
func (s *UserService) isUserDuplicated(usuario string) (bool,error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		// importa el orden de (bool, error)
		return false, err
	}

	for _, u := range users {
		if u.Usuario == usuario {
			return true, nil 
		}
	}

	return false, nil
}

//por que todos son string, no se les asigna un tipo de dato independiente
func (s *UserService) isNameDuplicated(nombre, apaterno, amaterno string) (bool,error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		// importa el orden de (bool, error)
		return false, err
	}

	for _, u := range users {
		if u.Nombre == nombre && u.Apaterno == apaterno && u.Amaterno == amaterno {
			return true, nil 
		}
	}
	
	return false, nil
}

func(s *UserService) CreateUser(user models.Usuario) error {
	isDuplicated, err != s.Repo.GetUserByUsername()(user.Usuario)
	if err != nil {
		return err
	}
	if isDuplicated != nil {
		return errors.New("el usuario ya existe")
	}
	isNameDuplicated, err := s.isNameDuplicated(user.Nombre, user.Apaterno, user.Amaterno)
	if err != nil {
		return err
	}
	if isNameDuplicated {
		return errors.New("el nombre completo ya existe")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id string, user models.Usuario) error {
	userID, err := s.Repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if userID == nil {
		return errors.New("usuario no encontrado")
	}
	userID.Nombre = user.Nombre
	userID.Apaterno = user.Apaterno
	userID.Amaterno = user.Amaterno
	userID.Direccion = user.Direccion
	userID.Telefono = user.Telefono
	userID.Ciudad = user.Ciudad
	userID.Estado = user.Estado
	userID.Rol = user.Rol
	userID.Imagen = user.Imagen

	if user.Password != ""  {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
		if err != nil {
			return nil
		}
		userID.password = string(hashedPassword)	
	}
	return s.Repo.UpdateUser(id, *userID)
}

func (s *UserService) DeleteUser(id string) error {
	userID, err := s.Repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if userID == nil {
		return errors.New("usuario no encontrado")
	}
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetUserByID(id string) (*models.Usuario, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.Usuario, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) GetUserByRol(rol string) ([]models.Usuario, error) {
	return s.Repo.GetUserByRol(rol)
}