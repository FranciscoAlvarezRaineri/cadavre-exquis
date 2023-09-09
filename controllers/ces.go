package controllers

import (
	"cadavre-exquis/ces"
	email_service "cadavre-exquis/email"
	"cadavre-exquis/users"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Params.ByName("id")
	ce, err := ces.GetCE(id)
	if err != nil {
		c.Set("templ", "error.gohtml")

		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	texts := ces.GetFullText(ce.Contributions)
	c.Set("main", "ce.gohtml")
	data := gin.H{
		"title": ce.Title,
		"texts": texts,
	}

	c.Status(http.StatusOK)
	c.Set("templ", "ce.gohtml")
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
		c.Set("templ", "error.gohtml")

		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	newCE, err := ces.CreateNewCE(title, length, characters_max, words_min, reveal_amount, uid, userName, text)
	if err != nil {
		c.Set("templ", "error.gohtml")

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resultUser, err := users.ContributedTo(uid, newCE)
	if err != nil || !resultUser {
		c.Set("templ", "error.gohtml")

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
	c.Set("templ", "create_success.gohtml")
	c.Next()
}

func NewCEForm(c *gin.Context) {
	uid := c.GetString("uid")
	if uid == "" {
		c.Status(http.StatusOK)
		c.Set("msg", "please, sign in first:")
		c.Set("templ", "signin.gohtml")
		c.Next()
		return
	}

	c.Status(http.StatusOK)
	c.Set("templ", "newce.gohtml")
	c.Next()
}

func GetRandomCE(c *gin.Context) {
	uid, _ := c.Get("uid")
	id, _ := c.Cookie("active_ce")

	ce, err := ces.GetRandomCE(uid.(string), id)
	if err != nil {
		c.Set("templ", "error.gohtml")
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	last_contribution := ces.LastContribution(ce)
	data := gin.H{
		"id":                ce.ID,
		"reveal":            ce.Reveal,
		"reveal_amount":     ce.RevealAmount,
		"last_contribution": last_contribution,
		"characters_max":    ce.CharactersMax,
		"words_min":         ce.WordsMin,
	}

	c.Status(http.StatusOK)
	c.Set("templ", "home.gohtml")
	c.Set("data", data)

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("active_ce", ce.ID, 3600*24, "/", "127.0.0.1", false, true)

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
		c.Set("templ", "error.gohtml")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	uid := c.GetString("uid")
	userName := c.GetString("userName")

	success, err := ces.UpdateCE(id, closed, reveal_amount, uid, userName, text)
	if err != nil || !success {
		c.Set("templ", "error.gohtml")
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
		c.Set("templ", "error.gohtml")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
	c.Set("templ", "contribution_success.gohtml")

	if closed {
		texts := ces.GetFullText(ce.Contributions)

		email := c.GetString("email")
		email_service.SendClosedEmail(email, userName, ce.ID, ce.Title)

		c.Set("templ", "ce.gohtml")
		c.Set("data", gin.H{"texts": texts})
	}

	c.Next()
}
