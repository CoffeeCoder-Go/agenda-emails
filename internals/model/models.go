package model

import (
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	Surname string `form:"surname"`
	Email string `form:"email"`
	ClassificationID uint `form:"classification_id"`
}

type Classification struct {
	gorm.Model
	Name string
	Color string `form:"color"`
	Emails []Email `form:"emails"`
}

