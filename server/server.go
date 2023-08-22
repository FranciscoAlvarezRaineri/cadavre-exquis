package main

import (
	"cadavre-exquis/ces"
	"cadavre-exquis/firebase/auth"
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/render"
	"cadavre-exquis/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	defer firestore.Close()

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	// router.Use(gin.Recovery())

	router.Static("/public", "./public")

	router.LoadHTMLGlob("views/**/*.html")

	router.Use(auth.AuthCheck)

	router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/home") })
	router.GET("/home", ces.GetRandomCE, render.HTML)
	router.GET("/user", users.GetUser, render.HTML)
	router.POST("/user", users.CreateUser, render.HTML)
	router.GET("/signin", users.SignIn, render.HTML)
	router.GET("/signup", users.SignUp, render.HTML)
	router.GET("/newce", ces.NewCE, render.HTML)
	router.GET("/ce/:id", ces.GetCE, render.HTML)

	router.Use(auth.AuthGuard)
	router.Use(users.GetUserMid)

	router.POST("/ces", ces.CreateCE, render.HTML)
	router.PUT("/ces/:id", ces.ContributeToCE, render.HTML)

	router.Run("localhost:8080")
}
