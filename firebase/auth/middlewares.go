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
		c.Set("templ", "signin.gohtml")
		c.Set("result", gin.H{"msg": "please, sign in first:"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uid := c.GetString("uid")
	auth, err := GetAuthByUID(uid)
	if err != nil {
		c.Set("templ", "error.gohtml")
		c.Set("result", gin.H{"msg": "something went wrong, please try again"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !auth.EmailVerified {
		c.Set("templ", "user.gohtml")
		c.Set("result", gin.H{"msg": "please verify your email address to continue."})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("verified", auth.EmailVerified)
	c.Next()
}
