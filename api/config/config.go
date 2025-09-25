package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ClientURL  string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env no encontrado, usando variables de entorno del sistema")
	}

	dbName := os.Getenv("DB_NAME")
	dbDriver := "sqlite3"
	if dbName == "" {
		dbName = "mi_base.db"
	}

	return Config{
		Port:      os.Getenv("PORT"),
		DBDriver:  dbDriver,
		DBName:    dbName,
		ClientURL: os.Getenv("CLIENT_URL"),
	}
}
