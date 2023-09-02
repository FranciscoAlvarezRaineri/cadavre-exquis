package users

import (
	"cadavre-exquis/models"
	"errors"
)

func GetUser(UID string) (*models.User, error) {
	user, err := getUserByUID(UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user_name string, email string, password string) (*models.User, error) {
	auth, err := createAuth(user_name, email, password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    auth.Email,
		UserName: auth.DisplayName,
		Code:     "test_code",
	}

	user, err := saveUserToDB(newUser, auth.UID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ConfirmEmail(uid string, code string) (*models.User, error) {
	user, err := getUserByUID(uid)
	if err != nil {
		return nil, err
	}

	if user.Code != code {
		return nil, errors.New("something went wrong, please try again")
	}

	_, err = confirmEmail(uid)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ContributedTo(uid string, ce *models.CE) (bool, error) {
	contributions := models.CEs{
		ce.ID: models.CERef{
			Title:  ce.Title,
			Reveal: ce.Reveal,
			Closed: ce.Closed,
		},
	}
	_, err := addCEToUser(uid, contributions)
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
