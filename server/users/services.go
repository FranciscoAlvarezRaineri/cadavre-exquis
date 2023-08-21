package users

import (
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"log"
)

func GetAuthByUID(UID string) (*auth.Auth, error) {
	return auth.GetAuthByUID(UID)
}

func GetUserByUID(UID string) (*User, error) {
	dsnap, err := firestore.GetDocFromColByID("users", UID)
	if err != nil {
		return nil, err
	}

	user := &User{}
	dsnap.DataTo(user)
	return user, nil
}

func createUser(user_name string, email string, password string) (*User, error) {
	auth, err := auth.CreateUser(user_name, email, password)
	if err != nil {
		return nil, err
	}

	newUser := &User{}
	newUser.Email = auth.Email
	newUser.UserName = auth.UserName

	dsnap, err := firestore.SetDocInCol("users", auth.UID, newUser)
	if err != nil {
		return nil, err
	}

	user := &User{}
	dsnap.DataTo(user)
	log.Printf("user: %v", user)
	return user, nil
}

func ContributedTo(uid string, ceRef CERef) (bool, error) {
	contributions := map[string]interface{}{
		ceRef.ID: map[string]interface{}{
			"closed": ceRef.Closed,
			"title":  ceRef.Title,
			"reveal": ceRef.Reveal,
		},
	}
	_, err := firestore.AddCEToUser(uid, contributions)
	if err != nil {
		return false, err
	}

	return true, err
}
