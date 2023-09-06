package controllers

import (
	"cadavre-exquis/ces"
	email_service "cadavre-exquis/email"
	"cadavre-exquis/users"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Params.ByName("id")
	ce, err := ces.GetCE(id)
	if err != nil {
		log.Print("hola")
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	texts := ces.GetFullText(ce.Contributions)
	data := gin.H{
		"title": ce.Title,
		"texts": texts,
		"main":  "ce",
	}
	templ := "ce.gohtml"

	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("data", data)
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
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": c.Errors})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	newCE, err := ces.CreateNewCE(title, length, characters_max, words_min, reveal_amount, uid, userName, text)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resultUser, err := users.ContributedTo(uid, newCE)
	if err != nil || !resultUser {
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	data := gin.H{}
	templ := "create_success.gohtml"

	c.Status(http.StatusCreated)
	c.Set("templ", templ)
	c.Set("data", data)
	c.Next()
}

func NewCEForm(c *gin.Context) {
	templ := "newce.gohtml"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.gohtml"
	}

	uid := c.GetString("uid")
	if uid == "" {
		templ = "signin.gohtml"
		if c.Request.Header.Get("HX-Request") != "true" {
			templ = "index.gohtml"
		}

		data := gin.H{
			"main": "signin",
			"msg":  "please, sign in first:",
		}
		c.Status(http.StatusOK)
		c.Set("templ", templ)
		c.Set("data", data)
		c.Next()
		return
	}

	data := gin.H{"main": "newce"}
	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("data", data)
	c.Next()
}

func GetRandomCE(c *gin.Context) {
	templ := "home.gohtml"
	uid, _ := c.Get("uid")

	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.gohtml"
	}

	id, err := c.Cookie("active_ce")

	ce, err := ces.GetRandomCE(uid.(string), id)
	if err != nil {
		c.Set("templ", "error.gohtml")
		if c.Request.Header.Get("HX-Request") != "true" {
			templ = "index.gohtml"
		}
		c.Set("data", gin.H{"error": err})
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	last_contribution := ces.LastContribution(ce)
	data := gin.H{
		"main":              "home",
		"id":                ce.ID,
		"reveal":            ce.Reveal,
		"reveal_amount":     ce.RevealAmount,
		"last_contribution": last_contribution,
		"characters_max":    ce.CharactersMax,
		"words_min":         ce.WordsMin,
	}

	c.SetCookie("active_ce", ce.ID, 3600, "/", "127.0.0.1", false, true)

	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("data", data)
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
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	success, err := ces.UpdateCE(id, closed, reveal_amount, uid, userName, text)
	if err != nil || !success {
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
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
		c.Set("templ", "index.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var texts []string
	templ := "contribution_success.gohtml"

	if closed {
		texts = ces.GetFullText(ce.Contributions)
		email := c.GetString("email")
		email_service.SendClosedEmail(email, userName, ce.ID, ce.Title)
		templ = "ce.gohtml"
	}

	data := gin.H{
		"texts": texts,
	}

	c.Status(http.StatusCreated)
	c.Set("templ", templ)
	c.Set("data", data)
	c.Next()
}
