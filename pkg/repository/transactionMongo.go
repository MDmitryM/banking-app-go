package repository

import (
	"context"

	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionMongo struct {
	db *MongoDB
}

func NewTransactionMongo(db *MongoDB) *TransactionMongo {
	return &TransactionMongo{
		db: db,
	}
}

func (r *TransactionMongo) CreateTransaction(transaction models.TransactionModel) (string, error) {
	transactionCollection := r.db.database.Collection("transactions")

	result, err := transactionCollection.InsertOne(context.Background(), transaction)
	if err != nil {
		return "", err
	}

	transactionID := result.InsertedID.(primitive.ObjectID).Hex()
	return transactionID, nil
}
