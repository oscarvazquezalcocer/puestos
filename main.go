package main

import (
	"net/http"
	"puestos/config"
	"puestos/models"
	"puestos/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la conexion a la base de datos"})
			c.Abort()
			return
		}
		c.Set("db", db)
		c.Next()
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func main() {

	config.ConfigureViper()

	db, err := gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{})
	if err != nil {
		panic("Error conexión a la base de datos")
	}

	// Realizar las migraciones antes de iniciar el servidor
	db.AutoMigrate(&models.Puesto{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Use(DatabaseMiddleware())
	routes.AddWEB(r)
	routes.AddAPI(r)

	r.Run(viper.GetString("PORT"))
}
