package repositories

import (
	"Backend-go/app/models"
	"Backend-go/config"
	"context"

	"google.golang.org/api/iterator"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user models.Usuario) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	//defer es como el await
	defer client.Close()
	_, _, err = client.Collection("usuarios_lenguajes").Add(ctx, user)
	return err
}

func (r *UserRepository) GetAllUsers() ([]models.Usuario, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.Usuario
	iter := client.Collection("usuarios_lenguajes").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user models.Usuario
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.Usuario, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	doc, err := client.Collection
	("usuarios_lenguajes").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user models.Usuario
	doc.DataTo(&user)
	user.ID = doc.Ref.ID
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.Usuario, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	doc, err := client.Collection("usuarios_lenguajes").Where("usuario", "==", username).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var user models.Usuario
	doc.DataTo(&user)
	user.ID = doc.Ref.ID
	return &user, nil
}

func (r *UserRepository) GetUserByRol(rol string) ([]models.Usuario, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.Usuario
	iter := client.Collection("usuarios_lenguajes").Documents(ctx).Where("rol", "==", rol).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user models.Usuario
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(id string, user models.Usuario) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("usuarios_lenguajes").Doc(id).Set(ctx, user)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	_, err = client.Collection("usuarios_lenguajes").Doc(id).Delete(ctx)
	return err
}
