package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/models/structs"
)

// Obtener inventario de libro específico
func GetInventoryById(db *sql.DB, id int) (structs.Inventory, error) {
	var inventory structs.Inventory
	err := db.QueryRow("SELECT available_quantity FROM inventory WHERE id=?", id).Scan(&inventory.AvailableQuantity)

	if err == sql.ErrNoRows {
		return inventory, errors.New("registro no encontrado")
	} else if err != nil {
		return inventory, err
	}

	return inventory, nil
}

// Actualizar inventario de un libro
func UpdateInventory(db *sql.DB, id int, inventoryUpdate map[string]any) error {
	sets := []string{}
	args := []any{}

	for field, value := range inventoryUpdate {
		sets = append(sets, fmt.Sprintf("%s=?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE inventory SET %s WHERE id=?", strings.Join(sets, ", "))
	_, err := db.Exec(query, args...)

	return err
}
