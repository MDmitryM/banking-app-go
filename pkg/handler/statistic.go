package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) monthStatistic(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/monthly",
	})
}

func (h *Handler) categotyStatistic(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/category",
	})
}

func (h *Handler) trendStatistic(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/trends",
	})
}
