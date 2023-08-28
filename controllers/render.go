package controllers

import (
	"github.com/gin-gonic/gin"
)

func RespondJSON(c *gin.Context) {
	c.Next()

	result := c.MustGet("result")

	c.JSON(-1, result)
}

func RenderHTML(c *gin.Context) {
	c.Next()

	templ := c.GetString("templ")

	result := c.MustGet("result").(gin.H)

	c.HTML(-1, templ, result)
}
