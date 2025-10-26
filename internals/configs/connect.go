package configs

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	var err error
	DB,err = gorm.Open(sqlite.Open("./db/development.db"),&gorm.Config{})

	if err != nil {
		log.Fatal(errors.New("Não foi possivel se conectar por causa de:"+err.Error()))
	}
}