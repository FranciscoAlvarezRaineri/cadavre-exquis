package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	user, err := GetUserByUID(uid)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.HTML(http.StatusOK, "contribution.html", gin.H{
		"msg": user.UID,
	})
}
