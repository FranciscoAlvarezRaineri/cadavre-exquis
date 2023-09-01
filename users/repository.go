package users

import (
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/models"
)

func createAuth(user_name string, email string, password string) (*auth.Auth, error) {
	return auth.CreateAuth(user_name, email, password)
}

func saveUserToDB(newUser *models.User, uid string) (*models.User, error) {
	dsnap, err := firestore.SetDoc("users", uid, newUser, nil)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	dsnap.DataTo(user)
	user.UID = dsnap.Ref.ID

	return user, nil
}

func addCEToUser(uid string, contributions models.CEs) (*models.User, error) {
	dsnap, err := firestore.SetDoc("users", uid, map[string]interface{}{"ces": contributions}, firestore.MergeAll)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	dsnap.DataTo(user)
	user.UID = dsnap.Ref.ID

	return user, nil
}

func getUserByUID(uid string) (*models.User, error) {
	dsnap, err := firestore.GetDoc("users", uid)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	dsnap.DataTo(user)
	user.UID = dsnap.Ref.ID

	return user, nil
}

func confirmEmail(uid string) (*auth.Auth, error) {
	return auth.ConfirmEmail(uid)
}
