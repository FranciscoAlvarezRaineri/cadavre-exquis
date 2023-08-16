package logs

import (
	"cadavre-exquis/firebase/firestore"
	"math/rand"
)

func GetRandomCEID() (string, error) {
	dsnaps, err := firestore.GetAllActiveCEs()
	if err != nil {
		return "", err
	}
	var ids []string
	for _, dsnap := range dsnaps {
		ids = append(ids, dsnap.Ref.ID)
	}
	id := ids[rand.Intn(len(ids))]
	return id, nil
}
