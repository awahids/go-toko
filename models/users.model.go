package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name	string		`json:"name" binding:"required" gorm:"type:varchar(50);not null"`
}