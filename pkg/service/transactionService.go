package service

import (
	"errors"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *TransactionService) DeleteTransaction(userID, transactionID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id format")
	}

	transactionObjID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return errors.New("invalid transaction id format")
	}

	return s.repo.DeleteTransaction(userObjID, transactionObjID)
}
