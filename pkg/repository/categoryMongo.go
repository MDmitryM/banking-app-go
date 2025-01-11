package repository

import (
	"context"
	"errors"

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
