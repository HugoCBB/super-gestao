package config

import (
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar arquivo env")
		return err
	}

	return nil

}
