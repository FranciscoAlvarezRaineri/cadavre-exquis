package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Env  string
	Data interface{}
}

func RespondJSON(c *gin.Context) {
	c.Next()

	data, _ := c.Get("data")

	c.JSON(-1, data)
}

func RenderHTML(c *gin.Context) {
	c.Set("templ", "index.gohtml")
	c.Set("data", gin.H{})
	result := &Result{Env: os.Getenv("ENV")}

	c.Next()

	templ := c.GetString("templ")
	data, _ := c.Get("data")
	result.Data = data

	c.HTML(-1, templ, result)
}
