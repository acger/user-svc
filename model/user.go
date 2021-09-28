package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account      string `gorm:"size:30;unique"`
	Name         string `gorm:"size:30"`
	Avatar       string `gorm:"default:avatar/0.jpg"`
	Password     string
	Gender       string `gorm:"size:10"`
	Status       string `gorm:"size:20"`
	Level        string `gorm:"size:20"`
	Birthday     string `gorm:"size:10"`
	Contact      string
	Introduction string `gorm:"size:1000"`
}
