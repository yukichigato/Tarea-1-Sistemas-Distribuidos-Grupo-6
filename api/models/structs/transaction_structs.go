package structs

type Transaction struct {
	TransactionId   int    `json:"transaction_id"`
	BookId          int    `json:"book_id"`
	BookName        string `json:"book_name"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
	BookPrice       int    `json:"book_price"`
}
