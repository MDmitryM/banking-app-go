package repository

import (
	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Authorization interface {
	CreateUser(user models.UserModel) (string, error)
	IsUserValid(email, password string) (string, error)
}

type Transaction interface {
	CreateTransaction(transaction models.TransactionModel) (string, error)
	DeleteTransaction(userID, transactionID primitive.ObjectID) error
}

type Statistic interface {
}

type Category interface {
}

type Repository struct {
	Authorization
	Transaction
	Statistic
	Category
}

func NewRepository(db *MongoDB) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		Transaction:   NewTransactionMongo(db),
	}
}
