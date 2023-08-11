package router

import (
	"cadavre-exquis/ces"
	"cadavre-exquis/users"

	"github.com/labstack/echo/v4"
	// "cadavre-exquis/firebase/auth"
)

func PublicRoutes(e *echo.Echo) {
	// e.POST("signin", users.SignIn)
	e.GET("/ces/:id", ces.GetCE)
}

func UserRoutes(e *echo.Echo) {
	e.GET("/user/:uid", users.GetUser)
}

func CesRoutes(e *echo.Echo) {
	e.PUT("/ces/:id", ces.ContributeToCE)
	e.POST("/ces", ces.CreateCE)
}
