package service

import "github.com/MDmitryM/banking-app-go/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
