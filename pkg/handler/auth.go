package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) signIn(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "/sign-in",
	})
}

func (h *Handler) signUp(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "/sign-up",
	})
}
