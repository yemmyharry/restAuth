package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(32) not null"`
	Email    string `json:"email" gorm:"type:varchar(32) unique not null"`
	Password string `json:"-" gorm:"type:varchar(256) not null"`
}

type PasswordReset struct {
	Id    uint   `gorm:"primary_key" json:"id"`
	Email string `gorm:"type:varchar(32) unique not null" json:"email"`
	Token string `gorm:"type:varchar(256) not null" json:"token"`
}
