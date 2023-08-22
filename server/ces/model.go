package ces

import (
	"log"
	"strings"
)

type Contribution struct {
	Text     string `firestore:"text"`
	Uid      string `firestore:"uid"`
	UserName string `firestore:"username"`
}

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

func createCE(ce CE, contribution Contribution) *CE {
	text := generateReveal(contribution.Text, ce.RevealAmount)
	newCE := CE{
		Title:         ce.Title,
		Contributions: []Contribution{contribution},
		Length:        ce.Length,
		CharactersMax: ce.CharactersMax,
		WordsMin:      ce.WordsMin,
		Reveal:        text,
		RevealAmount:  ce.RevealAmount,
		Closed:        false,
		Idle:          true,
		Public:        true}
	return &newCE
}

func generateReveal(text string, reveal_amount int) string {
	split := strings.Split(text, " ")
	log.Printf("split: %v. length: %v.", split, len(split)-reveal_amount)
	log.Printf("reveal: %v.", strings.Join(split[len(split)-reveal_amount:], " "))
	split = split[len(split)-reveal_amount:]
	output := strings.Join(split, " ")
	return output
}

func checkCompleteCE(ce *CE) bool {
	return ce.Length <= (len(ce.Contributions)-1) || ce.Closed
}
