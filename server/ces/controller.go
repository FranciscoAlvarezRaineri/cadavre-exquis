package ces

import (
	"cadavre-exquis/users"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCE(c *gin.Context) {
	id := c.Param("id")
	ce, err := GetCEById(id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Next()
		return
	}

	c.Status(http.StatusOK)
	c.Set("ce", ce)
	c.Next()
}

func CreateCE(c *gin.Context) {
	title := c.Request.FormValue("title")

	text := c.Request.FormValue("text")

	length, err := strconv.Atoi(c.Request.FormValue("length"))
	if err != nil {
		c.Error(err)
	}

	characters_limit, err := strconv.Atoi(c.Request.FormValue("characters_limit"))
	if err != nil {
		c.Error(err)
	}

	words_limit, err := strconv.Atoi(c.Request.FormValue("words_limit"))
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
	user, err := users.GetUserByUID(uid)
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}
	userName := user.UserName

	ce := CE{
		Title:           title,
		Length:          length,
		CharactersLimit: characters_limit,
		WordsLimit:      words_limit,
		RevealAmount:    reveal_amount,
	}

	contribution := Contribution{
		Uid:      uid, // username and uid should come from request!!!!!!!
		UserName: userName,
		Text:     text,
	}

	newCE, err := CreateNewCE(ce, contribution)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Next()
		return
	}

	c.Set("ce", newCE)
	c.Set("templ", "home.html")
	c.Next()
}

func ContributeToCE(c *gin.Context) {
	id := c.Param("id")

	last_contribution, err := strconv.ParseBool(c.Query("last_contribution"))
	if err != nil {
		c.Error(err)
	}

	reveal_amount, err := strconv.Atoi(c.Query("reveal_amount"))
	if err != nil {
		c.Error(err)
	}

	text := c.Request.FormValue("text")
	if len(text) < reveal_amount {
		err := fmt.Errorf("text is too short. it should be at least %v words long", reveal_amount)
		c.Error(err)
	}

	if len(c.Errors.Errors()) != 0 {
		c.Status(http.StatusBadRequest)
		c.Next()
		return
	}

	uid := c.GetString("uid")
	user, err := users.GetUserByUID(uid)
	if err != nil {
		c.Error(err)
		c.Next()
		return
	}
	userName := user.UserName

	contribution := Contribution{
		Uid:      uid, /// this data should come from request!!!!!!!!!!!!1
		UserName: userName,
		Text:     text,
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
