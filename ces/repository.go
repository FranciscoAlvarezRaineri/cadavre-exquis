package ces

import (
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/models"
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
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
	split = split[len(split)-reveal_amount:]
	output := strings.Join(split, " ")
	return output
}

func getRandomPublicIdleNotIDCE(uid string, id string) (*models.CE, error) {
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

	ces := []*models.CE{}
	for _, dsnap := range dsnaps {
		ce := &models.CE{}
		dsnap.DataTo(ce)
		ce.ID = dsnap.Ref.ID
		ces = append(ces, ce)
	}

	index := -1
	for i, ce := range ces {
		if ce.ID == id {
			index = i
		}
	}
	if index >= 0 {
		ces[index] = ces[len(ces)-1]
		ces = ces[:len(ces)-1]
	}

	filteredCEs := ces
	if uid != "" {
		count := 0
		for i, ce := range ces {
			log.Printf("contributionId: %s. userid: %s.", ce.ID, uid)
			for _, contribution := range ce.Contributions {
				if contribution.Uid == uid {
					log.Printf("ces: %v", filteredCEs)
					filteredCEs = removeFromSliceAtIndex(filteredCEs, (i - count))
					count = count + 1
					log.Printf("ces: %v", filteredCEs)
					break
				}
			}
		}
	}

	if len(dsnaps) == 0 {
		return nil, errors.New("no available text to contribute to, please start a new one")
	}

	ce := ces[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(ces))]

	return ce, nil
}

func removeFromSliceAtIndex(slice []*models.CE, i int) []*models.CE {
	return append(slice[:i], slice[i+1:]...)
}
