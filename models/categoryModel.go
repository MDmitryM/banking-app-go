package models

import (
	bankingApp "github.com/MDmitryM/banking-app-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       primitive.ObjectID `bson:"user_id"`
	CategoryName string             `bson:"category_name"`
	CategoryType string             `bson:"category_type"`
}

func ToCategoryModel(userID string, categoryInput bankingApp.Category) (CategoryModel, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return CategoryModel{}, err
	}

	return CategoryModel{
		UserID:       userObjID,
		CategoryName: categoryInput.CategoryName,
		CategoryType: categoryInput.CategoryType,
	}, nil
}

func (m *CategoryModel) ToCategoryDTO() bankingApp.Category {
	return bankingApp.Category{
		ID:           m.ID.Hex(),
		CategoryName: m.CategoryName,
		CategoryType: m.CategoryType,
	}
}
