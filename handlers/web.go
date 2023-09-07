package handlers

import (
	"net/http"
	"puestos/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var puestos []models.Puesto
	db.Find(&puestos)

	c.HTML(http.StatusOK, "list.html", gin.H{"puestos": puestos})
}

func ShowForm(c *gin.Context) {

	c.HTML(http.StatusOK, "create.html", nil)

}

func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var newPuesto models.Puesto
	if err := c.ShouldBind(&newPuesto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&newPuesto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//c.JSON(http.StatusOK, newUser)
	c.Redirect(http.StatusSeeOther, "/")
}

func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Puesto
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}

	c.HTML(http.StatusOK, "show.html", gin.H{"puesto": puesto})
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Puesto
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}
	if err := c.ShouldBind(&puesto); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": err.Error()})
		return
	}

	db.Save(&puesto)
	c.Redirect(http.StatusSeeOther, "/")
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var puesto models.Puesto
	result := db.First(&puesto, id)
	if result.Error != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Puesto no encontrado"})
		return
	}
	db.Delete(&puesto)
	c.Redirect(http.StatusSeeOther, "/")
}
