package model

import (
	"gorm.io/gorm"
)

type Color struct {
	Red uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue uint8 `json:"blue"`
}

type Email struct {
	gorm.Model
	Surname string `json:"surname"`
	Email string `json:"email"`
	ClassificationID uint `json:"classification_id"`
}

type Classification struct {
	gorm.Model
	Color Color `gorm:"embedded" json:"color"`
	Emails []Email `json:"emails"`
}

