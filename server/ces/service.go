package ces

import (
	"cadavre-exquis/firebase/firestore"
	"log"
)

func CreateNewCE(
	ce CE,
	contribution Contribution,
) (*CE, error) {
	newCE := CreateEntityCE(ce, contribution)
	dsnap, err := firestore.AddDocInCol("ces", newCE)
	if err != nil {
		return nil, err
	}
	result := &CE{}
	dsnap.DataTo(result)
	return result, nil
}

func GetCEById(id string) (*CE, error) {
	dsnap, err := firestore.GetDocFromColByID("ces", id)
	if err != nil {
		return nil, err
	}
	result := &CE{}
	dsnap.DataTo(result)
	return result, nil
}

func UpdateCE(ce *CE, contribution Contribution, id string) (*CE, error) {
	ce.Contributions = append(ce.Contributions, contribution)
	ce.Reveal = contribution.Text[len(contribution.Text)-ce.RevealAmount:]

	dsnap, result, err := firestore.SetDocInCol("ces", id, ce)
	if err != nil {
		return nil, err
	}
	log.Printf("CE updated succesfully: %v", result)
	updatedCE := &CE{}
	dsnap.DataTo(updatedCE)
	return updatedCE, nil
}
