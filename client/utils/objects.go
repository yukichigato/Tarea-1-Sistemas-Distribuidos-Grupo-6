package utils

type User struct {
	ID       	int    `json:"id"`
	Name     	string `json:"first_name"`
	Surname  	string `json:"last_name"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	USM_Pesos int    `json:"usm_pesos"`
}
	
type Book struct {
	ID                int    `json:"id"`
	Title             string `json:"book_name"`
	Category          string `json:"book_category"`
	Transaction_type  string `json:"transaction_type"`
	Price             int    `json:"price"`
	Status            string `json:"status"`
	Popularity        int    `json:"popularity_score"`
}

type Inventory struct {
	ID      						int `json:"id"`
	Available_quantity 	int `json:"quantity"`
}

type Loan struct {
	ID                int    `json:"id"`
	User_ID           int    `json:"user_id"`
	Book_ID           int    `json:"book_id"`
	Start_date        string `json:"start_date"`
	Return_date       string `json:"return_date"`
	Status            string `json:"status"`
}

type Sale struct {
	ID 							int    `json:"id"`
	User_ID           int    `json:"user_id"`
	Book_ID           int    `json:"book_id"`
	Sale_date         string `json:"sale_date"`
}

type Transaction struct {
	ID								int    `json:"id"`
	Book_ID           int    `json:"book_id"`
	Title						 	string `json:"book_name"`
	Type 							string `json:"transaction_type"`
	Transaction_date 	string `json:"transaction_date"`
	Price             int    `json:"price"`
}