package bankingApp

import "time"

type Transaction struct {
	ID          string    `json:"id,omitempty"`
	Amount      string    `json:"amount" validate:"required"`
	Type        string    `json:"type" validate:"required,oneof=income expence"`
	CategoryID  string    `json:"category_id"`
	Date        time.Time `json:"time" validate:"required"`
	Description string    `json:"description"`
}
