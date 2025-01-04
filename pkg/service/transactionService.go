package service

import (
	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateTransaction(userID string, transactionInput bankingApp.Transaction) (string, error) {
	trModel, err := models.ToTransactionModel(transactionInput, userID)
	if err != nil {
		return "", err
	}
	return s.repo.CreateTransaction(trModel)
}
