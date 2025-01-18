package bankingApp

// Структуры для ответа
type MonthlyStatistics struct {
	Month        string           `json:"month"` // формат "2024-01"
	TotalIncome  float64          `json:"total_income"`
	TotalExpense float64          `json:"total_expense"`
	Balance      float64          `json:"balance"`
	Categories   []CategoryAmount `json:"categories"`
}

type CategoryAmount struct {
	CategoryID   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	CategoryType string  `json:"category_type"`
	Amount       float64 `json:"amount"`
}
