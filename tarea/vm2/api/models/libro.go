package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// ListBooks retrieves all available books from the database.
func ListBooks(db *sql.DB) ([]structs.Book, error) {
	query := `
		SELECT 
			l.id, l.book_name, l.book_category, l.transaction_type, 
			l.price, l.status, l.popularity_score
		FROM books l
		JOIN inventory i ON l.id = i.id
		WHERE i.available_quantity > 0
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []structs.Book
	for rows.Next() {
		var book structs.Book
		if err := rows.Scan(
			&book.Id,
			&book.BookName,
			&book.BookCategory,
			&book.TransactionType,
			&book.Price,
			&book.Status,
			&book.PopularityScore,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookById gets a specific book by id.
func GetBookById(db *sql.DB, id int) (structs.Book, error) {
	query := `
		SELECT 
			l.id, l.book_name, l.book_category, l.transaction_type, 
			l.price, l.status, l.popularity_score
		FROM books l
		JOIN inventory i ON l.id = i.id
		WHERE l.id=? AND i.available_quantity > 0
	`
	var book structs.Book
	err := db.QueryRow(query, id).Scan(
		&book.Id,
		&book.BookName,
		&book.BookCategory,
		&book.TransactionType,
		&book.Price,
		&book.Status,
		&book.PopularityScore,
	)

	if err == sql.ErrNoRows {
		return book, errors.New("Libro no encontrado")
	} else if err != nil {
		return book, err
	}

	return book, nil
}

// UpdateBook updates a book (loan/sale).
func UpdateBook(db *sql.DB, id int, bookUpdates map[string]any) error {
	sets := []string{}
	args := []any{}

	for field, value := range bookUpdates {
		sets = append(sets, fmt.Sprintf("%s=?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE books SET %s WHERE id=?", strings.Join(sets, ", "))
	_, err := db.Exec(query, args...)

	return err
}

// InsertBook registers a new book.
func InsertBook(db *sql.DB, book structs.BookInput) error {
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM books WHERE book_name=?", book.BookName).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("Libro no existe en catalogo")
	}

	_, err = db.Exec(
		"INSERT INTO books (book_name, book_category, transaction_type, price, status, popularity_score) VALUES (?, ?, ?, ?, ?, ?)",
		book.BookName, book.BookCategory, book.TransactionType, book.Price, book.Status, 0,
	)
	return err
}