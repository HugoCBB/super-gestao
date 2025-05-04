package database

import (
	"fmt"
	"log"
	"os"

	"github.com/HugoCBB/super-gestao/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDataBase() {

	dsn := os.Getenv("DB")

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Erro ao se conectar ao banco de dados")
		return
	}

	fmt.Println("Conexao estabelecida com o banco de dados")
	DB.AutoMigrate(&models.User{}, &models.Task{})

}
