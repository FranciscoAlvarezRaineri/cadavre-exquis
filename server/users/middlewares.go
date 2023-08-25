package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserMid(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := GetUser(uid)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	c.Set("user", user)
	c.Set("userName", user.UserName)
	c.Next()
}
