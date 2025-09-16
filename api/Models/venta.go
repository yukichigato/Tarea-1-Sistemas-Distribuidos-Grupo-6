package models

import (
	"database/sql"
	"errors"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// Listar ventas
func ListSales(db *sql.DB) ([]structs.Sale, error) {
	rows, err := db.Query("SELECT id, user_id, book_id, sale_date FROM ventas")
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
	err := db.QueryRow("SELECT id, user_id, boook_id, sale_date FROM ventas WHERE id=$1", id).Scan(
		&sale.Id,
		&sale.UserId,
		&sale.BookId,
		&sale.SaleDate,
	)
	if err == sql.ErrNoRows {
		return sale, errors.New("venta no encontrada")
	}

	return sale, err
}

// Actualizar venta
func UpdateSale(db *sql.DB, id int, saleUpd structs.SaleUpdate) error {
	response, err := db.Exec("UPDATE ventas SET user_id=$1, book_id=$2, sale_date=$3 WHERE id=$4", saleUpd.UserId, saleUpd.BookId, saleUpd.SaleDate, id)
	if err != nil {
		return err
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no se encontro venta con esa id")
	}

	return nil
}
