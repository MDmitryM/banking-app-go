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
}

type Statistic interface {
}

type Category interface {
}

type Service struct {
	Authorization
	Transaction
	Statistic
	Category
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Transaction:   NewTransactionService(repo.Transaction),
	}
}
