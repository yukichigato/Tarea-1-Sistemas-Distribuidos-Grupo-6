package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/Controllers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/users", controllers.InsertarUsuarioHandler(db)) // Registrar nuevo usuario
}
