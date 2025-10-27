package configs

import (
	"errors"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var err error
	DB,err = gorm.Open(sqlite.Open(os.Getenv("DB")),&gorm.Config{})

	if err != nil {
		log.Fatal(errors.New("Não foi possivel se conectar por causa de:"+err.Error()))
	}

	log.Printf("[DB]Conectado com sucesso!")
}