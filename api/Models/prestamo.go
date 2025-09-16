package models

import (
	"database/sql"
	"errors"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// Listar préstamos
func ListLoans(db *sql.DB) ([]structs.Loan, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, start_date, return_date, status FROM prestamos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loans []structs.Loan
	for rows.Next() {
		var loan structs.Loan
		if err := rows.Scan(
			&loan.Id,
			&loan.UserId,
			&loan.BookId,
			&loan.StartDate,
			&loan.ReturnDate,
			&loan.Status,
		); err != nil {
			return nil, err
		}

		loans = append(loans, loan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return loans, err
}

// Obtener préstamo específico por id
func GetLoanById(db *sql.DB, id int) (structs.Loan, error) {
	var loan structs.Loan
	err := db.QueryRow("SELECT id, user_id, book_id, start_date, return_date, status FROM prestamos WHERE id=$1", id).Scan(
		&loan.Id,
		&loan.UserId,
		&loan.BookId,
		&loan.StartDate,
		&loan.ReturnDate,
		&loan.Status,
	)

	if err == sql.ErrNoRows {
		return loan, errors.New("prestamo no encontrado")
	}

	return loan, err
}

// Actualizar (estado) préstamo
func UpdateLoanStatus(db *sql.DB, id int, loanUpd structs.LoanStatusUpdate) error {
	response, err := db.Exec("UPDATE prestamos SET return_date = NOW(), status=$1 WHERE id=$2", loanUpd.Status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no se encontro prestamo con esa id")
	}

	return nil
}

// Insertar préstamo
func InsertLoan(db *sql.DB, loan structs.LoanInput) error {
	_, err := db.Exec(
		"INSERT INTO prestamos (user_id, book_id, start_date, status) VALUES ($1, $2, NOW(), $3)",
		loan.UserId, loan.BookId, "pendiente",
	)
	return err
}
