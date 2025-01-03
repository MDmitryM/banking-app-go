package handler

import (
	"net/http"

	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) monthStatistic(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/monthly " + userId,
	})
}

func (h *Handler) categotyStatistic(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/category " + userId,
	})
}

func (h *Handler) trendStatistic(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get statistics/trends " + userId,
	})
}
