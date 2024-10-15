package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *Firebase.App

func InitializeFirebaseApp() {
	opt := option.WithCredentialsFile("./firebaseServiceAccount.json")
	app, err := firebase.NewApp(context.Background(), nill, opt) //si funciona se guarda en app, si no, se guarda en err
	if err != nil {
		log.Fatalf("Error al inicializar Firebase app: %v", err)
	}
	FirebaseApp = app
	//return app
}

func GetAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error al obtener el cliente: %v", err)
	}
	return client
}
