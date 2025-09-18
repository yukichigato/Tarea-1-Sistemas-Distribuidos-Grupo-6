package models

import (
	"database/sql"
	"errors"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

// Listar usuarios
func ListUsers(db *sql.DB) ([]structs.User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, usm_pesos FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.UsmPesos,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Errores de iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Obtener usuario específico por id
func GetUserById(db *sql.DB, id int) (structs.User, error) {
	var user structs.User
	err := db.QueryRow("SELECT id, first_name, last_name, email, usm_pesos FROM usuarios WHERE id=?", id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.UsmPesos,
	)

	if err == sql.ErrNoRows {
		return user, errors.New("usuario no encontrado")
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// Insertar usuario en la tabla Usuario
func InsertUser(db *sql.DB, user structs.UserInput) error {
	// Verificar que el usuario no exista
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email=?", user.Email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("el correo ya esta registrado")
	}

	// Insertar usuario
	_, err = db.Exec(
		"INSERT INTO usuarios (first_name, last_name, email, password, usm_pesos) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password, 0,
	)
	return err
}

// Actualizar saldo usuario
func UpdateUserBalance(db *sql.DB, id int, delta int) error {
	_, err := db.Exec(
		"UPDATE usuarios SET usm_pesos = usm_pesos - ? WHERE id=?", delta, id)
	return err
}
