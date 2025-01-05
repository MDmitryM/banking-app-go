package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *TransactionMongo) GetTransactionByID(transactionID primitive.ObjectID) (models.TransactionModel, error) {
	transactionCollection := r.db.database.Collection("transactions")

	var transaction models.TransactionModel
	err := transactionCollection.FindOne(context.Background(), bson.M{"_id": transactionID}).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.TransactionModel{}, errors.New("transaction not found or not owned by the user")
		}
		return models.TransactionModel{}, fmt.Errorf("failed to get transaction: %w", err)
	}

	return transaction, nil
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

func (r *TransactionMongo) UpdateTransaction(transactionObjID primitive.ObjectID, trModelToUpdate models.TransactionModel) (models.TransactionModel, error) {
	transactionCollection := r.db.database.Collection("transactions")

	existingTransaction, err := r.GetTransactionByID(transactionObjID)
	if err != nil {
		return models.TransactionModel{}, err
	}

	if existingTransaction.UserID != trModelToUpdate.UserID {
		return models.TransactionModel{}, errors.New("operation not allowed")
	}

	trModelToUpdate.CreatedAt = existingTransaction.CreatedAt
	trModelToUpdate.ID = transactionObjID
	trModelToUpdate.UpdatedAt = time.Now()

	var updatedTransaction models.TransactionModel
	err = transactionCollection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": trModelToUpdate.ID},
		bson.M{"$set": trModelToUpdate},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedTransaction)
	if err != nil {
		return models.TransactionModel{}, nil
	}
	return updatedTransaction, nil
}

func (r *TransactionMongo) GetTransactions(userID primitive.ObjectID, offset, limit int) ([]models.TransactionModel, error) {
	transactionCollection := r.db.database.Collection("transactions")

	cursor, err := transactionCollection.Find(context.Background(),
		bson.M{"user_id": userID},
		options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactionModels []models.TransactionModel
	if err := cursor.All(context.Background(), &transactionModels); err != nil {
		return nil, err
	}

	return transactionModels, nil
}
