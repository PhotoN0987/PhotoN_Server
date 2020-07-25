package main

import (
	"Photon_Server/cmd/logger"
	"Photon_Server/cmd/server"
	"Photon_Server/pkg/config"
	"Photon_Server/pkg/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger.LoggingSetting()
	cfg := config.NewConfig()
	db := database.NewDB(cfg)

	server := server.NewServer(cfg, db)
	server.SetUpRouter()
	server.Run()
}
