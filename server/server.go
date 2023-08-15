package main

import (
	"cadavre-exquis/firebase/firestore"
	"cadavre-exquis/render"
	"cadavre-exquis/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	defer firestore.Close()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Static("/public", "public")

	e.Renderer = render.TemplateRenderer

	router.PublicRoutes(e)

	e.Use()

	// e.Use(authMiddleware)

	router.UserRoutes(e)

	router.CesRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
