package handler

import (
	"net/http"
	"strconv"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type addTransactionResponce struct {
	TransactionID string `json:"transaction_id"`
}

// @Summary Add Transaction
// @Description Добавить транзакцию пользователя
// @Tags Transactions
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body bankingApp.Transaction true "Transaction details"
// @Success 200 {object} addTransactionResponce "Успешное добавление транзакции, возвращается ID тр. в БД"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unautharized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router  /api/transactions [post]
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

// @Summary Get User Transaction
// @Description Получить транзакций пользователя по страницам
// @Tags Transactions
// @Security ApiKeyAuth
// @Produce json
// @Param page query string false "Page param"
// @Param pageSize query string false "Page Size param"
// @Success 200 {array}  bankingApp.Transaction "Список транзакций постранично"
// @Failure 401 {object} ErrorResponse "Unautharized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router  /api/transactions [get]
func (h *Handler) getTransactions(ctx echo.Context) error {
	pageParam := ctx.QueryParam("page")
	pageSizeParam := ctx.QueryParam("pageSize")

	//default params
	page := 1
	pageSize := 5

	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	if pageSizeParam != "" {
		if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 {
			pageSize = ps
		}
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	transactions, err := h.service.Transaction.GetTransactions(userID, page, pageSize)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	if transactions == nil {
		transactions = []bankingApp.Transaction{}
	}

	return ctx.JSON(http.StatusOK, transactions)
}

// @Summary Update Transaction
// @Description Обновить данные о транзакции пользователя
// @Tags Transactions
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body bankingApp.Transaction true "Transaction details"
// @Param id path string true "ID транзакции"
// @Success 200 {object} bankingApp.Transaction "Успешное обновление данных транзакции, возвращается обновленная запись из БД"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unautharized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router  /api/transactions/{id} [put]
func (h *Handler) updateTransaction(ctx echo.Context) error {
	var trInput bankingApp.Transaction
	if err := ctx.Bind(&trInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(trInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	transactionID := ctx.Param("id")

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId

	updatedTransaction, err := h.service.Transaction.UpdateTransaction(userId, transactionID, trInput)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, updatedTransaction)
}

// @Summary Delete Transaction
// @Description Удалить транзакцию пользователя по ID
// @Tags Transactions
// @Security ApiKeyAuth
// @Param   id path string true "ID транзакции"
// @Success 200 "Успешное обновление данных транзакции, возвращается обновленная запись из БД"
// @Failure 401 {object} ErrorResponse "Unautharized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router  /api/transactions/{id} [delete]
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
