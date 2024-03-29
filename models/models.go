package models

import "gorm.io/gorm"

type ResponseInformation struct {
	Message string `json:"message"`
}

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
