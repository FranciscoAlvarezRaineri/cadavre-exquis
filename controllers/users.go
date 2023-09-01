package controllers

import (
	email_service "cadavre-exquis/email"
	"cadavre-exquis/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	templ := "user.gohtml"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.gohtml"
	}

	uid := c.GetString("uid")
	if len(uid) == 0 {
		templ = "signin.gohtml"
		if c.Request.Header.Get("HX-Request") != "true" {
			templ = "index.gohtml"
		}

		c.Status(http.StatusOK)
		c.Set("templ", templ)
		c.Set("result", gin.H{"main": "signin"})
		c.Next()
		return
	}

	user, err := users.GetUser(uid)
	if err != nil {
		c.Set("templ", "error.gohtml")
		c.Set("result", gin.H{"error": err})
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	contributions := users.GetClosedContributions(user.Ces)

	result := gin.H{
		"main":      "user",
		"user_name": user.UserName,
		"ces":       contributions,
	}

	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}

func CreateUser(c *gin.Context) {
	user_name := c.Request.FormValue("user_name")

	email := c.Request.FormValue("email")

	password := c.Request.FormValue("password")

	user, err := users.CreateUser(user_name, email, password)
	if err != nil {
		c.Set("templ", "error.gohtml")
		c.Set("result", gin.H{"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("user: %v", user)

	email_service.SendConfirmationEmail(user.Email, user.UserName, user.UID, user.Code)

	c.Status(http.StatusCreated)
	c.Set("templ", "home.gohtml")
	c.Set("result", gin.H{"error": err})
	c.Next()
}

func ConfirmEmail(c *gin.Context) {
	uid := c.Param("uid")
	code := c.Param("code")
	user, err := users.ConfirmEmail(uid, code)
	if err != nil {
		c.Set("templ", "error.gohtml")
		c.Set("result", gin.H{"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
	c.Set("templ", "home.gohtml")
	c.Set("result", gin.H{"username": user.UserName})
	c.Next()
}

func SignIn(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signin.gohtml")
	c.Set("result", gin.H{"main": "signin"})
	c.Next()
}

func SignUp(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signup.gohtml")
	c.Set("result", gin.H{"main": "signup"})
	c.Next()
}
