package users

import (
	"cadavre-exquis/firebase/auth"
)

func GetUserByUID(UID string) (*auth.User, error) {
	return auth.GetUserByUID(UID)
}