package auth

import "github.com/gin-gonic/gin"

func ValidateToken(c *gin.Context) {
	idToken, err := c.Cookie("accessToken")
	if err != nil {
		c.Set("uid", "")
		c.Next()
		return
	}
	token, err := validateToken(idToken)
	if err != nil {
		c.Set("uid", "")
		c.Next()
		return
	}
	uid := token.UID
	c.Set("uid", uid)
	c.Next()
}
