package handler

import (
	"net/http"
	"os"

	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h *Handler) JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &service.JwtBankingClaims{}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Invalid or missing token",
			})
		},
	})
}
