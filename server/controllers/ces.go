package controllers

import (
	"cadavre-exquis/ces"
	"cadavre-exquis/users"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Param("id")
	ce, err := ces.GetCE(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	texts := ces.GetFullText(ce.Contributions)
	result := gin.H{
		"title": ce.Title,
		"texts": texts,
	}
	templ := "ce.html"

	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}

func CreateCE(c *gin.Context) {
	title := c.Request.FormValue("title")

	text := c.Request.FormValue("text")

	length, err := strconv.Atoi(c.Request.FormValue("length"))
	if err != nil {
		c.Error(err)
	}

	characters_max, err := strconv.Atoi(c.Request.FormValue("characters_max"))
	if err != nil {
		c.Error(err)
	}

	words_min, err := strconv.Atoi(c.Request.FormValue("words_min"))
	if err != nil {
		c.Error(err)
	}

	reveal_amount, err := strconv.Atoi(c.Request.FormValue("reveal_amount"))
	if err != nil {
		c.Error(err)
	}

	if len(c.Errors.Errors()) != 0 {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	newCE, err := ces.CreateNewCE(title, length, characters_max, words_min, reveal_amount, uid, userName, text)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resultUser, err := users.ContributedTo(uid, newCE)
	if err != nil || !resultUser {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := gin.H{}
	templ := "create_success.html"

	c.Status(http.StatusCreated)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}

func NewCE(c *gin.Context) {
	templ := "newce.html"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.html"
	}

	uid := c.GetString("uid")
	if uid == "" {
		templ = "signin.html"
		if c.Request.Header.Get("HX-Request") != "true" {
			templ = "index.html"
		}

		result := gin.H{
			"main": "signin",
			"msg":  "please, sign in first:",
		}
		c.Status(http.StatusOK)
		c.Set("templ", templ)
		c.Set("result", result)
		c.Next()
		return
	}

	result := gin.H{"main": "newce"}
	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}

func GetRandomCE(c *gin.Context) {
	templ := "home.html"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.html"
	}

	ce, err := ces.GetRandomCE()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	last_contribution := ces.LastContribution(ce)
	result := gin.H{
		"main":              "home",
		"id":                ce.ID,
		"reveal":            ce.Reveal,
		"reveal_amount":     ce.RevealAmount,
		"last_contribution": last_contribution,
		"characters_max":    ce.CharactersMax,
		"words_min":         ce.WordsMin,
	}

	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}

func ContributeToCE(c *gin.Context) {
	id := c.Param("id")

	closed, err := strconv.ParseBool(c.Query("last_contribution"))
	if err != nil {
		c.Error(err)
	}

	reveal_amount, err := strconv.Atoi(c.Query("reveal_amount"))
	if err != nil {
		c.Error(err)
	}

	text := c.Request.FormValue("text")
	if len(strings.Split(text, " ")) < reveal_amount {
		err := fmt.Errorf("text is too short. it should be at least %v words long", reveal_amount)
		c.Error(err)
	}

	if len(c.Errors.Errors()) != 0 {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	success, err := ces.UpdateCE(id, closed, reveal_amount, uid, userName, text)
	if err != nil || !success {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ce, err := ces.GetCE(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	successUser, err := users.ContributedTo(uid, ce)
	if err != nil || !successUser {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var texts []string
	templ := "contribution_success.html"

	if closed {
		texts = ces.GetFullText(ce.Contributions)
		templ = "ce.html"
	}

	result := gin.H{
		"texts": texts,
	}

	c.Status(http.StatusCreated)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}
