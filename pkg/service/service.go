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
	}
}
