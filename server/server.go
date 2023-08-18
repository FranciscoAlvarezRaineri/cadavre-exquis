package main

import (
	"cadavre-exquis/ces"
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/render"
	"cadavre-exquis/users"

	"github.com/gin-gonic/gin"
)

func main() {
	defer firestore.Close()

	router := gin.Default()

	router.Use(gin.Recovery())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Static("/public", "./public")

	router.LoadHTMLGlob("views/**/*.html")

	router.Use(auth.AuthCheck)

	router.GET("/", render.Index)
	router.GET("/home", ces.GetRandomCE, render.Home)
	router.GET("/user", users.GetUser, render.User)

	router.Use(auth.AuthGuard)

	router.GET("/newce", render.NewCEForm)
	router.POST("/ces", ces.CreateCE, render.CreateCE) // cambiar a un render específico de mensaje de creación de CE
	router.PUT("/ces/:id", ces.ContributeToCE, render.ContributeToCE)

	router.Use(func(c *gin.Context) {
		c.Abort()
	})

	router.Run("localhost:8080")
}
