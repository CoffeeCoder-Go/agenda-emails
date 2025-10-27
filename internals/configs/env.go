package configs

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitEnv(){
	files,err := filepath.Glob("./*.env")

	if err != nil {
		log.Fatalf("[Configuração]%v\n",err.Error())
	}

	if len(files) == 0 {
		log.Fatalf("[No Exist File]%v\n",err.Error())
	}

	err = godotenv.Load(files...)

	if err != nil {
		log.Fatalf("[ENV]%v\n",err.Error())
	}

	log.Printf("Arquivos carregados:%v\n",files)
}