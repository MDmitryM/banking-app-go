package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ErrorResponse представляет структуру ответа с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse создает новый ErrorResponse
func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

// SendError отправляет ошибку клиенту
func SendJSONError(c echo.Context, statusCode int, message string) error {
	logrus.Error(message)
	return c.JSON(statusCode, NewErrorResponse(message))
}
