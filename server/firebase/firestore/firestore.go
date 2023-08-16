package firestore

import (
	"log"
	"math/rand"

	"cadavre-exquis/context"
	"cadavre-exquis/firebase"

	"cloud.google.com/go/firestore"
)

var client = initFirestore()

func initFirestore() *firestore.Client {
	client, err := firebase.App.Firestore(context.Context)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func Close() {
	client.Close()
}

func GetDocFromColByID(collection string, id string) (*firestore.DocumentSnapshot, error) {
	dsnap, err := client.Collection(collection).Doc(id).Get(context.Context)
	if err != nil {
		return nil, err
	}
	return dsnap, nil
}

func AddDocInCol(collection string, doc interface{}) (*firestore.DocumentSnapshot, error) {
	newDoc, _, err := client.Collection(collection).Add(context.Context, doc)
	if err != nil {
		return nil, err
	}

	dsnap, err := newDoc.Get(context.Context)
	if err != nil {
		return nil, err
	}

	return dsnap, nil
}

func SetDocInCol(collection string, id string, doc interface{}) (*firestore.DocumentSnapshot, *firestore.WriteResult, error) {
	result, err := client.Collection(collection).Doc(id).Set(context.Context, doc)
	if err != nil {
		return nil, nil, err
	}

	dsnap, err := client.Collection(collection).Doc(id).Get(context.Context)
	if err != nil {
		return nil, nil, err
	}

	return dsnap, result, nil
}

func GetAllActiveCEs() ([]*firestore.DocumentSnapshot, error) {
	return client.Collection("logs/ces/active").Documents(context.Context).GetAll()
}

func GetRandomPublicCEs() (*firestore.DocumentSnapshot, error) {
	dsnaps, err := client.Collection("ces").Where("public", "==", true).Where("closed", "==", false).Documents(context.Context).GetAll()
	if err != nil {
		return nil, err
	}

	return dsnaps[rand.Intn(len(dsnaps))], nil
}

func UpdateCE(id string, contribution interface{}, reveal string, closed bool) (*firestore.WriteResult, error) {
	return client.Collection("ces").Doc(id).Update(context.Context, []firestore.Update{
		{Path: "contributions", Value: firestore.ArrayUnion(contribution)},
		{Path: "reveal", Value: reveal},
		{Path: "closed", Value: closed},
	})
}
