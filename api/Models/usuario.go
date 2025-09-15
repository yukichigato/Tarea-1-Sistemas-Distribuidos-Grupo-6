package models

import (
	"database/sql"
	"errors"
)

// Listar usuarios
func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, usm_pesos FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
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

	return users, err
}

// Obtener usuario específico
func GetUserById(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, first_name, last_name, email, usm_pesos FROM usuarios WHERE id=$1", id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.UsmPesos,
	)

	if err == sql.ErrNoRows {
		return user, errors.New("usuario no encontrado")
	}

	return user, err
}

// Insertar usuario en la tabla Usuario
func InsertUser(db *sql.DB, user UserInput) error {
	// Verificar que el usuario no exista
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email=$1", user.Email).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("el correo ya esta registrado")
	}

	// Insertar usuario
	_, err = db.Exec(
		"INSERT INTO usuarios (first_name, last_name, email, password, usm_pesos) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.Password, user.UsmPesos,
	)
	return err
}

// Actualizar (saldo) usuario
func UpdateUser(db *sql.DB, id int, userUpd UserUpdate) error {
	_, err := db.Exec(
		"UPDATE usuarios SET usm_pesos=$1 WHERE id=$2", userUpd.UsmPesos, id,
	)
	return err
}
