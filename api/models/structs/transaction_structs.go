package structs

import (
	"time"
)

type Transaction struct {
	BookName        string    `json:"book_name"`
	BookId          int       `json:"book_id"`
	UserName        string    `json:"user_name"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	BookPrice       int       `json:"book_price"`
}
