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
	Books   []int `json:"books"`   // Para libros adquiridos
	Penalty int   `json:"penalty"` // Para devoluciones tardías
}
