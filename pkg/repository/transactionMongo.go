package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *TransactionMongo) DeleteTransaction(userObjID, transactionObjID primitive.ObjectID) error {
	transactionCollection := r.db.database.Collection("transactions")

	// Фильтр для удаления
	filter := bson.M{
		"_id":     transactionObjID,
		"user_id": userObjID,
	}

	// Удаляем документ с фильтром
	delResult, err := transactionCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting transaction: %w", err)
	}

	// Проверяем, был ли документ удалён
	if delResult.DeletedCount == 0 {
		return errors.New("transaction not found or not owned by the user")
	}

	return nil
}
