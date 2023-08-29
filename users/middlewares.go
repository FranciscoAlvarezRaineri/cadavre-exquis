package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserMid(c *gin.Context) {
	uid := c.GetString("uid")
	user, err := GetUser(uid)
	if err != nil {
		c.Set("result", gin.H{})
		c.Set("templ", "error.html")
		c.AbortWithError(http.StatusNotFound, err)
	}
	c.Set("user", user)
	c.Set("userName", user.UserName)
	c.Set("email", user.Email)
	c.Next()
}
