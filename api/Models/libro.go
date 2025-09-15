package models

import (
	"database/sql"
	"errors"
)

// Listar libros disponibles
func ListBooks(db *sql.DB) ([]Book, error) {
	query := `
		SELECT l.id, l.book_name, l.book_category, l.tansaction_type, l.price, l.status, l.popularity_score, i.available_quantity
		FROM libros l
		JOIN inventario i ON l.id = i.id
		WHERE i.available_quantity > 0
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(
			&book.Id,
			&book.BookName,
			&book.BookCategory,
			&book.TransactionType,
			&book.Price,
			&book.Status,
			&book.PopularityScore,
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

// Obtener libro específico
func GetBookById(db *sql.DB, id int) (Book, error) {
	var book Book
	err := db.QueryRow("SELECT * FROM libros WHERE id=$1", id).Scan(
		&book.Id,
		&book.BookName,
		&book.BookCategory,
		&book.TransactionType,
		&book.Price,
		&book.Status,
		&book.PopularityScore,
	)

	if err == sql.ErrNoRows {
		return book, errors.New("libro no encontrado")
	}

	return book, err
}

// Obtener inventario de libro específico
func GetInventory(db *sql.DB, id int) (int, error) {
	var available_quantity int
	err := db.QueryRow("SELECT available_quantity FROM inventario WHERE id=$1", id).Scan(&available_quantity)

	if err == sql.ErrNoRows {
		return available_quantity, errors.New("registro no encontrado")
	}

	return available_quantity, err
}

// Actualizar libro (prestamo/compra)
func UpdateBook(db *sql.DB, id int, popularity_score int, available_quantity int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Actualizar inventario
	_, err = tx.Exec(
		"UPDATE inventario SET available_quantity=$1 WHERE id=$2", available_quantity, id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Actualizar popularidad del libro
	_, err = tx.Exec(
		"UPDATE libros SET popularity_score=$1 WHERE id=$2", popularity_score, id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Actualizar estado
	if available_quantity == 0 {
		_, err = tx.Exec(
			"UPDATE libros SET status=$1 WHERE id=$2", "agotado", id,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// Registrar libro
func InsertBook(db *sql.DB, book BookInput) error {
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
