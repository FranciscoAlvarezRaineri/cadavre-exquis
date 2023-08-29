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
	c.Set("templ", "index.html")
	c.Set("result", gin.H{})

	c.Next()

	templ := c.GetString("templ")
	result, _ := c.Get("result")

	c.HTML(-1, templ, result)
}
