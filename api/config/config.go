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

	return Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ClientURL:  os.Getenv("CLIENT_URL"),
	}
}
