package bankingApp

type Category struct {
	CategoryName string `json:"category_name" validate:"required"`
	CategoryType string `json:"category_type" validate:"required,oneof=income expence"`
}
