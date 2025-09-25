package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
)

func ListUsers(db *sql.DB) ([]structs.User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, password, usm_pesos FROM users")
	if err != nil {
		log.Println("Error al ejecutar query:", err)
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
			&user.Password,
			&user.UsmPesos,
		); err != nil {
			log.Println("Error al escanear fila:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error de iteración de filas:", err)
		return nil, err
	}

	log.Printf("Usuarios cargados: %+v\n", users)
	return users, nil
}

// Obtener usuario específico por id
func GetUserById(db *sql.DB, id int) (structs.User, error) {
	var user structs.User
	err := db.QueryRow("SELECT id, first_name, last_name, email, usm_pesos FROM users WHERE id=?", id).Scan(
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
		"INSERT INTO users (first_name, last_name, email, password, usm_pesos) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password, 0,
	)
	return err
}

// Actualizar usuario
func UpdateUser(db *sql.DB, id int, userUpdates map[string]any) error {
	sets := []string{}
	args := []any{}

	for field, value := range userUpdates {
		sets = append(sets, fmt.Sprintf("%s=?", field))
		args = append(args, value)
	}
	args = append(args, id)

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=?", strings.Join(sets, ", "))
	_, err := db.Exec(query, args...)

	return err
}
