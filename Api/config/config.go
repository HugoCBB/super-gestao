package config

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Erro ao carregar arquivo env")
		return err
	}

	return nil

}
