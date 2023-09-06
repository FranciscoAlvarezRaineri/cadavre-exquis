package controllers

import (
	email_service "cadavre-exquis/email"
	"cadavre-exquis/users"
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
		c.Set("main", "signin.gohtml")

		c.Next()
		return
	}

	user, err := users.GetUser(uid)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("main", "error.gohtml")

		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	contributions := users.GetClosedContributions(user.Ces)

	data := gin.H{
		"user_name": user.UserName,
		"ces":       contributions,
	}

	c.Status(http.StatusOK)
	c.Set("main", "user.gohtml")
	c.Set("templ", templ)
	c.Set("data", data)
	c.Next()
}

func CreateUser(c *gin.Context) {
	user_name := c.Request.FormValue("user_name")

	email := c.Request.FormValue("email")

	password := c.Request.FormValue("password")

	user, err := users.CreateUser(user_name, email, password)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("main", "error.gohtml")

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	email_service.SendConfirmationEmail(user.Email, user.UserName, user.UID, user.Code)

	c.Status(http.StatusCreated)
	c.Set("templ", "home.gohtml")
	c.Set("main", "main.gohtml")
	c.Set("data", gin.H{})
	c.Next()
}

func ConfirmEmail(c *gin.Context) {
	uid := c.Param("uid")
	code := c.Param("code")
	user, err := users.ConfirmEmail(uid, code)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("main", "error.gohtml")

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
	c.Set("templ", "home.gohtml")
	c.Set("data", gin.H{"username": user.UserName})
	c.Next()
}

func SignIn(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signin.gohtml")
	c.Set("main", "signin.gohtml")
	c.Next()
}

func SignUp(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signup.gohtml")
	c.Set("main", "signin.gohtml")
	c.Next()
}
