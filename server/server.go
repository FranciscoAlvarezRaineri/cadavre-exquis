package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	
	ctx := context.Background()
	app, err := connectFirebase(ctx)
	if err != nil {
		return
	}
	
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	
	e.GET("/", func(c echo.Context) error {
		doc, err := docAsMap(ctx, client)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "No such document in Firestore.")
		}
		return c.JSON(http.StatusOK, doc)
	})
	

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func connectFirebase(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("./cadavre-exquis-9c7af-firebase-adminsdk-uuwij-a4516003e3.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func docAsMap(ctx context.Context, client *firestore.Client) (map[string]interface{}, error) {
	dsnap, err := client.Collection("ces").Doc("17RnDMF2H0EYjQHbAZrW").Get(ctx)
	if err != nil {
					return nil, err
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	return m, nil
}