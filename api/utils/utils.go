package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Función para convertir entrada json a struct
func BindJSON(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return false
	}
	return true
}

// Función para obtener id de URL
func ParseID(c *gin.Context) (int, bool) {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return 0, false
	}
	return id, true
}
