package main

import (
	"cadavre-exquis/ces"
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/render"

	"github.com/gin-gonic/gin"
)

func main() {
	defer firestore.Close()

	router := gin.Default()

	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Static("/public", "./public")

	
	router.LoadHTMLGlob("views/**/*.html")
	
	router.GET("/", render.Index)
	router.GET("/home", ces.GetRandomCE, render.Home)
	router.GET("/user", render.User)

	router.Use(auth.ValidateToken)

	router.GET("/newce", render.NewCE)
	router.PUT("/ces/:id", ces.ContributeToCE, render.ContributeToCE)

	router.POST("/ces", ces.CreateCE)

	router.Run("localhost:8080")
}
