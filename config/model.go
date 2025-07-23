package config

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Port int `json:"port"`
}
