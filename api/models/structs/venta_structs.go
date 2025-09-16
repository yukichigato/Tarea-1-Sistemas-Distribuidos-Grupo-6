package structs

import (
	"time"
)

type Sale struct {
	Id       int       `json:"id"`
	UserId   int       `json:"user_id"`
	BookId   int       `json:"book_id"`
	SaleDate time.Time `json:"sale_date"`
}

type SaleUpdate struct {
	UserId   int       `json:"user_id"`
	BookId   int       `json:"book_id"`
	SaleDate time.Time `json:"sale_date"`
}

type SaleInput struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}
