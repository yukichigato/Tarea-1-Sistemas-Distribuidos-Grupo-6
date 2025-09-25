package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controllers "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/controllers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Endpoints solicitados en la tarea
	router.GET("/users", controllers.ListUsersHandler(db))        // Listar todos los usuarios
	router.GET("/users/:id", controllers.GetUserHandler(db))      // Obtener usuario por id
	router.PATCH("/users/:id", controllers.UpdateUserHandler(db)) // Actualiza saldo de usm pesos del usuario
	router.POST("/users", controllers.InsertUserHandler(db))      // Registrar nuevo usuario

	router.GET("/books", controllers.ListBooksHandler(db))        // Listar todos los libros
	router.GET("/books/:id", controllers.GetBookHandler(db))      // Obtener libro por id
	router.PATCH("/books/:id", controllers.UpdateBookHandler(db)) // Actualizar popularidad al adquirir un libro
	router.POST("/books", controllers.InsertBookHandler(db))      // Registrar nuevo libro

	router.GET("/loans", controllers.ListLoansHandler(db))        // Listar todos los préstamos
	router.GET("/loans/:id", controllers.GetLoanHandler(db))      // Obtener préstamo por id
	router.PATCH("/loans/:id", controllers.UpdateLoanHandler(db)) // Actualizar estado de préstamo
	router.POST("/loans", controllers.InsertLoanHandler(db))      // Registrar préstamo

	router.GET("/sales", controllers.ListSalesHandler(db))        // Listar todas las ventas
	router.GET("/sales/:id", controllers.GetSaleHandler(db))      // Obtener venta por id
	router.PATCH("/sales/:id", controllers.UpdateSaleHandler(db)) // Actualizar venta
	router.POST("/sales", controllers.InsertSaleHandler(db))      // Registrar nueva venta

	router.GET("/transactions", controllers.ListTransactionsHandler(db)) // Listar todas las transacciones

	// Endpoints adicionales, no solicitados en la tarea
	router.POST("/login", controllers.LoginHandler(db)) // Validar inicio de sesión

	router.GET("/inventory/:id", controllers.GetInventoryHandler(db))      // Obtener inventario de un libro por id
	router.PATCH("/inventory/:id", controllers.UpdateInventoryHandler(db)) // Actualizar inventario de un libro
}
