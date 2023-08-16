package auth

import (
	"cadavre-exquis/context"
	"cadavre-exquis/firebase"
	"log"

	"firebase.google.com/go/auth"
)

var client = initAuth()

type User struct {
	UID      string
	userName string
	email    string
}

// var uid = "7YzaAOa0m2XYrPTi9pNCqe5hniY2"

func initAuth() *auth.Client {
	var client, err = firebase.App.Auth(context.Context)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func GetUserByUID(uid string) (*User, error) {
	userRecord, err := client.GetUser(context.Context, uid)
	if err != nil {
		return nil, err
	}

	/*
		userDB, err := firestore.GetDocFromColByID("users", userRecord.UID)
		if err != nil {
			log.Fatalf("error getting user from db %s: %v\n", uid, err)
		}*/

	user := &User{}
	user.UID = userRecord.UID
	user.userName = userRecord.DisplayName
	user.email = userRecord.Email

	log.Printf("Successfully fetched user data: %v\n", user.email)

	return user, nil
}

func validateToken(idToken string) (*auth.Token, error) {
	return client.VerifyIDToken(context.Context, idToken)
}
