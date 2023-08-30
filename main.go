package main

import (
	"cadavre-exquis/controllers"
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/users"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer firestore.Close()

	godotenv.Load(".env")

	router := gin.Default()

	router.ForwardedByClientIP = true

	router.SetTrustedProxies([]string{
		"0.0.0.0",
		"100.20.92.101",
		"44.225.181.72",
		"44.227.217.144",
	})

	router.Use(gin.Recovery())

	router.Static("/public", "./public")

	router.LoadHTMLGlob("views/**/*.gohtml")

	router.Use(controllers.RenderHTML)

	router.Use(auth.AuthCheck)

	router.GET("/", controllers.GetRandomCE)
	router.GET("/home", controllers.GetRandomCE)
	router.GET("/user", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.GET("/signin", controllers.SignIn)
	router.GET("/signup", controllers.SignUp)
	router.GET("/newce", controllers.NewCE)

	router.Use(auth.AuthGuard)

	router.Use(users.GetUserMid)

	router.POST("/ces", controllers.CreateCE)
	router.GET("/ces/:id", controllers.GetCE)
	router.PUT("/ces/:id", controllers.ContributeToCE)

	router.Run(os.Getenv("HOST_PORT"))
}
