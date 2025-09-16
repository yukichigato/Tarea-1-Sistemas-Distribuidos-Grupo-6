package models

import (
	"database/sql"
	"errors"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// Listar libros disponibles
func ListBooks(db *sql.DB) ([]structs.Book, error) {
	query := `
		SELECT l.id, l.book_name, l.book_category, l.tansaction_type, l.price, l.status, l.popularity_score, l.created_at, i.available_quantity
		FROM libros l
		JOIN inventario i ON l.id = i.id
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
			&book.PublicationDate,
			&book.Inventory.AvailableQuantity,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	// Errores de iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// Obtener libro específico por id
func GetBookById(db *sql.DB, id int) (structs.Book, error) {
	query := `
		SELECT l.id, l.book_name, l.book_category, l.tansaction_type, l.price, l.status, l.popularity_score, l.created_at, i.available_quantity
		FROM libros l
		JOIN inventario i ON l.id = i.id
		WHERE l.id=$1 AND i.available_quantity > 0
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
		&book.PublicationDate,
		&book.Inventory.AvailableQuantity,
	)

	if err == sql.ErrNoRows {
		return book, errors.New("libro no encontrado")
	}

	return book, err
}

// Actualizar libro (prestamo/compra)
func UpdateBook(db *sql.DB, id int, popularity_score int, available_quantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Actualizar inventario
	_, err = tx.Exec(
		"UPDATE inventario SET available_quantity=$1 WHERE id=$2", available_quantity, id,
	)
	if err != nil {
		return err
	}

	// Actualizar popularidad del libro
	_, err = tx.Exec(
		"UPDATE libros SET popularity_score=$1 WHERE id=$2", popularity_score, id,
	)
	if err != nil {
		return err
	}

	// Actualizar estado
	if available_quantity == 0 {
		_, err = tx.Exec(
			"UPDATE libros SET status=$1 WHERE id=$2", "agotado", id,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Registrar libro
func InsertBook(db *sql.DB, book structs.BookInput) error {
	// Verificar que el libro no exista
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM libros WHERE book_name=$1", book.BookName).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("el libro ya existe en el catalogo")
	}

	// Insertar libro
	_, err = db.Exec(
		"INSERT INTO libros (book_name, book_category, transaction_type, price, status, popularity_score, inventory) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		book.BookName, book.BookCategory, book.TransactionType, book.Price, book.Status, book.PopularityScore, book.Inventory,
	)
	return err
}
