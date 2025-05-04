package main

import (
	"github.com/HugoCBB/super-gestao/config"
	"github.com/HugoCBB/super-gestao/internal/database"
	"github.com/HugoCBB/super-gestao/internal/routers"
)

func init() {
	config.LoadEnv()

}

func main() {
	database.ConnectDataBase()
	routers.HandleRequest()

}
