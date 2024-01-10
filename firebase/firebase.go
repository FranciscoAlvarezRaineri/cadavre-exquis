package firebase

import (
	"context"
	"log"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

var App = initApp()

func initApp() *firebase.App {
	opt := option.WithCredentialsFile("./cadavre-exquis-9c7af-firebase-adminsdk-uuwij-a4516003e3.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	return app
}
