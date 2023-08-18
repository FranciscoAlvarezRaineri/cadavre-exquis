package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	user, err := GetUserByUID(uid)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.Next()
		return
	}
	c.Status(http.StatusOK)
	c.Set("templ", "contribution.html")
	c.Set("user", user)
}
