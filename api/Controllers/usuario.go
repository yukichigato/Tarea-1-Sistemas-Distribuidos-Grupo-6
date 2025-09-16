package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/models/structs"
	"github.com/yukichigato/Tarea-1-Sistemas-Distribuidos-Grupo-6/api/utils"
)

// Handler para registrar usuario
func InsertUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input structs.UserInput
		if !utils.BindJSON(c, &input) {
			return
		}

		// Llamar modelo
		if err := models.InsertUser(db, input); err != nil {
			if err.Error() == "el correo ya esta registrado" {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "usuario registrado con exito"})
	}
}

// Handler para validar inicio de sesión
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input structs.UserLogin
		if !utils.BindJSON(c, &input) {
			return
		}

		users, err := models.ListUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, u := range users {
			if u.Email == input.Email && u.Password == input.Password {
				c.JSON(http.StatusOK, gin.H{"message": "login exitoso"})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email o contraseña incorrectos"})
	}
}

// Handler para actualizar saldo de usm pesos
func BalanceHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		var input structs.UserBalanceUpdate
		if !utils.BindJSON(c, &input) {
			return
		}

		if err := models.UpdateUserBalance(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "saldo actualizado correctamente"})
	}
}

// Handler para listar usuarios
func ListUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.ListUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// Handler para obtener usuario
func GetUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := utils.ParseID(c)
		if !ok {
			return
		}

		user, err := models.GetUserById(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
