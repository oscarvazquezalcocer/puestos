package models

import "gorm.io/gorm"

type Puesto struct {
	gorm.Model
	Nombre string `json:"nombre" form:"nombre" binding:"required" gorm:"unique"`
	Clave  string `json:"clave"  form:"clave"  binding:"required" gorm:"unique"`
}
