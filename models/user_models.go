package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"unique"`
	Password string
}
