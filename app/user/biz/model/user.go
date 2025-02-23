package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex;size:255;not null"`
	Password string `json:"password" gorm:"size:255;not null"`
}

func (u *User) TableName() string {
	return "user"
}

func GetUserByEmail(ctx context.Context,db *gorm.DB, email string) (*User, error) {
	user := &User{}
	if err := db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(ctx context.Context,db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
