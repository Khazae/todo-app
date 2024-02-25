package models

import (
	"gorm.io/gorm"
)


type Task struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status" gorm:"default:false"`
	gorm.Model
}
