package firestore

import (
	"log"

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
