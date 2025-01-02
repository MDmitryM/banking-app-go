package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) addCategory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "post /categories",
	})
}

func (h *Handler) getCategories(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get categories list /categories",
	})
}

func (h *Handler) updateCategory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "put /categories/id",
	})
}

func (h *Handler) deleteCategory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpint": "delete /categories/id",
	})
}
