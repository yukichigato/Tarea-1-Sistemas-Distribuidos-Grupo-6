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

// Hanlder para listar préstamos
func ListLoansHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")

		var loans []structs.Loan
		var err error

		if idStr != "" {
			userId, convErr := strconv.Atoi(idStr)
			if convErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "id de usuario inválido"})
				return
			}
			loans, err = models.GetUserLoans(db, userId)
		} else {
			loans, err = models.ListLoans(db)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"loans": loans})
	}
}

// Handler para obtener préstamo
func GetLoanHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		loan, err := models.GetLoanById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, loan)
	}
}

// Handler para actualizar estado de un préstamo
func UpdateLoanHandler(db *sql.DB) gin.HandlerFunc {
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

		if err := models.UpdateLoan(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "prestamo actualizado con exito"})

	}
}

// Handler para insertar préstamo
func InsertLoanHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input structs.LoanInput
		if !utils.BindJSON(c, &input) {
			return
		}

		if err := models.InsertLoan(db, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "prestamo registrado con exito"})
	}
}
