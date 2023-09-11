package users

import (
	"cadavre-exquis/firebase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	idToken := c.GetHeader("Authorization")
	if idToken == "" {
		c.Set("uid", "")
		c.Next()
		return
	}

	token, err := auth.ValidateToken(idToken)
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
		c.Set("msg", "please, sign in first:")

		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	uid := c.GetString("uid")
	auth, err := auth.GetAuthByUID(uid)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("main", "error.gohtml")
		c.Set("msg", "something went wrong, please try again")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	/*
		if !auth.EmailVerified {
			c.Set("templ", "user.gohtml")
			c.Set("data", gin.H{"msg": "please verify your email address to continue."})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	*/
	c.Set("verified", auth.EmailVerified)
	c.Next()
}

func GetUserMid(c *gin.Context) {
	uid := c.GetString("uid")

	user, err := GetUser(uid)
	if err != nil {
		c.Set("templ", "index.gohtml")
		c.Set("main", "error.gohtml")
		c.Set("data", gin.H{
			"main":  "error",
			"error": err})
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.Set("user", user)
	c.Set("userName", user.UserName)
	c.Set("email", user.Email)
	c.Next()
}
