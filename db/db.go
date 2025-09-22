package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB abre la conexión e inicializa todas las tablas
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	createUser := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		usm_pesos INTEGER DEFAULT 0
	);`

	createBook := `CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		book_name TEXT,
		book_category TEXT,
		transaction_type TEXT,
		price INTEGER,
		status TEXT,
		popularity_score INTEGER DEFAULT 0
	);`

	createInventory := `CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY,
		available_quantity INTEGER,
		FOREIGN KEY(id) REFERENCES books(id)
	);`

	createLoan := `CREATE TABLE IF NOT EXISTS loans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		start_date TEXT,
		return_date TEXT,
		status TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	createSale := `CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		sale_date TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	createTransaction := `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		type TEXT,
		date TEXT,
		amount INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	statements := []string{
		createUser,
		createBook,
		createInventory,
		createLoan,
		createSale,
		createTransaction,
	}

	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Error creando tablas: %v\n", err)
		}
	}
}
