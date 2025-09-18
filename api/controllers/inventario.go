package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para obtener inventario de un libro
func GetInventoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		inventory, err := models.GetInventoryById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, inventory)
	}
}

// Handler para actualizar inventario de un libro
func UpdateInventoryHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		if err := models.UpdateInventory(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "inventario actualizado correctamente"})
	}
}
