package handler

import (
	"net/http"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/labstack/echo/v4"
)

type singUpResponce struct {
	Id          string `json:"id"`
	AccessToken string `json:"access_token"`
}

func (h *Handler) signUp(ctx echo.Context) error {
	var input bankingApp.User
	if err := ctx.Bind(&input); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(input); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	userId, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	token, err := h.service.GenerateToken(input.Email, input.Password)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, singUpResponce{
		Id:          userId,
		AccessToken: token,
	})
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponce struct {
	AccessToken string `json:"acess_token"`
}

func (h *Handler) signIn(ctx echo.Context) error {
	var input signInInput
	if err := ctx.Bind(&input); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(input); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	token, err := h.service.GenerateToken(input.Email, input.Password)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, signInResponce{
		AccessToken: token,
	})
}
