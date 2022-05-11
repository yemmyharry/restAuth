package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(32) not null"`
	Email    string `json:"email" gorm:"type:varchar(32) unique not null"`
	Password string `json:"password" gorm:"type:varchar(256) not null"`
}
