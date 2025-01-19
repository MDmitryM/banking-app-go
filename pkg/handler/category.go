package handler

import (
	"net/http"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/repository"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type addCategoryResponce struct {
	CategoryID string `json:"category_id"`
}

// @Summary     Add category
// @Tags        Categories
// @Description Создание пользовательской категории
// @Security	ApiKeyAuth
// @Accept 		json
// @Produce 	json
// @Param input body bankingApp.Category true "Category details"
// @Success		200  {object} addCategoryResponce "Успешное создание категории"
// @Failure     400  {object} ErrorResponse "Invalid body"
// @Failure		401  {object} ErrorResponse	"Unauthorized"
// @Failure	    500  {object} ErrorResponse "Internal server error"
// @Router		/api/categories [post]
func (h *Handler) addCategory(ctx echo.Context) error {
	var catInput bankingApp.Category
	if err := ctx.Bind(&catInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(catInput); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	catID, err := h.service.Category.CreateCategory(userID, catInput)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, addCategoryResponce{
		CategoryID: catID,
	})
}

// @Summary 	Get categories
// @Description Получения всех категорий пользователя
// @Tags 		Categories
// @Security	ApiKeyAuth
// @Produce		json
// @Success		200  {array}    bankingApp.Category "Список категорй пользователя"
// @Failure		401  {object}   ErrorResponse	    "Unauthorized"
// @Failure	    500  {object}   ErrorResponse       "Internal server error"
// @Router		/api/categories [get]
func (h *Handler) getCategories(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	cachedCategories, err := h.service.CachedCategory.GetUserCachedCategories(userID)
	if err == nil {
		return ctx.JSON(http.StatusOK, cachedCategories)
	}
	if err == repository.ErrCacheNotFound {
		logrus.Print(err.Error())
	}

	userCategories, err := h.service.Category.GetUserCategories(userID)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	if userCategories == nil {
		userCategories = []bankingApp.Category{}
	} else {
		h.service.CachedCategory.CacheUserCategories(userID, userCategories)
	}

	return ctx.JSON(http.StatusOK, userCategories)
}

type CategoryNameInput struct {
	UpdatedName string `json:"category_name" validate:"required"`
}

// @Summary 	Update category name
// @Description Обновление названия категории пользователя
// @Tags 		Categories
// @Security	ApiKeyAuth
// @Accept 		json
// @Produce		json
// @Param		category 	body 		CategoryNameInput	true  "Новое название категории"
// @Param		id			path		string				true  "ID категории у которой изменяется имя"
// @Success		200  											  "Обновлено успешно"
// @Failure		401  		{object}   	ErrorResponse	    "Unauthorized"
// @Failure	    500  		{object}   	ErrorResponse       "Internal server error"
// @Router		/api/categories/{id} [put]
func (h *Handler) updateCategory(ctx echo.Context) error {
	categoryID := ctx.Param("id")
	if categoryID == "" {
		return SendJSONError(ctx, http.StatusBadRequest, "invalid category id")
	}

	var categoryDTO CategoryNameInput
	if err := ctx.Bind(&categoryDTO); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(categoryDTO); err != nil {
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userId := claims.UserId

	if err := h.service.Category.UpdateCategoryName(userId, categoryID, categoryDTO.UpdatedName); err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

// @Summary 	Delete category
// @Description Удаление категории пользователя
// @Tags 		Categories
// @Security	ApiKeyAuth
// @Produce		json
// @Param		id			path		string		  true  "ID удаляемой категории"
// @Success		200  										"Удаление успешно"
// @Failure		400 		{object}	ErrorResponse		"Internal server error"
// @Failure		401  		{object}   	ErrorResponse	    "Unauthorized"
// @Failure	    500  		{object}   	ErrorResponse       "Internal server error"
// @Router		/api/categories/{id} [delete]
func (h *Handler) deleteCategory(ctx echo.Context) error {
	categoryID := ctx.Param("id")
	if categoryID == "" {
		return SendJSONError(ctx, http.StatusBadRequest, "invalid category id")
	}

	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	err := h.service.Category.DeleteUserCategory(userID, categoryID)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
