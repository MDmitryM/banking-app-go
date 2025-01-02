package handler

import (
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRouts(echo *echo.Echo) {
	//Authorization
	auth := echo.Group("/auth")     // /auth
	auth.POST("/sign-in", h.signIn) // /auth/sign-in
	auth.POST("/sign-up", h.signUp) // /auth/sign-up

	//for /api/* user must be authorized with JWT
	api := echo.Group("/api") // /api

	//Transactions
	transactions := api.Group("/transactions")       // /api/transactions
	transactions.POST("/", h.addTransaction)         // add transaction
	transactions.GET("/", h.getTransactions)         // get transactions list with pagination
	transactions.PUT("/:id", h.updateTransaction)    // /api/transactions/:id update transaction info
	transactions.DELETE("/:id", h.deleteTransaction) // /api/transactions/:id delete transaction

	//Statistics
	statistics := api.Group("/statistics")           // /api/statistics
	statistics.GET("/monthly", h.monthStatistic)     // data for month
	statistics.GET("/category", h.categotyStatistic) // data for category
	statistics.GET("/trends", h.trendStatistic)      // trends in user's income/expenses

	//Categories
	categories := api.Group("/categories")        // api/categories
	categories.POST("/", h.addCategory)           // add category
	categories.GET("/", h.getCategories)          // categories list
	categories.PUT("/:name", h.updateCategory)    // change category name
	categories.DELETE("/:name", h.deleteCategory) // delete category
}
