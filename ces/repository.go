package ces

import (
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/models"
	"log"
	"strings"
)

func saveCEToDB(newCE *models.CE) (*models.CE, error) {
	dsnap, err := firestore.AddDoc("ces", newCE)
	if err != nil {
		return nil, err
	}

	result := &models.CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID

	return result, nil
}

func getCEById(id string) (*models.CE, error) {
	dsnap, err := firestore.GetDoc("ces", id)
	if err != nil {
		return nil, err
	}
	result := &models.CE{}
	dsnap.DataTo(result)
	result.ID = dsnap.Ref.ID
	return result, nil
}

func updateCE(id string, contribution *models.Contribution, reveal string, closed bool) (bool, error) {
	update := firestore.UpdateArray{
		{Path: "contributions", Value: firestore.ArrayUnion(contribution)},
		{Path: "reveal", Value: reveal},
		{Path: "closed", Value: closed},
	}
	_, err := firestore.UpdateDoc("ces", id, update)

	if err != nil {
		return false, err
	}

	return true, err
}

func generateReveal(text string, reveal_amount int) string {
	split := strings.Split(text, " ")
	log.Printf("split: %v. length: %v.", split, len(split)-reveal_amount)
	log.Printf("reveal: %v.", strings.Join(split[len(split)-reveal_amount:], " "))
	split = split[len(split)-reveal_amount:]
	output := strings.Join(split, " ")
	return output
}

func getRandomPublicIdleCE() (*models.CE, error) {
	where1 := firestore.Where{
		Key:      "public",
		Operator: "==",
		Value:    true,
	}
	where2 := firestore.Where{
		Key:      "closed",
		Operator: "==",
		Value:    false,
	}

	dsnap, err := firestore.GetAll2Where("ces", where1, where2)
	if err != nil {
		return nil, err
	}

	ce := &models.CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
	return ce, nil
}
