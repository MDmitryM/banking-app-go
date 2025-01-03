package handler

import (
	"net/http"

	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) addTransaction(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "post /transactions " + userId,
	})
}

func (h *Handler) getTransactions(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get /transactions " + userId,
	})
}

func (h *Handler) updateTransaction(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "put /transactions/id " + userId,
	})
}

func (h *Handler) deleteTransaction(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "delete /transactions/id " + userId,
	})
}
