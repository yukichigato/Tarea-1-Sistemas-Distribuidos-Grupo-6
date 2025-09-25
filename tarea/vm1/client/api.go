package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// BACKEND_URL defines the base URL for all API calls
//const BACKEND_URL = "http://10.10.30.252:8080"
const BACKEND_URL = "http://localhost:8080"
// CreateSale creates a new sale record via HTTP POST request to the sales endpoint
func CreateSale(userID int, bookID int) error {
	sale := Sale{
		User_ID: userID,
		Book_ID: bookID,
	}
	saleJson, err := json.Marshal(sale)
	if err != nil {
		panic(err)
	}

	res, err := http.Post(BACKEND_URL+"/sales", "application/json", bytes.NewBuffer(saleJson))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create sale, status code: %d", res.StatusCode)
	}

	return nil
}

// CreateLoan creates a new loan record via HTTP POST request to the loans endpoint
func CreateLoan(userID int, bookID int) error {
	fmt.Println("Creating loan for user", userID, "and book", bookID)
	loan := Loan{
		User_ID: userID,
		Book_ID: bookID,
	}
	loanJson, err := json.Marshal(loan)
	if err != nil {
		panic(err)
	}

	res, err := http.Post(BACKEND_URL+"/loans", "application/json", bytes.NewBuffer(loanJson))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create loan, status code: %d", res.StatusCode)
	}

	return nil
}

// FetchUser retrieves user information by email and password via HTTP GET request
func FetchUser(email string, password string) (User, error) {
	res, err := http.Get(BACKEND_URL + "/users?email=" + email + "&password=" + password)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var user User
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	return user, nil
}

// CreateUser creates a new user account via HTTP POST request to the users endpoint
func CreateUser(newUser User) error {
	userJson, err := json.Marshal(newUser)
	if err != nil {
		return fmt.Errorf("error marshaling user data: %v", err)
	}

	res, err := http.Post(BACKEND_URL+"/users", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create user, status code: %d", res.StatusCode)
	}

	return nil
}

// FetchBooks retrieves all available books from the server via HTTP GET request
func FetchBooks() ([]Book, error) {
	res, err := http.Get(BACKEND_URL + "/books")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var books []Book
	err = json.NewDecoder(res.Body).Decode(&books)
	if err != nil {
		panic(err)
	}

	return books, nil
}













// FetchBook retrieves a specific book by ID via HTTP GET request
func FetchBook(id int) (Book, error) {
	res, err := http.Get(BACKEND_URL + "/books/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var book Book
	err = json.NewDecoder(res.Body).Decode(&book)
	if err != nil {
		panic(err)
	}

	return book, nil
}

// FetchUserLoans retrieves all loans for a specific user by user ID via HTTP GET request
func FetchUserLoans(id int) ([]Loan, error) {
	res, err := http.Get(BACKEND_URL + "/loans?user_id=" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var loans []Loan
	err = json.NewDecoder(res.Body).Decode(&loans)
	if err != nil {
		panic(err)
	}
	return loans, nil
}

// FetchTransactions retrieves all transactions from the server via HTTP GET request
func FetchTransactions() ([]Transaction, error) {
	res, err := http.Get(BACKEND_URL + "/transactions")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var transactions []Transaction
	err = json.NewDecoder(res.Body).Decode(&transactions)
	if err != nil {
		panic(err)
	}
	return transactions, nil
}

// FetchUserTransactions combines user loans and sales to create a unified transaction history
func FetchUserTransactions(id int) ([]Transaction, error) {
	loans, err := FetchUserLoans(id)
	if err != nil {
		return nil, err
	}

	sales, err := FetchUserSales(id)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, loan := range loans {
		book, err := FetchBook(loan.Book_ID)
		if err != nil {
			return nil, err
		}
		transaction := Transaction{
			ID:               loan.ID,
			Book_ID:          loan.Book_ID,
			Title:            book.Title,
			Type:             "Arriendo",
			Transaction_date: loan.Start_date,
			Price:            book.Price,
		}
		transactions = append(transactions, transaction)
	}

	for _, sale := range sales {
		book, err := FetchBook(sale.Book_ID)
		if err != nil {
			return nil, err
		}
		transaction := Transaction{
			ID:               sale.ID,
			Book_ID:          sale.Book_ID,
			Title:            book.Title,
			Type:             "Venta",
			Transaction_date: sale.Sale_date,
			Price:            book.Price,
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// FetchUserSales retrieves all sales for a specific user by user ID via HTTP GET request
func FetchUserSales(id int) ([]Sale, error) {
	res, err := http.Get(BACKEND_URL + "/sales?user_id=" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var sales []Sale
	err = json.NewDecoder(res.Body).Decode(&sales)
	if err != nil {
		panic(err)
	}
	return sales, nil
}

// UpdateUserUSMPesos updates a user's USM_Pesos balance via HTTP PATCH request
func UpdateUserUSMPesos(id int, newUSM_Pesos int) error {
	payload := map[string]int{
		"usm_pesos": newUSM_Pesos,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPatch, BACKEND_URL+"/user/"+strconv.Itoa(id), bytes.NewBuffer(payloadJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user usm pesos, status code: %d", res.StatusCode)
	}

	return nil
}

// updateLoanStatus updates the status of a loan to "Returned" via HTTP PATCH request
func updateLoanStatus(loanID int) error {
	loanUpdate := map[string]string{
		"status": "Returned",
	}
	loanUpdateJson, err := json.Marshal(loanUpdate)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPatch, BACKEND_URL+"/loans/"+strconv.Itoa(loanID), bytes.NewBuffer(loanUpdateJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update loan status, status code: %d", res.StatusCode)
	}

	return nil
}