package handler

import (
	"net/http"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type addTransactionResponce struct {
	TransactionID string `json:"transaction_id"`
}

func (h *Handler) addTransaction(ctx echo.Context) error {
	var transactionInput bankingApp.Transaction
	if err := ctx.Bind(&transactionInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(transactionInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId

	transactionId, err := h.service.Transaction.CreateTransaction(userId, transactionInput)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, addTransactionResponce{
		TransactionID: transactionId,
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
	transactionID := ctx.Param("id")

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	err := h.service.Transaction.DeleteTransaction(userID, transactionID)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
