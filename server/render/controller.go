package render

import (
	"log"
	"net/http"

	"cadavre-exquis/ces"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	uid, _ := c.Get("uid")
	log.Printf("UID: %s", uid)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"main": "/home",
	})
}

func Home(c *gin.Context) {
	templ := "home.html"
	if c.GetHeader("HX-Request") != "true" {
		templ = "index.html"
	}
	getCE, _ := c.Get("ce")
	ce := getCE.(*ces.CE)
	last_contribution := c.GetBool("last_contribution")

	c.HTML(http.StatusOK, templ, gin.H{
		"id":                ce.ID,
		"reveal":            ce.Reveal,
		"reveal_amount":     ce.RevealAmount,
		"last_contribution": last_contribution,
		"main":              "/home",
	})
}

func User(c *gin.Context) {
	templ := "signin.html"
	if c.GetHeader("HX-Request") != "true" {
		templ = "index.html"
	}

	c.HTML(http.StatusOK, templ, gin.H{
		"main": "/user",
	})
}

func NewCE(c *gin.Context) {
	templ := "newce.html"
	if c.GetHeader("HX-Request") != "true" {
		templ = "index.html"
	}

	c.HTML(http.StatusOK, templ, gin.H{
		"main": "/newce",
	})
}

func ContributeToCE(c *gin.Context) {
	err := c.Errors
	if err != nil {
		templ := "contribution_error.html"
		c.HTML(http.StatusOK, templ, gin.H{
			"main": "/newce",
			"msg":  err[0],
		})
		return
	}

	templ := c.GetString("templ")
	texts, _ := c.Get("texts")
	c.HTML(http.StatusOK, templ, gin.H{
		"main":  "/newce",
		"texts": texts,
	})
}
