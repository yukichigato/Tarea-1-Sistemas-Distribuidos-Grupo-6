package models

import (
	"database/sql"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/models/structs"
)

// Listar todas las transacciones
func ListTransactions(db *sql.DB) ([]structs.Transaction, error) {
	query := `
		SELECT
			p.id AS transaction_id,
			l.id AS book_id,
			l.book_name AS book_name,
			'Arriendo' AS transaction_type,
			p.start_date AS transaction_date,
			l.price AS book_price
		FROM loans p
		JOIN books l ON p.book_id = l.id

		UNION ALL

		SELECT
			v.id AS transaction_id,
			l.id AS book_id,
			l.book_name AS book_name,
			'Compra' AS transaction_type,
			v.sale_date AS transaction_date,
			l.price AS book_price
		FROM sales v
		JOIN books l ON v.book_id = l.id
		ORDER BY transaction_date DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []structs.Transaction

	for rows.Next() {
		var t structs.Transaction
		if err := rows.Scan(
			&t.TransactionId,
			&t.BookId,
			&t.BookName,
			&t.TransactionType,
			&t.TransactionDate,
			&t.BookPrice,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
