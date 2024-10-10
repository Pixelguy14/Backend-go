package services

import (
	"backend-go/app/models"
	"backend-go/app/repositories"
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
	isDuplicated, err != s.isUserDuplicated(user.Usuario)
	if err != nil {
		return err
	}
	if isDuplicated {
		return errors.New("El usuario ya existe")
	}
	isNameDuplicated, err := s.isNameDuplicated(user.Nombre, user.Apaterno, user.Amaterno)
	if err != nil {
		return err
	}
	if isNameDuplicated {
		return errors.New("El nombre ya existe")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Imagen = "default.png"
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
