package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) addTransaction(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "post /transactions",
	})
}

func (h *Handler) getTransactions(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get /transactions",
	})
}

func (h *Handler) updateTransaction(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "put /transactions/id",
	})
}

func (h *Handler) deleteTransaction(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "delete /transactions/id",
	})
}
