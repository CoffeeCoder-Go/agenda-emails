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
	Apelido string `json:"apelido"`
	Email string `json:"email"`
	ClassificacaoID uint `json:"classificacao_id"`
}

type Classification struct {
	gorm.Model
	Color Color `gorm:"embedded" json:"color"`
	Emails []Email `gorm:"foreignKey:ClassificacaoID" json:"emails"`
}

