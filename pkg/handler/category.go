package handler

import (
	"net/http"

	bankingApp "github.com/MDmitryM/banking-app-go"
	"github.com/MDmitryM/banking-app-go/pkg/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type addCategoryResponce struct {
	CategoryID string `json:"category_id"`
}

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

func (h *Handler) getCategories(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*service.JwtBankingClaims)
	userID := claims.UserId

	userCategories, err := h.service.Category.GetUserCategories(userID)
	if err != nil {
		return SendJSONError(ctx, http.StatusInternalServerError, err.Error())
	}

	if userCategories == nil {
		userCategories = []bankingApp.Category{}
	}

	return ctx.JSON(http.StatusOK, userCategories)
}

type CategoryNameInput struct {
	UpdatedName string `json:"category_name" validate:"required"`
}

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
		return SendJSONError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}

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
