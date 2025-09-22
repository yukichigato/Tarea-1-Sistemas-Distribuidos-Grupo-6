package main

import (
	"fmt"
	"log"
	
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/bd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Inicia la base de datos usando tu paquete 'bd'
	db := bd.IniciarBD("mi_base.db")
	defer db.Close()

	fmt.Println("Poblando la base de datos...")

	// Inserta un usuario de ejemplo
	_, err := db.Exec(`INSERT INTO usuarios (nombre, apellido, correo, contrasena, usm_pesos) VALUES (?, ?, ?, ?, ?);`,
		"Cristobal", "Alvarez", "cristobalalvarez@gmail.com", "123456", 100)
	if err != nil {
		log.Fatalf("Error al insertar usuario: %v", err)
	}

	// Inserta un libro de ejemplo
	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"El principito", "Infantil", "Venta", 10, "Disponible", 5)
	if err != nil {
		log.Fatalf("Error al insertar libro: %v", err)
	}
	
	// Inserta inventario para el libro
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 1, 5)
	if err != nil {
		log.Fatalf("Error al insertar inventario: %v", err)
	}

	fmt.Println("¡Base de datos poblada con éxito!")
}