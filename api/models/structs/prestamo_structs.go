package structs

import (
	"time"
)

type Loan struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	StartDate  time.Time `json:"start_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
}

type LoanStatusUpdate struct {
	Status string `json:"status"`
}

type LoanInput struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
