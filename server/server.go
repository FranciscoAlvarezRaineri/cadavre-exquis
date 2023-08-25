package main

import (
	"cadavre-exquis/controllers"
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/users"

	"github.com/gin-gonic/gin"
)

func main() {
	defer firestore.Close()

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(gin.Recovery())

	router.Static("/public", "./public")

	router.LoadHTMLGlob("views/**/*.html")

	router.Use(controllers.RenderHTML)

	router.Use(auth.AuthCheck)

	router.GET("/", controllers.GetRandomCE)
	router.GET("/home", controllers.GetRandomCE)
	router.GET("/user", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.GET("/signin", controllers.SignIn)
	router.GET("/signup", controllers.SignUp)
	router.GET("/ce/:id", controllers.GetCE)
	router.GET("/newce", controllers.NewCE)

	router.Use(auth.AuthGuard)

	router.Use(users.GetUserMid)

	router.POST("/ces", controllers.CreateCE)
	router.PUT("/ces/:id", controllers.ContributeToCE)

	router.Run("localhost:8080")
}
