package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/models"
)

// Handler para listar transacciones
func ListTransactionsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactions, err := models.ListTransactions(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}
