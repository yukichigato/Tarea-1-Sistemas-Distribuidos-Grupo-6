package bd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func IniciarBD(ruta string) *sql.DB {
	bd, err := sql.Open("sqlite3", ruta)
	if err != nil {
		log.Fatal(err)
	}

	crearTablas(bd)
	return bd
}

func crearTablas(bd *sql.DB) {
	crearUsuarios := `CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nombre TEXT,
		apellido TEXT,
		correo TEXT UNIQUE,
		contrasena TEXT,
		usm_pesos INTEGER DEFAULT 0
	);`

	crearLibros := `CREATE TABLE IF NOT EXISTS libros (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		titulo TEXT,
		categoria TEXT,
		tipo_transaccion TEXT,
		precio INTEGER,
		estado TEXT,
		puntuacion_popularidad INTEGER DEFAULT 0
	);`

	crearInventario := `CREATE TABLE IF NOT EXISTS inventario (
		id INTEGER PRIMARY KEY,
		cantidad_disponible INTEGER,
		FOREIGN KEY(id) REFERENCES libros(id)
	);`

	crearPrestamos := `CREATE TABLE IF NOT EXISTS prestamos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		usuario_id INTEGER,
		libro_id INTEGER,
		fecha_inicio TEXT,
		fecha_devolucion TEXT,
		estado TEXT,
		FOREIGN KEY(usuario_id) REFERENCES usuarios(id),
		FOREIGN KEY(libro_id) REFERENCES libros(id)
	);`

	crearVentas := `CREATE TABLE IF NOT EXISTS ventas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		usuario_id INTEGER,
		libro_id INTEGER,
		fecha_venta TEXT,
		FOREIGN KEY(usuario_id) REFERENCES usuarios(id),
		FOREIGN KEY(libro_id) REFERENCES libros(id)
	);`

	sentencias := []string{
		crearUsuarios,
		crearLibros,
		crearInventario,
		crearPrestamos,
		crearVentas,
	}

	for _, stmt := range sentencias {
		_, err := bd.Exec(stmt)
		if err != nil {
			log.Fatalf("Error creando tablas: %v\n", err)
		}
	}
}

