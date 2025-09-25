package bd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// IniciarBD abre la base de datos y crea las tablas si no existen
func IniciarBD(ruta string) *sql.DB {
	bd, err := sql.Open("sqlite3", ruta)
	if err != nil {
		log.Fatal(err)
	}

	crearTablas(bd)
	return bd
}

// crearTablas define la estructura de la base de datos según la especificación
func crearTablas(bd *sql.DB) {
	crearUsers := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		usm_pesos INTEGER DEFAULT 0
	);`

	crearBooks := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		book_name TEXT,
		book_category TEXT,
		transaction_type TEXT,
		price INTEGER,
		status TEXT,
		popularity_score INTEGER DEFAULT 0
	);`

	crearInventory := `CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY,
		available_quantity INTEGER,
		FOREIGN KEY(id) REFERENCES books(id)
	);`

	crearLoans := `CREATE TABLE IF NOT EXISTS loans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		start_date TEXT,
		return_date TEXT,
		status TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	crearSales := `CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		sale_date TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	sentencias := []string{
		crearUsers,
		crearBooks,
		crearInventory,
		crearLoans,
		crearSales,
	}

	for _, stmt := range sentencias {
		_, err := bd.Exec(stmt)
		if err != nil {
			log.Fatalf("Error creando tablas: %v\n", err)
		}
	}
}
