package models

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	UsmPesos  int
}

type UserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UsmPesos  int    `json:"usm_pesos"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	UsmPesos int `json:"usm_pesos"`
}
