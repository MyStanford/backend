package config

type Config struct {
	Server   ServerConfig   `json:"server"`
	Models   []ModelConfig  `json:"models"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type ModelConfig struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Model string `json:"model"`
	Path  string `json:"path"`
	Key   string `json:"key"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Dsn  string `json:"dsn"`
}
