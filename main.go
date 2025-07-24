// Author :  dinglz
// Time   :  2025.07.24

package main

import (
	"mystanford/config"
	"mystanford/database"
	"mystanford/folder"
	"mystanford/logger"
	"mystanford/server"
)

func init() {
	logger.InitLogger()
	folder.InitFolder()
	config.InitConfig()
	database.InitDatabase()
	server.InitServer()
}

func main() {
	logger.Logger.Info("Starting My Stanford Backend...")
	logger.Logger.Info("Dream start from here!")

	server.Start(config.NowConfig.Server.Port)
}
