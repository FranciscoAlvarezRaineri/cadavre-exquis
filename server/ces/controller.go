package ces

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Param("id")
	ce, err := GetCEById(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(http.StatusOK, "contribution.html", gin.H{
		"msg": ce.Title,
	})
}

func CreateCE(c *gin.Context) {
	c.Request.ParseForm()
	log.Printf("form: %v", c.Request.Form)

	title := c.Request.FormValue("title")

	text := c.Request.FormValue("text")

	length, err := strconv.Atoi(c.Request.FormValue("length"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	characters_limit, err := strconv.Atoi(c.Request.FormValue("characters_limit"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	words_limit, err := strconv.Atoi(c.Request.FormValue("words_limit"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	reveal_amount, err := strconv.Atoi(c.Request.FormValue("reveal_amount"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ce := CE{
		Title:           title,
		Length:          length,
		CharactersLimit: characters_limit,
		WordsLimit:      words_limit,
		RevealAmount:    reveal_amount,
	}

	contribution := Contribution{
		Uid:      "123456",
		UserName: "prueba",
		Text:     text,
	}

	newCE, err := CreateNewCE(ce, contribution)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"id":     newCE.ID,
		"reveal": newCE.Reveal,
	})
}

func ContributeToCE(c *gin.Context) {
	id := c.Param("id")
	last_contribution, err := strconv.ParseBool(c.Query("last_contribution"))
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}
	reveal_amount, err := strconv.Atoi(c.Query("reveal_amount"))
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}

	text := c.Request.FormValue("text")
	if text == "" {
		err := errors.New("text can't be empty")
		c.Error(err)
		c.Next()
		return
	}

	contribution := Contribution{
		Uid:      "123456",
		UserName: "prueba",
		Text:     text,
	}

	if len(text) < 4 {
		err := fmt.Errorf("text is too short. it should be at least %v words long", reveal_amount)
		c.Error(err)
		c.Next()
		return
	}

	result, err := UpdateCE(id, contribution, last_contribution, reveal_amount)
	if err != nil || !result {
		c.Error(err)
		c.Next()
		return
	}

	c.Set("texts", []string{})
	c.Set("templ", "contribution_success.html")

	if last_contribution {
		ce, err := GetCEById(id)
		if err != nil {
			c.Error(err)
			c.Next()
			return
		}
		texts := getFullText(ce.Contributions)
		c.Set("texts", texts)
		c.Set("templ", "ce.html")
	}

	c.Next()
}

func GetRandomCE(c *gin.Context) {
	ce, err := GetRandomPublicCE()
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}

	c.Set("ce", ce)
	c.Set("last_contribution", lastContribution(ce))

	c.Next()
}
