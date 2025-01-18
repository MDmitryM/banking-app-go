package service

import (
	"time"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatisticService struct {
	repo repository.Statistic
}

func NewStatisticService(repo repository.Statistic) *StatisticService {
	return &StatisticService{
		repo: repo,
	}
}

func (s *StatisticService) GetMonthlyStatistic(userID, month string) (*bankingApp.MonthlyStatistics, error) {
	parsedMonth, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, err
	}

	startDate := time.Date(parsedMonth.Year(), parsedMonth.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	return s.repo.GetMonthlyStatistic(userObjID, startDate, endDate)
}
