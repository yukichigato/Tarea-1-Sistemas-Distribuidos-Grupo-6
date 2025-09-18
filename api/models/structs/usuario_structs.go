package structs

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UsmPesos  int    `json:"usm_pesos"`
}

type UserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserBalanceUpdate struct {
	BookList []int `json:"books_id_list"` // Para libros adquiridos
	Deposit  *int  `json:"deposit"`       // Para abonar usm pesos
	LateFee  *int  `json:"late_fee"`      // Para devoluciones tardías
}
