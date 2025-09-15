package models

type Inventory struct {
	AvailableQuantity int `json:"available_quantity"`
}

type Book struct {
	Id              int       `json:"id"`
	BookName        string    `json:"book_name"`
	BookCategory    string    `json:"book_category"`
	TransactionType string    `json:"transaction_type"`
	Price           float64   `json:"price"`
	Status          string    `json:"status"`
	PopularityScore int       `json:"popularity_score"`
	Inventory       Inventory `json:"inventory"`
}

type BookInput struct {
	BookName        string    `json:"book_name"`
	BookCategory    string    `json:"book_category"`
	TransactionType string    `json:"transaction_type"`
	Price           float64   `json:"price"`
	Status          string    `json:"status"`
	PopularityScore int       `json:"popularity_score"`
	Inventory       Inventory `json:"inventory"`
}
