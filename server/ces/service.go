package ces

import (
	"cadavre-exquis/firebase/firestore"
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
	ce := models.CE{
		Title:         title,
		Length:        length,
		CharactersMax: characters_max,
		WordsMin:      words_min,
		RevealAmount:  reveal_amount,
	}

	contribution := models.Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}

	newCE := createCE(ce, contribution)
	dsnap, err := firestore.AddDocInCol("ces", newCE)
	if err != nil {
		return nil, err
	}
	result := &models.CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID
	return result, nil
}

func GetCEById(id string) (*models.CE, error) {
	dsnap, err := firestore.GetDocFromColByID("ces", id)
	if err != nil {
		return nil, err
	}
	result := &models.CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID
	return result, nil
}

func UpdateCE(id string, closed bool, reveal_amount int, uid string, userName string, text string) (bool, error) {
	contribution := models.Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}

	reveal := generateReveal(contribution.Text, reveal_amount)

	_, err := firestore.UpdateCE(id, contribution, reveal, closed)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetRandomPublicCE() (*models.CE, error) {
	dsnap, err := firestore.GetRandomPublicCE()
	if err != nil {
		return nil, err
	}

	ce := &models.CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
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
