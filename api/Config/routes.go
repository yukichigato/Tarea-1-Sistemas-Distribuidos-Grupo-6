package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/Controllers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Endpoints solicitados en la tarea
	router.GET("/users", controllers.ListUsers(db))            // Listar todos los usuarios
	router.GET("/users/:id", controllers.GetUser(db))          // Obtener usuario por id
	router.POST("/users", controllers.InsertUserHandler(db))   // Registrar nuevo usuario
	router.PATCH("/users/:id", controllers.BalanceHandler(db)) // Actualiza saldo de usm pesos del usuario

	// Endpoints adicionales, no solicitados en la tarea
	router.POST("/login", controllers.LoginHandler(db)) // Validar inicio de sesión
}
