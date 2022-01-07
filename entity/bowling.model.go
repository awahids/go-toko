package entity

import "gorm.io/gorm"

type Bowling struct {
	*gorm.Model
	Quantity	int		`json:"quantity" binding:"required" gorm:"type:int(50);not null"`
	UsersId		int		`json:"userId" binding:"required" gorm:"type:int(100); not null"`
}