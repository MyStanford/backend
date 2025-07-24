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
			Models: []ModelConfig{
				{
					Type:  "openai",
					Name:  "kimik2",
					Model: "kimi-k2-0711-preview",
					Path:  "https://api.moonshot.cn/v1",
					Key:   "示例模型配置，需要使用k2请上官网申请",
				},
			},
			Database: DatabaseConfig{
				Type: "sqlite",
				Dsn:  "data/data.db",
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
