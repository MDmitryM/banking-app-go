package handler

import (
	_ "github.com/MDmitryM/banking-app-go/docs"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRouts(echo *echo.Echo) {
	//Swagger
	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	//Authorization
	auth := echo.Group("/auth")     // /auth
	auth.POST("/sign-in", h.signIn) // /auth/sign-in
	auth.POST("/sign-up", h.signUp) // /auth/sign-up

	//for /api/* user must be authorized with JWT
	api := echo.Group("/api", h.JWTMiddleware()) // /api

	//Transactions
	transactions := api.Group("/transactions")       // /api/transactions
	transactions.POST("", h.addTransaction)          // add transaction
	transactions.GET("", h.getTransactions)          // get transactions list with pagination
	transactions.PUT("/:id", h.updateTransaction)    // /api/transactions/:id update transaction info
	transactions.DELETE("/:id", h.deleteTransaction) // /api/transactions/:id delete transaction

	//Statistics
	statistics := api.Group("/statistics")       // /api/statistics
	statistics.GET("/monthly", h.monthStatistic) // data for month

	//Categories
	categories := api.Group("/categories")      // api/categories
	categories.POST("", h.addCategory)          // add category
	categories.GET("", h.getCategories)         // categories list
	categories.PUT("/:id", h.updateCategory)    // change category name
	categories.DELETE("/:id", h.deleteCategory) // delete category
}
