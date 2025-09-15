package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/controllers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Endpoints solicitados en la tarea
	router.GET("/users", controllers.ListUsersHandler(db))     // Listar todos los usuarios
	router.GET("/users/:id", controllers.GetUserHandler(db))   // Obtener usuario por id
	router.POST("/users", controllers.InsertUserHandler(db))   // Registrar nuevo usuario
	router.PATCH("/users/:id", controllers.BalanceHandler(db)) // Actualiza saldo de usm pesos del usuario

	router.GET("/books", controllers.ListBooksHandler(db))        // Listar todos los libros
	router.GET("/books/:id", controllers.GetBookHandler(db))      // Obtener libro por id
	router.PATCH("/books/:id", controllers.UpdateBookHandler(db)) // Actualizar popularidad y existencias al adquirir un libro
	router.POST("/books", controllers.InsertBookHandler(db))      // Registrar nuevo libro

	// Endpoints adicionales, no solicitados en la tarea
	router.POST("/login", controllers.LoginHandler(db)) // Validar inicio de sesión
}
