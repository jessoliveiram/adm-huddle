package main

import (
	"github.com/jessoliveiram/adm-huddle/internal/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/", api.Healthcheck)

	e.Start(":8080")
}
