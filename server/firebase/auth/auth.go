package auth

import (
	"cadavre-exquis/context"
	"cadavre-exquis/firebase"
	"log"

	"firebase.google.com/go/auth"
)

var client = initAuth()

type Auth struct {
	UID      string
	UserName string
	Email    string
}

// var uid = "7YzaAOa0m2XYrPTi9pNCqe5hniY2"

func initAuth() *auth.Client {
	client, err := firebase.App.Auth(context.Context)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetAuthByUID(uid string) (*Auth, error) {
	userRecord, err := client.GetUser(context.Context, uid)
	if err != nil {
		return nil, err
	}

	auth := &Auth{}
	auth.UID = userRecord.UID

	log.Printf("Successfully fetched user data: %v\n", auth.UID)

	return auth, nil
}

func validateToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(context.Context, idToken)
}

func CreateUser(user_name string, email string, password string) (*Auth, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName(user_name).
		Disabled(false)
		// PhotoURL("http://www.example.com/12345678/photo.png").
	userRecord, err := client.CreateUser(context.Context, params)
	if err != nil {
		return nil, err
	}

	auth := &Auth{}
	auth.UID = userRecord.UID
	auth.UserName = userRecord.DisplayName
	auth.Email = userRecord.Email

	return auth, nil
}
