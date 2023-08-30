package controllers

import (
	"cadavre-exquis/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	templ := "user.html"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.html"
	}

	uid := c.GetString("uid")
	if len(uid) == 0 {
		templ = "signin.html"
		if c.Request.Header.Get("HX-Request") != "true" {
			templ = "index.html"
		}

		c.Status(http.StatusOK)
		c.Set("templ", templ)
		c.Set("result", gin.H{"main": "signin"})
		c.Next()
		return
	}

	user, err := users.GetUser(uid)
	if err != nil {
		c.Set("templ", "error.html")
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

	_, err := users.CreateUser(user_name, email, password)
	if err != nil {
		c.Set("templ", "error.html")
		c.Set("result", gin.H{"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
	c.Set("templ", "home.html")
	c.Set("result", gin.H{"error": err})
	c.Next()
}

func SignIn(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signin.html")
	c.Set("result", gin.H{"main": "signin"})
	c.Next()
}

func SignUp(c *gin.Context) {
	c.Status(http.StatusOK)
	c.Set("templ", "signup.html")
	c.Set("result", gin.H{"main": "signup"})
	c.Next()
}
