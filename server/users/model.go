package users

type CERef struct {
	ID     string `firestore:"id"`
	Closed bool   `firestore:"closed"`
	Title  string `firestore:"title"`
	Reveal string `firestore:"reveal"`
}

type User struct {
	UID         string           `firestore:"-"`
	UserName    string           `firestore:"user_name"`
	Email       string           `firestore:"email"`
	Ces         map[string]CERef `firestore:"ces"`
	Created     []CERef          `firestore:"created"`
	Contributed []CERef          `firestore:"contributed"`
}
