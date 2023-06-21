package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Welcome to Gastroguru")
}
