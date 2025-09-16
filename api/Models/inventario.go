package models

import (
	"database/sql"
	"errors"
)

// Obtener inventario de libro específico
func GetInventory(db *sql.DB, id int) (int, error) {
	var available_quantity int
	err := db.QueryRow("SELECT available_quantity FROM inventario WHERE id=$1", id).Scan(&available_quantity)

	if err == sql.ErrNoRows {
		return available_quantity, errors.New("registro no encontrado")
	}

	return available_quantity, err
}
