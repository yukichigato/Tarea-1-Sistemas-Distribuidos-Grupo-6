package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para listar ventas
func ListSalesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")

		var sales []structs.Sale
		var err error

		if idStr != "" {
			userId, convErr := strconv.Atoi(idStr)
			if convErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "id de usuario inválido"})
				return
			}
			sales, err = models.GetUserSales(db, userId)
		} else {
			sales, err = models.ListSales(db)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, sales)
	}
}

// Handler para obtener venta
func GetSaleHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		sale, err := models.GetSaleById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sale)
	}
}

// Handler para actualizar venta
func UpdateSaleHandler(db *sql.DB) gin.HandlerFunc {
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

		if err := models.UpdateSale(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "prestamo actualizado con exito"})

	}
}

// Handler para registrar venta
func InsertSaleHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input structs.SaleInput
		if !utils.BindJSON(c, &input) {
			return
		}

		if err := models.InsertSale(db, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "venta registrada con exito"})
	}
}
