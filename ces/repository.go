package ces

import (
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/models"
	"errors"
	"log"
	"math/rand"
	"strings"
)

func saveCEToDB(newCE *models.CE) (*models.CE, error) {
	dsnap, err := firestore.AddDoc("ces", newCE)
	if err != nil {
		return nil, err
	}

	ce := &models.CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID

	return ce, nil
}

func getCEById(id string) (*models.CE, error) {
	dsnap, err := firestore.GetDoc("ces", id)
	if err != nil {
		return nil, err
	}
	ce := &models.CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
	return ce, nil
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

func getAllRandomPublicIdleCEs(uid string) (*models.CE, error) {
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

	dsnaps, err := firestore.GetAll2Where("ces", where1, where2)
	if err != nil {
		return nil, err
	}

	if len(dsnaps) == 0 {
		return nil, errors.New("no available text to contribute to, please start a new one")
	}

	dsnap := dsnaps[rand.Intn(len(dsnaps))]
	ce := &models.CE{}
	dsnap.DataTo(ce)
	ce.ID = dsnap.Ref.ID
	return ce, nil

	/*
		ces := []*models.CE{}
		for _, dsnap := range dsnaps {
			var ce *models.CE
			mapstructure.Decode(dsnap.Data(), &ce)
			ce.ID = dsnap.Ref.ID
			ces = append(ces, ce)
		}
	*/
}

/*
func filterOutCEsByContributor(uid string, ces []*models.CE) []*models.CE {
	result := []*models.CE{}
	for _, ce := range ces {
		for _, contribution := range ce.Contributions {
			if contribution.Uid != uid {
				result = append(result, ce)
			}

		}
	}

	return result
}
*/
