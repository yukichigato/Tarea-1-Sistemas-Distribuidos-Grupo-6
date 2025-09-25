package structs

type Loan struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	BookId     int    `json:"book_id"`
	StartDate  string `json:"start_date"`
	ReturnDate string `json:"return_date"`
	Status     string `json:"status"`
}

type LoanInput struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
