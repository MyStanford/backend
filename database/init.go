package database

import (
	"errors"
	"mystanford/config"
	"mystanford/logger"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	var e error
	switch config.NowConfig.Database.Type {
	case "sqlite":
		DB, e = gorm.Open(sqlite.Open(config.NowConfig.Database.Dsn))
	case "mysql":
		DB, e = gorm.Open(mysql.Open(config.NowConfig.Database.Dsn))
	default:
		e = errors.New("不支持的数据库类型: " + config.NowConfig.Database.Type + " !")
	}
	if e != nil {
		logger.Logger.Error("connect database error:", e.Error())
		os.Exit(1)
	}
	e = DB.AutoMigrate(&Person{})
	if e != nil {
		logger.Logger.Error("migrate error:", e.Error())
		os.Exit(1)
	}
	PersonLoadDefault()
}
