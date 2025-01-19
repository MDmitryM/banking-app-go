package service

import (
	"encoding/json"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
)

type StatisticCachedService struct {
	repo repository.CachedStatistic
}

func NewStatisticCachedService(repo repository.CachedStatistic) *StatisticCachedService {
	return &StatisticCachedService{
		repo: repo,
	}
}

func (s *StatisticCachedService) CacheUserStatistic(userID, month string, stats *bankingApp.MonthlyStatistics) error {
	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	return s.repo.CacheUserStatistic(userID, month, string(data))
}

func (s *StatisticCachedService) GetUserCachedStatistic(userID, month string) (*bankingApp.MonthlyStatistics, error) {
	data, err := s.repo.GetUserCachedStatistic(userID, month)
	if err != nil {
		return nil, err
	}

	var stats bankingApp.MonthlyStatistics
	if err := json.Unmarshal([]byte(data), &stats); err != nil {
		return nil, err
	}

	return &stats, nil
}

func (s *StatisticCachedService) DeleteCachedStatisticByMonth(userID, month string) error {
	return s.repo.DeleteCachedStatisticByMonth(userID, month)
}

func (s *StatisticCachedService) InvalidateUserStatisticCache(userID string) error {
	return s.repo.DeleteAllUserCachedStatistics(userID)
}
