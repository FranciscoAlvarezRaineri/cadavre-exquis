package ces

import (
	"cadavre-exquis/users"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Param("id")
	ce, err := GetCEById(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	texts := getFullText(ce.Contributions)
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
		c.Status(http.StatusBadRequest)
		c.Next()
		return
	}

	uid := c.GetString("uid")
	user := c.MustGet("user").(*users.User)
	userName := user.UserName

	ce := CE{
		Title:         title,
		Length:        length,
		CharactersMax: characters_max,
		WordsMin:      words_min,
		RevealAmount:  reveal_amount,
	}

	contribution := Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}

	newCE, err := CreateNewCE(ce, contribution)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Next()
		return
	}

	ceRef := users.CERef{
		ID:     newCE.ID,
		Title:  newCE.Title,
		Reveal: newCE.Reveal,
		Closed: newCE.Closed,
	}

	resultUser, err := users.ContributedTo(uid, ceRef)
	if err != nil || !resultUser {
		c.Error(err)
		c.Next()
		return
	}

	result := gin.H{}
	templ := "create_success.html"

	c.Status(http.StatusCreated)
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
		c.Status(http.StatusBadRequest)
		c.Next()
		return
	}

	uid := c.GetString("uid")
	user := c.MustGet("user").(*users.User)
	userName := user.UserName

	contribution := Contribution{
		Uid:      uid,
		UserName: userName,
		Text:     text,
	}

	success, err := UpdateCE(id, contribution, closed, reveal_amount)
	if err != nil || !success {
		c.Error(err)
		c.Next()
		return
	}

	ce, err := GetCEById(id)
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}

	ceRef := users.CERef{
		ID:     ce.ID,
		Title:  ce.Title,
		Reveal: ce.Reveal,
		Closed: ce.Closed,
	}

	successUser, err := users.ContributedTo(uid, ceRef)
	if err != nil || !successUser {
		c.Error(err)
		c.Next()
		return
	}

	var texts []string
	templ := "contribution_success.html"

	if closed {
		texts = getFullText(ce.Contributions)
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

func GetRandomCE(c *gin.Context) {
	templ := "home.html"
	if c.Request.Header.Get("HX-Request") != "true" {
		templ = "index.html"
	}

	ce, err := GetRandomPublicCE()
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	last_contribution := lastContribution(ce)
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

func NewCE(c *gin.Context) {
	templ := "newce.html"
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
		c.Set("result", gin.H{
			"main": "signin",
			"msg":  "please, sign in first:",
		})
		c.Next()
		return
	}

	result := gin.H{"main": "newce"}
	c.Status(http.StatusOK)
	c.Set("templ", templ)
	c.Set("result", result)
	c.Next()
}
