package ces

type Contribution struct {
	Text     string `firestore:"text,omitempty"`
	Uid      string `firestore:"uid,omitempty"`
	UserName string `firestore:"username,omitempty"`
}

type CE struct {
	ID              string         `firestore:"id,omitempty"`
	Title           string         `firestore:"title,omitempty"`
	Contributions   []Contribution `firestore:"contributions,omitempty"`
	LengthLimit     int            `firestore:"length_limit,omitempty"`
	CharactersLimit int            `firestore:"characters_limit,omitempty"`
	WordsLimit      int            `firestore:"words_limit,omitempty"`
	Reveal          string         `firestore:"reveal,omitempty"`
	RevealAmount    int            `firestore:"reveal_amount,omitempty"`
	Closed          bool           `firestore:"closed,omitempty"`
	Idle            bool           `firestore:"idle,omitempty"`
}

func CreateEntityCE(ce CE, contribution Contribution) *CE {
	newCE := CE{
		Title:           ce.Title,
		Contributions:   []Contribution{contribution},
		LengthLimit:     ce.LengthLimit,
		CharactersLimit: ce.CharactersLimit,
		WordsLimit:      ce.WordsLimit,
		Reveal:          contribution.Text[len(contribution.Text)-ce.RevealAmount:],
		RevealAmount:    ce.RevealAmount,
		Closed:          false,
		Idle:            true}
	return &newCE
}
