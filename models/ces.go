package models

type Contribution struct {
	Text     string `firestore:"text"`
	Uid      string `firestore:"uid"`
	UserName string `firestore:"username"`
}

type CERef struct {
	ID     string `firestore:"-"`
	Closed bool   `firestore:"closed"`
	Title  string `firestore:"title"`
	Reveal string `firestore:"reveal"`
}

type CEs map[string]CERef

type CE struct {
	ID            string         `firestore:"-"`
	Title         string         `firestore:"title"`
	Contributions []Contribution `firestore:"contributions"`
	Length        int            `firestore:"length"`
	CharactersMax int            `firestore:"characters_max"`
	WordsMin      int            `firestore:"words_min"`
	Reveal        string         `firestore:"reveal"`
	RevealAmount  int            `firestore:"reveal_amount"`
	Closed        bool           `firestore:"closed"`
	Idle          bool           `firestore:"idle"`
	Public        bool           `firestore:"public"`
}
