package ces

import (
	"cadavre-exquis/models"
)

func CreateNewCE(
	title string,
	length int,
	characters_max int,
	words_min int,
	reveal_amount int,
	uid string,
	userName string,
	text string,
) (*models.CE, error) {
	contribution := models.Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}
	reveal := generateReveal(contribution.Text, reveal_amount)
	newCE := &models.CE{
		Title:         title,
		Contributions: []models.Contribution{contribution},
		Length:        length,
		CharactersMax: characters_max,
		WordsMin:      words_min,
		Reveal:        reveal,
		RevealAmount:  reveal_amount,
		Closed:        false,
		Idle:          true,
		Public:        true,
	}

	ce, err := saveCEToDB(newCE)
	if err != nil {
		return nil, err
	}

	return ce, nil
}

func GetCE(id string) (*models.CE, error) {
	ce, err := getCEById(id)
	if err != nil {
		return nil, err
	}

	return ce, nil
}

func UpdateCE(id string, closed bool, reveal_amount int, uid string, userName string, text string) (bool, error) {
	contribution := &models.Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}
	reveal := generateReveal(contribution.Text, reveal_amount)

	_, err := updateCE(id, contribution, reveal, closed)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetRandomCE() (*models.CE, error) {
	ce, err := getRandomPublicIdleCE()
	if err != nil {
		return nil, err
	}

	return ce, nil
}

func LastContribution(ce *models.CE) bool {
	return ce.Length == (len(ce.Contributions) + 1)
}

func GetFullText(contributions []models.Contribution) []string {
	var fullText []string
	for _, contribution := range contributions {
		fullText = append(fullText, contribution.Text)
	}
	return fullText
}
