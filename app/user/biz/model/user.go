package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex;size:255;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
}