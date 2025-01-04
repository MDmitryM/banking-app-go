package models

import (
	"fmt"
	"strconv"
	"time"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Константы для дефолтной категории
const (
	DefaultCategoryID   = "000000000000000000000000" // Специальный ObjectID для дефолтной категории
	DefaultCategoryName = "Без категории"
)

type TransactionModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Amount      float64            `bson:"amount"`
	Type        string             `bson:"type"`
	CategoryID  primitive.ObjectID `bson:"category_id"`
	Date        time.Time          `bosn:"date"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func ToTransactionModel(dto bankingApp.Transaction, userId string) (TransactionModel, error) {
	userObjID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return TransactionModel{}, fmt.Errorf("invalid user id: %w", err)
	}

	amount, err := strconv.ParseFloat(dto.Amount, 64)
	if err != nil {
		return TransactionModel{}, fmt.Errorf("invalid amount format: %w", err)
	}

	categoryID := dto.CategoryID
	if categoryID == "" {
		categoryID = DefaultCategoryID
	}

	objID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		objID, _ = primitive.ObjectIDFromHex(DefaultCategoryID)
	}

	now := time.Now()

	date := dto.Date
	if date.IsZero() {
		date = now
	}
	return TransactionModel{
		UserID:      userObjID,
		Amount:      amount,
		Type:        dto.Type,
		CategoryID:  objID,
		Date:        now,
		Description: dto.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}
