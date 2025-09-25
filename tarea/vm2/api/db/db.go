package bd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)
	return db
}


func createTables(db *sql.DB) {
	createUsers := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		usm_pesos INTEGER DEFAULT 0
	);`

	createBooks := `CREATE TABLE IF NOT EXISTS books (
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

	createLoans := `CREATE TABLE IF NOT EXISTS loans (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		start_date TEXT,
		return_date TEXT, 
		status TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	createSales := `CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		book_id INTEGER,
		sale_date TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id),
		FOREIGN KEY(book_id) REFERENCES books(id)
	);`

	statements := []string{
		createUsers,
		createBooks,
		createInventory,
		createLoans,
		createSales,
	}

	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Error creating tables: %v\n", err)
		}
	}
}
