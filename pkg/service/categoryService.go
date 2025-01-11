package service

import (
	"errors"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/models"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) CreateCategory(userID string, categoryInput bankingApp.Category) (string, error) {
	categoryModel, err := models.ToCategoryModel(userID, categoryInput)
	if err != nil {
		return "", err
	}

	createdCategoryID, err := s.repo.CreateCategory(categoryModel)
	if err != nil {
		return "", err
	}

	return createdCategoryID, nil
}

func (s *CategoryService) GetUserCategories(userID string) ([]bankingApp.Category, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	categoriesModels, err := s.repo.GetUserCategories(userObjID)
	if err != nil {
		return nil, err
	}

	var categoriesDTO []bankingApp.Category
	for _, v := range categoriesModels {
		categoriesDTO = append(categoriesDTO, v.ToCategoryDTO())
	}

	return categoriesDTO, nil
}

func (s *CategoryService) DeleteUserCategory(userID, categoryID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	categoryObjID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return errors.New("invalid category ID")
	}

	return s.repo.DeleteUserCategory(userObjID, categoryObjID)
}
