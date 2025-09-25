package structs

type Sale struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	BookId   int    `json:"book_id"`
	SaleDate string `json:"sale_date"`
}

type SaleInput struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
