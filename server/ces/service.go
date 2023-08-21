package ces

import (
	"cadavre-exquis/firebase/firestore"
)

func CreateNewCE(
	ce CE,
	contribution Contribution,
) (*CE, error) {
	newCE := createCE(ce, contribution)
	dsnap, err := firestore.AddDocInCol("ces", newCE)
	if err != nil {
		return nil, err
	}
	result := &CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID
	return result, nil
}

func GetCEById(id string) (*CE, error) {
	dsnap, err := firestore.GetDocFromColByID("ces", id)
	if err != nil {
		return nil, err
	}
	result := &CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID
	return result, nil
}

func UpdateCE(id string, contribution Contribution, closed bool, reveal_amount int) (bool, error) {
	reveal := generateReveal(contribution.Text, reveal_amount)

	_, err := firestore.UpdateCE(id, contribution, reveal, closed)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetRandomPublicCE() (*CE, error) {
	dsnap, err := firestore.GetRandomPublicCE()
	if err != nil {
		return nil, err
	}

	ce := &CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
	return ce, nil
}

func lastContribution(ce *CE) bool {
	return ce.Length == (len(ce.Contributions) + 1)
}

func getFullText(contributions []Contribution) []string {
	var fullText []string
	for _, contribution := range contributions {
		fullText = append(fullText, contribution.Text)
	}
	return fullText
}
