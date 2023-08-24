package handlers

import (
	"itsva-puestos/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var funciones []models.Funcion
	db.Find(&funciones)

	c.HTML(http.StatusOK, "list.html", gin.H{"funciones": funciones})
}

func ShowForm(c *gin.Context) {

	c.HTML(http.StatusOK, "create.html", nil)

}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var newFuncion models.Funcion
	if err := c.ShouldBind(&newFuncion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&newFuncion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, newUser)
	c.Redirect(http.StatusSeeOther, "/")
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var funcion models.Funcion
	result := db.First(&funcion, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Funcion no encontrada"})
		return
	}

	c.HTML(http.StatusOK, "show.html", gin.H{"funcion": funcion})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var funcion models.Funcion
	result := db.First(&funcion, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}
	if err := c.ShouldBind(&funcion); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})
		return
	}

	db.Save(&funcion)
	c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var funcion models.Funcion
	result := db.First(&funcion, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Funcion no encontrada"})
		return
	}
	db.Delete(&funcion)
	c.Redirect(http.StatusSeeOther, "/")
}
