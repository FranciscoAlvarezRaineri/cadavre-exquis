package firestore

import (
	"log"

	"cadavre-exquis/context"
	"cadavre-exquis/firebase"

	"cloud.google.com/go/firestore"
)

type UpdateArray = []firestore.Update

var ArrayUnion = firestore.ArrayUnion
var MergeAll = firestore.MergeAll

var Client = initFirestore()

func initFirestore() *firestore.Client {
	client, err := firebase.App.Firestore(context.Context)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func Close() {
	Client.Close()
}

func GetDoc(collection string, id string) (*firestore.DocumentSnapshot, error) {
	dsnap, err := Client.Collection(collection).Doc(id).Get(context.Context)
	if err != nil {
		return nil, err
	}
	return dsnap, nil
}

func AddDoc(collection string, doc interface{}) (*firestore.DocumentSnapshot, error) {
	newDoc, _, err := Client.Collection(collection).Add(context.Context, doc)
	if err != nil {
		return nil, err
	}

	dsnap, err := newDoc.Get(context.Context)
	if err != nil {
		return nil, err
	}

	return dsnap, nil
}

func SetDoc(collection string, id string, doc interface{}, opts firestore.SetOption) (*firestore.DocumentSnapshot, error) {
	if opts == nil {
		_, err := Client.Collection(collection).Doc(id).Set(context.Context, doc)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := Client.Collection(collection).Doc(id).Set(context.Context, doc, opts)
		if err != nil {
			return nil, err
		}
	}

	dsnap, err := Client.Collection(collection).Doc(id).Get(context.Context)
	if err != nil {
		return nil, err
	}

	return dsnap, nil
}

func UpdateDoc(collection string, id string, update []firestore.Update) (*firestore.WriteResult, error) {
	return Client.Collection(collection).Doc(id).Update(context.Context, update)
}

type Where struct {
	Key      string
	Operator string
	Value    any
}

func GetAll2Where(collection string, where1 Where, where2 Where) ([]*firestore.DocumentSnapshot, error) {
	dsnaps, err := Client.Collection(collection).Where(where1.Key, where1.Operator, where1.Value).Where(where2.Key, where2.Operator, where2.Value).Documents(context.Context).GetAll()
	if err != nil {
		return nil, err
	}

	return dsnaps, nil
}
