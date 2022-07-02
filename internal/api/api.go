package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "working")
}
