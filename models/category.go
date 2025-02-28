package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required"`
}
