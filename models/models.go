package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Id         int    `json:"id"`
	Course     int    `json:"course"`
	Department string `json:"department"`
	Phone      int    `json:"phone"`
	Adress     string `json:"adress"`
}
