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

func GetCatalog() ([]Book, error) {
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