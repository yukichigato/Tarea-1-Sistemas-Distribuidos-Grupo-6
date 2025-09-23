package main

import (
	"fmt"
	"log"
	
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/bd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := bd.IniciarBD("mi_base.db")
	defer db.Close()

	fmt.Println("Poblando la base de datos con datos de ejemplo...")

	// Inserta un usuario de ejemplo
	_, err := db.Exec(`INSERT INTO usuarios (nombre, apellido, correo, contrasena, usm_pesos) VALUES (?, ?, ?, ?, ?);`,
		"Cristobal", "Alvarez", "cristobalalvarez@gmail.com", "123456", 100)
	if err != nil {
		log.Fatalf("Error al insertar usuario: %v", err)
	}

	// Inserta los libros del catálogo
	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"El principito", "Infantil", "Venta", 10, "Disponible", 4)
	if err != nil {
		log.Fatalf("Error al insertar libro 1: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 1, 6)
	if err != nil {
		log.Fatalf("Error al insertar inventario para libro 1: %v", err)
	}

	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"Papelucho", "Infantil", "Arriendo", 2, "Disponible", 2)
	if err != nil {
		log.Fatalf("Error al insertar libro 2: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 2, 11)
	if err != nil {
		log.Fatalf("Error al insertar inventario para libro 2: %v", err)
	}

	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"Metro 2033", "Guerra", "Venta", 19, "Disponible", 1)
	if err != nil {
		log.Fatalf("Error al insertar libro 3: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 3, 1)
	if err != nil {
		log.Fatalf("Error al insertar inventario para libro 3: %v", err)
	}

	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"Divergente", "Distopia", "Arriendo", 3, "Disponible", 1)
	if err != nil {
		log.Fatalf("Error al insertar libro 4: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 4, 1)
	if err != nil {
		log.Fatalf("Error al insertar inventario para libro 4: %v", err)
	}

	_, err = db.Exec(`INSERT INTO libros (titulo, categoria, tipo_transaccion, precio, estado, puntuacion_popularidad) VALUES (?, ?, ?, ?, ?, ?);`,
		"Cien años de soledad", "Ficción", "Venta", 25, "Agotado", 0)
	if err != nil {
		log.Fatalf("Error al insertar libro agotado: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO inventario (id, cantidad_disponible) VALUES (?, ?);`, 5, 0)
	if err != nil {
		log.Fatalf("Error al insertar inventario para libro agotado: %v", err)
	}

	// Inserta datos de prestamos 
	_, err = db.Exec(`INSERT INTO prestamos (usuario_id, libro_id, fecha_inicio, fecha_devolucion, estado) VALUES (?, ?, ?, ?, ?);`,
		1, 1, "23/08/2025", "17/09/2025", "Devuelto")
	if err != nil {
		log.Fatalf("Error al insertar préstamo 1: %v", err)
	}

	_, err = db.Exec(`INSERT INTO prestamos (usuario_id, libro_id, fecha_inicio, fecha_devolucion, estado) VALUES (?, ?, ?, ?, ?);`,
		1, 2, "20/08/2025", "14/09/2025", "En curso")
	if err != nil {
		log.Fatalf("Error al insertar préstamo 2: %v", err)
	}

	// Inserta datos de ventas
	_, err = db.Exec(`INSERT INTO ventas (usuario_id, libro_id, fecha_venta) VALUES (?, ?, ?);`,
		1, 1, "05/08/2025")
	if err != nil {
		log.Fatalf("Error al insertar venta 1: %v", err)
	}
	
	_, err = db.Exec(`INSERT INTO ventas (usuario_id, libro_id, fecha_venta) VALUES (?, ?, ?);`,
		1, 2, "17/08/2025")
	if err != nil {
		log.Fatalf("Error al insertar venta 2: %v", err)
	}

	fmt.Println("¡Base de datos poblada con éxito!")
}
