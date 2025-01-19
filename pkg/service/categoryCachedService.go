package service

import (
	"encoding/json"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
)

type CategoryCachedService struct {
	repo repository.CachedCategory
}

func NewCachedCategory(repo repository.CachedCategory) *CategoryCachedService {
	return &CategoryCachedService{
		repo: repo,
	}
}

func (s *CategoryCachedService) GetUserCachedCategories(userID string) ([]bankingApp.Category, error) {
	cachedData, err := s.repo.GetUserCachedCategories(userID)
	if err == nil {
		var userCategories []bankingApp.Category
		if err := json.Unmarshal([]byte(cachedData), &userCategories); err != nil {
			return nil, err
		}
		return userCategories, nil
	}
	return nil, err
}

func (s *CategoryCachedService) CacheUserCategories(userID string, categories []bankingApp.Category) error {
	data, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	return s.repo.CacheUserCategories(userID, string(data))
}

func (s *CategoryCachedService) InvalidateUserCache(userID string) error {
	return s.repo.DeleteUserCachedCategories(userID)
}
