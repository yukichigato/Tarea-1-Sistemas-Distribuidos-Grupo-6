package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para listar todos los libros disponibles
func ListBooksHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := models.ListBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"books": books})
	}
}

// Handler para obtener libro
func GetBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		book, err := models.GetBookById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

// Handler para actualizar libro (prestamo/compra)
func UpdateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		var input map[string]any
		if !utils.BindJSON(c, &input) {
			return
		}
		delete(input, "id")

		if len(input) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No hay campos válidos para actualizar"})
			return
		}

		if err := models.UpdateBook(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "libro actualizado con exito"})

	}
}

// handler para registrar libro
func InsertBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input structs.BookInput
		if !utils.BindJSON(c, &input) {
			return
		}

		if err := models.InsertBook(db, input); err != nil {
			if err.Error() == "el libro ya existe en el catalogo" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "libro registrado con exito"})
	}
}
