package models

import (
	"database/sql"
	"errors"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	UsmPesos  int
}

// Insertar usuario en la tabla Usuario
func InsertarUsuario(db *sql.DB, user User) error {
	// Verificar que el usuario no exista
	var cant_usuarios_iguales int
	err := db.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email=$1", user.Email).Scan(&cant_usuarios_iguales)
	if err != nil {
		return err
	}
	if cant_usuarios_iguales > 0 {
		return errors.New("El correo ya esta registrado")
	}

	// Insertar usuario
	_, err = db.Exec(
		"INSERT INTO usuarios (first_name, last_name, email, password, usm_pesos) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.Password, user.UsmPesos,
	)
	return err
}
