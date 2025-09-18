package models

import (
	"database/sql"
	"errors"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// Obtener inventario de libro específico
func GetInventoryById(db *sql.DB, id int) (structs.Inventory, error) {
	var inventory structs.Inventory
	err := db.QueryRow("SELECT available_quantity FROM inventario WHERE id=?", id).Scan(&inventory.AvailableQuantity)

	if err == sql.ErrNoRows {
		return inventory, errors.New("registro no encontrado")
	} else if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// Actualizar inventario de un libro
func UpdateInventory(db *sql.DB, id int) error {
	_, err := db.Exec(
		"UPDATE inventario SET available quantity = available quantity - 1 WHERE id=?", id)
	return err
}
