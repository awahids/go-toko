package entity

import "gorm.io/gorm"

type Users struct {
	*gorm.Model
	Name	string		`json:"name" binding:"required" gorm:"type:varchar(50);not null"`
}