package ces

import (
	"log"
	"strings"
)

type Contribution struct {
	Text     string `firestore:"text,omitempty"`
	Uid      string `firestore:"uid,omitempty"`
	UserName string `firestore:"username,omitempty"`
}

type CE struct {
	ID              string         `firestore:"id,omitempty"`
	Title           string         `firestore:"title,omitempty"`
	Contributions   []Contribution `firestore:"contributions,omitempty"`
	Length          int            `firestore:"length,omitempty"`
	CharactersLimit int            `firestore:"characters_limit,omitempty"`
	WordsLimit      int            `firestore:"words_limit,omitempty"`
	Reveal          string         `firestore:"reveal,omitempty"`
	RevealAmount    int            `firestore:"reveal_amount,omitempty"`
	Closed          bool           `firestore:"closed,omitempty"`
	Idle            bool           `firestore:"idle,omitempty"`
	Public          bool           `firestore:"public,omitempty"`
}

func createCE(ce CE, contribution Contribution) *CE {
	text := generateReveal(contribution.Text, ce.RevealAmount)
	newCE := CE{
		Title:           ce.Title,
		Contributions:   []Contribution{contribution},
		Length:          ce.Length,
		CharactersLimit: ce.CharactersLimit,
		WordsLimit:      ce.WordsLimit,
		Reveal:          text,
		RevealAmount:    ce.RevealAmount,
		Closed:          false,
		Idle:            true,
		Public:          true}
	return &newCE
}

func generateReveal(text string, reveal_amount int) string {
	split := strings.Split(text, " ")
	log.Printf("split: %v. length: %v.", split, len(split)-reveal_amount)
	split = split[len(split)-reveal_amount:]
	output := strings.Join(split, " ")
	return output
}

func checkCompleteCE(ce *CE) bool {
	return ce.Length <= (len(ce.Contributions) - 1) || ce.Closed
}