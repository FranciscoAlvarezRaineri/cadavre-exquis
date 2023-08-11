package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	uid := c.Param("uid")
	user, err := GetUserByUID(uid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Not a valid UserID.")
	}
	return c.JSON(http.StatusOK, user)
}
