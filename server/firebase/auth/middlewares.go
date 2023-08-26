package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	idToken, err := c.Cookie("accessToken")
	if err != nil {
		c.Error(err)
		c.Set("uid", "")
		c.Next()
		return
	}
	token, err := validateToken(idToken)
	if err != nil {
		c.Error(err)
		c.Set("uid", "")
		c.Next()
		return
	}

	c.Set("uid", token.UID)
	c.Next()
}

func AuthGuard(c *gin.Context) {
	if len(c.Errors.Errors()) != 0 {
		c.Set("templ", "error.html")
		c.Set("result", gin.H{"msg": "please, sign in first:"})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
