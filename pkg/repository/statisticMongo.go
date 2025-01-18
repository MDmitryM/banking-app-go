package repository

import (
	"context"
	"fmt"
	"time"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticMongo struct {
	db *MongoDB
}

func NewStatisticMongo(db *MongoDB) *StatisticMongo {
	return &StatisticMongo{
		db: db,
	}
}

func (r *StatisticMongo) GetMonthlyStatistic(userObjID primitive.ObjectID, startDate, endDate time.Time) (*bankingApp.MonthlyStatistics, error) {
	transactionCollection := r.db.database.Collection("transactions")

	pipeline := mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "user_id", Value: userObjID},
				{Key: "date", Value: bson.D{
					{Key: "$gte", Value: startDate},
					{Key: "$lte", Value: endDate},
				}},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "categories"},
				{Key: "localField", Value: "category_id"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "category"},
			}},
		},
		bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$category"},
				{Key: "preserveNullAndEmptyArrays", Value: true}, // Сохраняем транзакции без категорий
			}},
		},
		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "category", Value: bson.D{
					{Key: "$cond", Value: bson.D{
						{Key: "if", Value: bson.D{
							{Key: "$eq", Value: bson.A{bson.D{
								{Key: "$ifNull", Value: bson.A{"$category", nil}},
							}, nil}},
						}},
						{Key: "then", Value: bson.D{
							{Key: "_id", Value: models.DefaultCategoryID},
							{Key: "category_name", Value: "Без категории"},
							{Key: "category_type", Value: "expense"}, // Предполагаем что транзакции без категории - это расходы
						}},
						{Key: "else", Value: "$category"},
					}},
				}},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "category_id", Value: "$category._id"},
					{Key: "category_name", Value: "$category.category_name"},
					{Key: "category_type", Value: "$category.category_type"},
				}},
				{Key: "amount", Value: bson.D{
					{Key: "$sum", Value: "$amount"},
				}},
			}},
		},
	}

	cursor, err := transactionCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate transactions: %w", err)
	}
	defer cursor.Close(context.Background())

	var result []struct {
		ID struct {
			CategoryID   string `bson:"category_id"`
			CategoryName string `bson:"category_name"`
			CategoryType string `bson:"category_type"`
		} `bson:"_id"`
		Amount float64 `bson:"amount"`
	}

	if err := cursor.All(context.Background(), &result); err != nil {
		return nil, fmt.Errorf("failed to decode results: %w", err)
	}

	stats := &bankingApp.MonthlyStatistics{
		Month:      startDate.Format("2006-01"),
		Categories: make([]bankingApp.CategoryAmount, 0, len(result)),
	}

	for _, r := range result {
		categoryAmount := bankingApp.CategoryAmount{
			CategoryID:   r.ID.CategoryID,
			CategoryName: r.ID.CategoryName,
			CategoryType: r.ID.CategoryType,
			Amount:       r.Amount,
		}
		stats.Categories = append(stats.Categories, categoryAmount)

		if r.ID.CategoryType == "income" {
			stats.TotalIncome += r.Amount
		} else {
			stats.TotalExpense += r.Amount
		}
	}

	stats.Balance = stats.TotalIncome - stats.TotalExpense

	return stats, nil
}
