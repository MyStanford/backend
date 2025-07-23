package config

import (
	"encoding/json"
	"mystanford/logger"
	"mystanford/utils"
	"os"
	"path/filepath"
)

func InitConfig() {
	rootPath, _ := os.Getwd()
	if !utils.FileExist(filepath.Join(rootPath, "data", "config.json")) {
		logger.Logger.Info("配置文件不存在，正在新建配置文件...")
		NowConfig = Config{
			Server: ServerConfig{
				Port: 3000,
			},
		}
		newConfig, _ := json.MarshalIndent(NowConfig, "", "    ")
		e := os.WriteFile(filepath.Join(rootPath, "data", "config.json"), newConfig, os.ModePerm)
		if e != nil {
			logger.Logger.Error(e.Error())
			return
		}
		return
	}
	logger.Logger.Info("正在加载配置文件...")
	configContent, e := os.ReadFile(filepath.Join(rootPath, "data", "config.json"))
	if e != nil {
		logger.Logger.Error(e.Error())
		return
	}
	e = json.Unmarshal(configContent, &NowConfig)
	if e != nil {
		logger.Logger.Error(e.Error())
		return
	}
}
