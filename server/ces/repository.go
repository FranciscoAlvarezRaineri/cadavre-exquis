package ces

import (
	"cadavre-exquis/models"
	"log"
	"strings"
)

func createCE(ce models.CE, contribution models.Contribution) *models.CE {
	text := generateReveal(contribution.Text, ce.RevealAmount)
	newCE := models.CE{
		Title:         ce.Title,
		Contributions: []models.Contribution{contribution},
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
