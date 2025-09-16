package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para listar ventas
func ListSalesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sales, err := models.ListSales((db))
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

		var input structs.SaleUpdate
		if !utils.BindJSON(c, &input) {
			return
		}

		if err := models.UpdateSale(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "venta actualizada correctamente"})
	}
}
