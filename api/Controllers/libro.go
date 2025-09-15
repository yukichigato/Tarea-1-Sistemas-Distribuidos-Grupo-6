package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para listar todos los libros
func ListBooksHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := models.ListBooks(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, books)
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

		c.JSON(http.StatusOK, gin.H{
			"id":               book.Id,
			"book_name":        book.BookName,
			"book_category":    book.BookCategory,
			"transaction_type": book.TransactionType,
			"price":            book.Price,
			"status":           book.Status,
			"popularity_score": book.PopularityScore,
		})
	}
}

// Handler para actualizar libro (prestamo/compra)
func UpdateBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		// Verificar existencias del libro
		available_quantity, err := models.GetInventory(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if available_quantity == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "inventario insuficiente del libro solicitado"})
			return
		}

		// Obtener popularidad del libro
		book, err := models.GetBookById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		newQuantity := available_quantity - 1
		newPopularity := book.PopularityScore + 1

		// Actualizar popularidad e inventario del libro
		if err := models.UpdateBook(db, id, newPopularity, newQuantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo actualizar el libro"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "libro actualizado correctamente"})

	}
}

// handler para registrar libro
func InsertBookHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.BookInput
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
