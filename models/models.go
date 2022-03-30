package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	ID         int    `json:"uid"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Course     int    `json:"course"`
	Department string `json:"department"`
	Phone      int    `json:"phone"`
	Adress     string `json:"adress"`
	Email      string `json:"email"`
}
