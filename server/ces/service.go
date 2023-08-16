package ces

import (
	"cadavre-exquis/firebase/firestore"
	"log"
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

func UpdateCE(id string, contribution Contribution, last_contribution bool, reveal_amount int) (bool, error) {
	closed := false
	if last_contribution {
		closed = true
	}
	reveal := generateReveal(contribution.Text, reveal_amount)

	_, err := firestore.UpdateCE(id, contribution, reveal, closed)
	if err != nil {
		return false, err
	}

	return true, err
}

/*
func UpdateCE(ce *CE, contribution Contribution) (*CE, error) {
ce.Contributions = append(ce.Contributions, contribution)
		ce.Reveal = generateReveal(contribution.Text, ce.RevealAmount)
		if checkCompleteCE(ce) {
			ce.Closed = true
		}
		dsnap, _, err := firestore.SetDocInCol("ces", ce.ID, ce)
		if err != nil {
			return nil, err
		}

		updatedCE := &CE{}
		dsnap.DataTo(updatedCE)
		updatedCE.ID = dsnap.Ref.ID

		return updatedCE, nil
}*/

func GetRandomPublicCE() (*CE, error) {
	dsnap, err := firestore.GetRandomPublicCEs()
	if err != nil {
		return nil, err
	}

	ce := &CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
	return ce, nil
}

func lastContribution(ce *CE) bool {
	log.Printf("lenght: %v. contributions: %v. result: %v.", ce.Length, len(ce.Contributions), ce.Length == (len(ce.Contributions)+1))
	return ce.Length == (len(ce.Contributions) + 1)
}

func getFullText(contributions []Contribution) []string {
	var fullText []string
	for _, contribution := range contributions {
		fullText = append(fullText, contribution.Text)
	}
	return fullText
}
