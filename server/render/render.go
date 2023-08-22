package render

import (
	"github.com/gin-gonic/gin"
)

func HTML(c *gin.Context) {
	templ := c.GetString("templ")

	result := c.MustGet("result").(gin.H)

	c.HTML(-1, templ, result)
}

func JSON(c *gin.Context) {
	result := c.MustGet("result")

	c.JSON(-1, result)
}
