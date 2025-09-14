package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/Models"
)

// Handler para registrar usuario
func InsertarUsuarioHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
			UsmPesos  int    `json:"usm_pesos"`
		}

		// Convertir a struct la entrada
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
			return
		}

		user := models.User{
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  input.Password,
			UsmPesos:  input.UsmPesos,
		}

		// Llamar modelo
		if err := models.InsertarUsuario(db, user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "usuario registrado con exito"})
	}
}
