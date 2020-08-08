package main

import (
	"photon-server/cmd/logger"
	"photon-server/cmd/server"
	"photon-server/pkg/config"
	"photon-server/pkg/database"

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
