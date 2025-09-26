package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/models/structs"
)

// Listar ventas
func ListSales(db *sql.DB) ([]structs.Sale, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, sale_date FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []structs.Sale
	for rows.Next() {
		var sale structs.Sale
		if err := rows.Scan(
			&sale.Id,
			&sale.UserId,
			&sale.BookId,
			&sale.SaleDate,
		); err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sales, nil
}

// Obtener venta específica por id
func GetSaleById(db *sql.DB, id int) (structs.Sale, error) {
	var sale structs.Sale
	err := db.QueryRow("SELECT id, user_id, book_id, sale_date FROM sales WHERE id=?", id).Scan(
		&sale.Id,
		&sale.UserId,
		&sale.BookId,
		&sale.SaleDate,
	)
	if err == sql.ErrNoRows {
		return sale, errors.New("venta no encontrada")
	} else if err != nil {
		return sale, err
	}

	return sale, nil
}

// Actualizar venta
func UpdateSale(db *sql.DB, id int, saleUpdates map[string]any) error {
	sets := []string{}
	args := []any{}

	for field, value := range saleUpdates {
		sets = append(sets, fmt.Sprintf("%s=?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE sales SET %s WHERE id=?", strings.Join(sets, ", "))
	_, err := db.Exec(query, args...)

	return err
}

// Registrar venta
func InsertSale(db *sql.DB, sale structs.SaleInput) error {
	_, err := db.Exec(
		"INSERT INTO sales (user_id, book_id, sale_date) VALUES (?, ?, Date('now'))",
		sale.UserId, sale.BookId,
	)
	return err
}

// Listar compras de un usuario por user_id
func GetUserSales(db *sql.DB, userId int) ([]structs.Sale, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, sale_date FROM sales WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userSales []structs.Sale
	for rows.Next() {
		var userSale structs.Sale
		if err := rows.Scan(
			&userSale.Id,
			&userSale.UserId,
			&userSale.BookId,
			&userSale.SaleDate,
		); err != nil {
			return nil, err
		}

		userSales = append(userSales, userSale)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userSales, nil
}
