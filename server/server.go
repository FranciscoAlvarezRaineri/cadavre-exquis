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

	router.SetTrustedProxies([]string{"127.0.0.1"})

	// router.Use(gin.Recovery())

	router.Static("/public", "./public")

	router.LoadHTMLGlob("views/**/*.html")

	router.Use(render.HTML)

	router.Use(auth.AuthCheck)

	router.GET("/", ces.GetRandomCE)
	router.GET("/home", ces.GetRandomCE)
	router.GET("/user", users.GetUser)
	router.POST("/user", users.CreateUser)
	router.GET("/signin", users.SignIn)
	router.GET("/signup", users.SignUp)
	router.GET("/ce/:id", ces.GetCE)
	router.GET("/newce", ces.NewCE)

	router.Use(auth.AuthGuard)
	router.Use(users.GetUserMid)

	router.POST("/ces", ces.CreateCE)
	router.PUT("/ces/:id", ces.ContributeToCE)

	router.Run("localhost:8080")
}
