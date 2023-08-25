package users

import (
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/models"
	"log"
)

func GetUserByUID(UID string) (*models.User, error) {
	dsnap, err := firestore.GetDocFromColByID("users", UID)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	dsnap.DataTo(user)
	return user, nil
}

func CreateUser(user_name string, email string, password string) (*models.User, error) {
	auth, err := auth.CreateUser(user_name, email, password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{}
	newUser.Email = auth.Email
	newUser.UserName = auth.DisplayName

	dsnap, err := firestore.SetDocInCol("users", auth.UID, newUser)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	dsnap.DataTo(user)
	log.Printf("user: %v", user)
	return user, nil
}

func ContributedTo(uid string, ce *models.CE) (bool, error) {
	ceRef := models.CERef{
		ID:     ce.ID,
		Title:  ce.Title,
		Reveal: ce.Reveal,
		Closed: ce.Closed,
	}

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

func GetClosedContributions(ces map[string]models.CERef) *[]models.CERef {
	var contributions []models.CERef
	for id, ce := range ces {
		if ce.Closed {
			contribution := ce
			contribution.ID = id
			contributions = append(contributions, contribution)
		}
	}

	return &contributions
}
