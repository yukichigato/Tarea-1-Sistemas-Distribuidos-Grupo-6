package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/models/structs"
)

// Listar préstamos
func ListLoans(db *sql.DB) ([]structs.Loan, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, start_date, return_date, status FROM loans")
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

	return loans, nil
}

// Obtener préstamo específico por id
func GetLoanById(db *sql.DB, id int) (structs.Loan, error) {
	var loan structs.Loan
	err := db.QueryRow("SELECT id, user_id, book_id, start_date, return_date, status FROM loans WHERE id=?", id).Scan(
		&loan.Id,
		&loan.UserId,
		&loan.BookId,
		&loan.StartDate,
		&loan.ReturnDate,
		&loan.Status,
	)

	if err == sql.ErrNoRows {
		return loan, errors.New("prestamo no encontrado")
	} else if err != nil {
		return loan, err
	}

	return loan, nil
}

// Actualizar préstamo
func UpdateLoan(db *sql.DB, id int, loanUpdates map[string]any) error {
	sets := []string{}
	args := []any{}

	for field, value := range loanUpdates {
		sets = append(sets, fmt.Sprintf("%s=?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE loans SET %s WHERE id=?", strings.Join(sets, ", "))
	_, err := db.Exec(query, args...)

	return err
}

// Insertar préstamo
func InsertLoan(db *sql.DB, loan structs.LoanInput) error {
	_, err := db.Exec(
		"INSERT INTO loans (user_id, book_id, start_date, return_date, status) VALUES (?, ?, DATE('now'), DATE('now', '+7 days'), ?)",
		loan.UserId, loan.BookId, "pendiente",
	)
	return err
}

// Listar préstamos de un usuario por user_id
func GetUserLoans(db *sql.DB, userId int) ([]structs.Loan, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, start_date, return_date, status FROM loans WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userLoans []structs.Loan
	for rows.Next() {
		var userLoan structs.Loan
		if err := rows.Scan(
			&userLoan.Id,
			&userLoan.UserId,
			&userLoan.BookId,
			&userLoan.StartDate,
			&userLoan.ReturnDate,
			&userLoan.Status,
		); err != nil {
			return nil, err
		}

		userLoans = append(userLoans, userLoan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userLoans, nil
}
