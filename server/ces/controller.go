package ces

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCE(c echo.Context) error {
	id := c.Param("id")
	ce, err := GetCEById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, ce)
}

func CreateCE(c echo.Context) error {
	Title := c.QueryParam("title")

	Text := c.QueryParam("text")

	LengthLimit, err := strconv.Atoi(c.QueryParam("length"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid length limit.")
	}

	CharactersLimit, err := strconv.Atoi(c.QueryParam("character-limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid character limit.")
	}

	WordsLimit, err := strconv.Atoi(c.QueryParam("words-limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid words limit.")
	}

	RevealAmount, err := strconv.Atoi(c.QueryParam("reveal-limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid reveal limit.")
	}

	ce := CE{
		Title:           Title,
		LengthLimit:     LengthLimit,
		CharactersLimit: CharactersLimit,
		WordsLimit:      WordsLimit,
		RevealAmount:    RevealAmount,
	}

	contribution := Contribution{
		Uid:      "123456",
		UserName: "prueba",
		Text:     Text,
	}

	newCE, err := CreateNewCE(ce, contribution)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, newCE)
}

func ContributeToCE(c echo.Context) error {
	id := c.Param("id")

	Text := c.QueryParam("text")

	contribution := Contribution{
		Uid:      "123456",
		UserName: "prueba",
		Text:     Text,
	}

	oldCE, err := GetCEById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newCE, err := UpdateCE(oldCE, contribution, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, newCE)
}
