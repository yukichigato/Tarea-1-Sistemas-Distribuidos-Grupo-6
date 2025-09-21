package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	SERVER_ADDRESS = "http://localhost:8080"
)

// CalculateTotal calculates the total price for a list of book IDs.
// It fetches each book's price and sums them up, ignoring any books that can't be retrieved.
func CalculateTotal(ids []int) int {
	if len(ids) == 0 {
		return 0
	}
	total := 0
	for _, id := range ids {
		book, err := GetBookByID(id)
		if err == nil {
			total += book.Price
		}
	}
	return total
}

// RegisterUser sends a POST request to register a new user on the server.
// It marshals the user data to JSON and sends it to the /user endpoint.
// Returns true if registration is successful (HTTP 201), false otherwise.
func RegisterUser(user User) bool {
	payload, err := json.Marshal(user)
	if err != nil {
		return false
	}

	client, err := http.Post(SERVER_ADDRESS + "/user", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return false
	}
	defer client.Body.Close()

	return client.StatusCode == http.StatusCreated
}

// LoginUser authenticates a user by sending a GET request with email and password.
// It retrieves user data from the server and returns a pointer to the User struct.
// Returns the user data and nil error on success, nil and error on failure.
func LoginUser(email, password string) (*User, error) {
	response, err := http.Get(SERVER_ADDRESS + "/user?email=" + email + "&password=" + password)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return nil, err
	}
	
	return &user, nil
}

// GetBooks retrieves all books from the server.
// Returns a slice of Book structs and an error if the request fails.
func GetBooks() ([]Book, error) {
	response, err := http.Get(SERVER_ADDRESS + "/books")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var books []Book
	if err := json.NewDecoder(response.Body).Decode(&books); err != nil {
		return nil, err
	}
	
	return books, nil
}

// GetBookByID retrieves a specific book by its ID from the server.
// Returns the Book struct and an error if the book is not found or request fails.
func GetBookByID(id int) (Book, error) {
	response, err := http.Get(SERVER_ADDRESS + "/books/" + strconv.Itoa(id))
	if err != nil {
		return Book{}, err
	}
	defer response.Body.Close()

	var book Book
	if err := json.NewDecoder(response.Body).Decode(&book); err != nil {
		return Book{}, err
	}
	
	return book, nil
}

// GetUserLoans retrieves all loans for a specific user from the server.
// Returns a slice of Loan structs and an error if the request fails.
func GetUserLoans(userID int) ([]Loan, error) {
	response, err := http.Get(SERVER_ADDRESS + "/loans?user_id=" + strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var loans []Loan
	if err := json.NewDecoder(response.Body).Decode(&loans); err != nil {
		return nil, err
	}
	
	return loans, nil
}

// GetUserTransactions retrieves all transactions for a specific user from the server.
// Returns a slice of Transaction structs and an error if the request fails.
func GetUserTransactions(userID int) ([]Transaction, error) {
	response, err := http.Get(SERVER_ADDRESS + "/transactions?user_id=" + strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var transactions []Transaction
	if err := json.NewDecoder(response.Body).Decode(&transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func UpdateUserPesos(userID int, newPesos int) error {
	payload := map[string]int{"usm_pesos": newPesos}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, SERVER_ADDRESS+"/user/"+strconv.Itoa(userID), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}