package handler

import (
	"net/http"

	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) addCategory(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"user_id": "post /categories " + userId,
	})
}

func (h *Handler) getCategories(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "get categories list /categories " + userId,
	})
}

func (h *Handler) updateCategory(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpoint": "put /categories/id " + userId,
	})
}

func (h *Handler) deleteCategory(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId
	return ctx.JSON(http.StatusOK, echo.Map{
		"endpint": "delete /categories/id " + userId,
	})
}
