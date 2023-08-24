package handlers

import (
	"itsva-puestos/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var funciones []models.Funcion
	db.Find(&funciones)

	c.JSON(http.StatusOK, gin.H{"funciones": funciones})
}

func ShowAPI(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var funcion models.Funcion
	if err := db.First(&funcion, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Funcion no encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"funcion": funcion})
}
