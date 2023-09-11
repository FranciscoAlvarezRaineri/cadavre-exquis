package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Env   string
	Main  string
	Msg   string
	Data  interface{}
	Error error
	Auth  string
}

func RespondJSON(c *gin.Context) {
	c.Next()

	data, _ := c.Get("data")

	c.JSON(-1, data)
}

func RenderHTML(c *gin.Context) {
	c.Set("templ", "index.gohtml")
	c.Set("main", "home.gohtml")
	c.Set("msg", "")
	c.Set("data", gin.H{})

	c.Next()

	templ := c.GetString("templ")
	main := c.GetString("main")
	if c.Request.Header.Get("HX-Request") != "true" {
		main = templ
		templ = "index.gohtml"
	}

	result := &Result{}
	result.Env = os.Getenv("ENV")
	result.Main = main
	result.Msg = c.GetString("msg")
	result.Auth = c.GetHeader("Authorization")
	result.Error = c.Errors.Last()

	data, _ := c.Get("data")
	result.Data = data

	c.HTML(-1, templ, result)
}
