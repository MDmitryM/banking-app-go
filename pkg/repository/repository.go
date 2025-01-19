package repository

import (
	"errors"
	"time"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrCacheNotFound = errors.New("cache not found")

type Authorization interface {
	CreateUser(user models.UserModel) (string, error)
	IsUserValid(email, password string) (string, error)
}

type Transaction interface {
	CreateTransaction(transaction models.TransactionModel) (string, error)
	DeleteTransaction(userID, transactionID primitive.ObjectID) (time.Time, error)
	UpdateTransaction(transactionID primitive.ObjectID, updatedTransactionModel models.TransactionModel) (models.TransactionModel, error)
	GetTransactionByID(transactioID primitive.ObjectID) (models.TransactionModel, error)
	GetTransactions(userID primitive.ObjectID, offset, limit int) ([]models.TransactionModel, error)
}

type Statistic interface {
	GetMonthlyStatistic(userID primitive.ObjectID, startDate, endDate time.Time) (*bankingApp.MonthlyStatistics, error)
}

type Category interface {
	CreateCategory(categoryToCreate models.CategoryModel) (string, error)
	GetUserCategories(userID primitive.ObjectID) ([]models.CategoryModel, error)
	DeleteUserCategory(userObjID, categoryObjID primitive.ObjectID) error
	UpdateCategoryName(userObjID, categoryObjID primitive.ObjectID, updated string) error
}

type CachedStatistic interface {
	CacheUserStatistic(userID, month, stats string) error
	GetUserCachedStatistic(userID, month string) (string, error)
	DeleteCachedStatisticByMonth(userID, month string) error
	DeleteAllUserCachedStatistics(userID string) error
}

type CachedCategory interface {
	CacheUserCategories(userID, data string) error
	GetUserCachedCategories(userID string) (string, error)
	DeleteUserCachedCategories(userId string) error
}

type Repository struct {
	Authorization
	Transaction
	Statistic
	Category
	CachedStatistic
	CachedCategory
}

func NewRepository(db *MongoDB, redisDb *RedisDB) *Repository {
	return &Repository{
		Authorization:   NewAuthMongo(db),
		Transaction:     NewTransactionMongo(db),
		Category:        NewCategoryMongo(db),
		Statistic:       NewStatisticMongo(db),
		CachedCategory:  NewCategoryRedis(redisDb),
		CachedStatistic: NewStatisticRedis(redisDb),
	}
}
