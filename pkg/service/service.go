package service

import (
	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user bankingApp.User) (string, error)
	GenerateToken(email, password string) (string, error)
}

type Transaction interface {
	CreateTransaction(userID string, transactionInput bankingApp.Transaction) (string, error)
	DeleteTransaction(userID, transactionID string) error
	UpdateTransaction(userID, transactionID string, transactionInput bankingApp.Transaction) (bankingApp.Transaction, error)
	GetTransactions(userID string, page, pageSize int) ([]bankingApp.Transaction, error)
}

type Statistic interface {
	GetMonthlyStatistic(userID, month string) (*bankingApp.MonthlyStatistics, error)
}

type Category interface {
	CreateCategory(userID string, categoryInput bankingApp.Category) (string, error)
	GetUserCategories(userID string) ([]bankingApp.Category, error)
	DeleteUserCategory(userID, categoryID string) error
	UpdateCategoryName(userID, categoryID string, updatedCategoryName string) error
}

type CachedStatistic interface {
}

type CachedCategory interface {
	CacheUserCategories(userID string, categories []bankingApp.Category) error
	GetUserCachedCategories(userID string) ([]bankingApp.Category, error)
	InvalidateUserCache(userID string) error
}

type Service struct {
	Authorization
	Transaction
	Statistic
	Category
	CachedStatistic
	CachedCategory
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repo.Authorization),
		Transaction:    NewTransactionService(repo.Transaction),
		Category:       NewCategoryService(repo.Category),
		Statistic:      NewStatisticService(repo.Statistic),
		CachedCategory: NewCachedCategory(repo.CachedCategory),
	}
}
