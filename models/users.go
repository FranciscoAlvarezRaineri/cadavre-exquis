package models

type User struct {
	UID         string           `firestore:"-"`
	UserName    string           `firestore:"user_name"`
	Email       string           `firestore:"email"`
	Ces         map[string]CERef `firestore:"ces"`
	Created     map[string]CERef `firestore:"created"`
	Contributed map[string]CERef `firestore:"contributed"`
	Code        string           `firestore:"code"`
}
