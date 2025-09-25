package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	config "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/Config"
)

func main() {
	cfg := config.LoadConfig()

	// Conectar db
	db, err := bd.InitDB(cfg.DBDriver, cfg.DBName)
	if err != nil {
		log.Fatalf("No se pudo inicializar la DB: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	config.SetupRoutes(router, db)

	// Healtcheck ping-pong
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Iniciar aplicación
	address := ":" + cfg.Port
	if address == ":" {
		address = ":8080"
	}
	log.Printf("API escuchando en %s\n", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
