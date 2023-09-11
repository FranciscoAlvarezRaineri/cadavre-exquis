package auth

import (
	"cadavre-exquis/context"
	"cadavre-exquis/firebase"
	"log"

	"firebase.google.com/go/auth"
)

var client = initAuth()

type Auth = auth.UserRecord

func initAuth() *auth.Client {
	client, err := firebase.App.Auth(context.Context)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetAuthByUID(uid string) (*auth.UserRecord, error) {
	userRecord, err := client.GetUser(context.Context, uid)
	if err != nil {
		return nil, err
	}

	return userRecord, nil
}

func ValidateToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(context.Context, idToken)
}

func CreateAuth(user_name string, email string, password string) (*auth.UserRecord, error) {
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

	return userRecord, nil
}

func ConfirmEmail(uid string) (*auth.UserRecord, error) {
	params := (&auth.UserToUpdate{}).EmailVerified(true)
	userRecord, err := client.UpdateUser(context.Context, uid, params)
	if err != nil {
		return nil, err
	}
	return userRecord, nil
}
