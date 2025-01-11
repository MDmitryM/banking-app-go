package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryMongo struct {
	db *MongoDB
}

func NewCategoryMongo(db *MongoDB) *CategoryMongo {
	return &CategoryMongo{
		db: db,
	}
}

func (r *CategoryMongo) CreateCategory(categoryToCreate models.CategoryModel) (string, error) {
	categoriesCollection := r.db.database.Collection("categories")

	filter := bson.M{
		"user_id":       categoryToCreate.UserID,
		"category_name": categoryToCreate.CategoryName,
	}

	var categoryByName models.CategoryModel
	err := categoriesCollection.FindOne(context.Background(), filter).Decode(&categoryByName)
	if err == nil {
		return "", errors.New("category already exists")
	}

	if err != mongo.ErrNoDocuments {
		return "", err
	}

	result, err := categoriesCollection.InsertOne(context.Background(), categoryToCreate)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *CategoryMongo) GetUserCategories(userID primitive.ObjectID) ([]models.CategoryModel, error) {
	categoriesCollection := r.db.database.Collection("categories")

	cursor, err := categoriesCollection.Find(context.Background(),
		bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var categoriesModels []models.CategoryModel
	if err := cursor.All(context.Background(), &categoriesModels); err != nil {
		return nil, err
	}

	return categoriesModels, nil
}

// Transaction numbers are only allowed on a replica set member or mongos
func (r *CategoryMongo) DeleteUserCategory(userObjID, categoryObjID primitive.ObjectID) error {
	// Начинаем сессию
	session, err := r.db.client.StartSession()
	if err != nil {
		return fmt.Errorf("failed to start session: %w", err)
	}
	defer session.EndSession(context.Background())

	// Выполняем транзакцию
	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {
		// Начало транзакции
		if err := session.StartTransaction(); err != nil {
			return fmt.Errorf("failed to start transaction: %w", err)
		}

		// Проверяем существование категории
		categoryCollection := r.db.database.Collection("categories")
		deleteResult, err := categoryCollection.DeleteOne(sessCtx, bson.M{
			"_id":     categoryObjID,
			"user_id": userObjID,
		})
		if err != nil {
			_ = session.AbortTransaction(sessCtx)
			return fmt.Errorf("failed to delete category: %w", err)
		}
		if deleteResult.DeletedCount == 0 {
			_ = session.AbortTransaction(sessCtx)
			return errors.New("category is not found")
		}

		// Обновляем транзакции с удалённой категорией
		transactionCollection := r.db.database.Collection("transactions")
		transactionFilter := bson.M{
			"user_id":     userObjID,
			"category_id": categoryObjID,
		}
		update := bson.M{"$set": bson.M{"category_id": models.DefaultCategoryID}}
		_, err = transactionCollection.UpdateMany(sessCtx, transactionFilter, update)
		if err != nil {
			_ = session.AbortTransaction(sessCtx)
			return fmt.Errorf("failed to update transactions: %w", err)
		}

		// Коммит транзакции
		if err := session.CommitTransaction(sessCtx); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
		return nil
	})

	return err
}
