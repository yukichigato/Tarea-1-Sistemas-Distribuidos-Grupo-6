package structs

import (
	"time"
)

type Transaction struct {
	TransactionId   int       `json:"transaction_id"`
	BookId          int       `json:"book_id"`
	BookName        string    `json:"book_name"`
	TransactionType string    `json:"transaction_type"`
	TransactionDate time.Time `json:"transaction_date"`
	BookPrice       int       `json:"book_price"`
}
