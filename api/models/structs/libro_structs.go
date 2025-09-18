package structs

type Book struct {
	Id              int    `json:"id"`
	BookName        string `json:"book_name"`
	BookCategory    string `json:"book_category"`
	TransactionType string `json:"transaction_type"`
	Price           int    `json:"price"`
	Status          string `json:"status"`
	PopularityScore int    `json:"popularity_score"`
}

type BookInput struct {
	BookName        string `json:"book_name"`
	BookCategory    string `json:"book_category"`
	TransactionType string `json:"transaction_type"`
	Price           int    `json:"price"`
	Status          string `json:"status"`
}
