package handler

import (
	"net/http"
	"time"

	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) monthStatistic(ctx echo.Context) error {
	month := ctx.QueryParam("month")
	if month == "" {
		month = time.Now().Format("2006-01")
	}

	_, err := time.Parse("2006-01", month)
	if err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, "invalid month format. Use YYYY-MM")
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId

	stats, err := h.service.Statistic.GetMonthlyStatistic(userId, month)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, stats)
}
