package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// fetchCartInfo retrieves book details for a list of book IDs in the shopping cart
func fetchCartInfo(cart []int) ([]Book, error) {
	var books []Book
	for _, bookID := range cart {
		book, err := FetchBook(bookID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// calculateDaysLeft calculates the number of days between two dates in DD-MM-YYYY format
func calculateDaysLeft(start string, end string) (int, error) {
	d1, err := time.Parse("02-01-2006", start)
	if err != nil {
		return 0, err
	}
	d2, err := time.Parse("02-01-2006", end)
	if err != nil {
		return 0, err
	}
	return int(d2.Sub(d1).Hours() / 24), nil
}

// optimizeCart removes the most expensive books from cart until total price fits within budget
func optimizeCart(cart []Book, budget int) []Book {
	optimizedCart := make([]Book, len(cart))
	copy(optimizedCart, cart)

	sort.Slice(optimizedCart, func(i, j int) bool {
		return optimizedCart[i].Price > optimizedCart[j].Price
	})

	for calculateBookTotal(optimizedCart) > budget && len(optimizedCart) > 0 {
		optimizedCart = optimizedCart[1:]
	}

	return optimizedCart
}

// calculateBookTotal calculates the total price of a list of books
func calculateBookTotal(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.Price
	}
	return total
}

// readInput prompts the user for input and returns the trimmed string response
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}