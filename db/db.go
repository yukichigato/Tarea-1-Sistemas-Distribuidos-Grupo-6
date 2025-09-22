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
        first_name TEXT,
        last_name TEXT,
        email TEXT UNIQUE,
        password TEXT,
        usm_pesos INTEGER DEFAULT 0
    );`

    crearLibros := `CREATE TABLE IF NOT EXISTS libros (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        book_name TEXT,
        book_category TEXT,
        transaction_type TEXT,
        price INTEGER,
        status TEXT,
        popularity_score INTEGER DEFAULT 0
    );`

    crearInventario := `CREATE TABLE IF NOT EXISTS inventario (
        id INTEGER PRIMARY KEY,
        available_quantity INTEGER,
        FOREIGN KEY(id) REFERENCES libros(id)
    );`

    crearPrestamos := `CREATE TABLE IF NOT EXISTS prestamos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        book_id INTEGER,
        start_date TEXT,
        return_date TEXT,
        status TEXT,
        FOREIGN KEY(user_id) REFERENCES usuarios(id),
        FOREIGN KEY(book_id) REFERENCES libros(id)
    );`

    crearVentas := `CREATE TABLE IF NOT EXISTS ventas (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        book_id INTEGER,
        sale_date TEXT,
        FOREIGN KEY(user_id) REFERENCES usuarios(id),
        FOREIGN KEY(book_id) REFERENCES libros(id)
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
