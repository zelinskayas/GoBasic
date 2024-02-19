package api

import "github.com/zelinskayas/GoBasic/goodRestApiWithDBAuth/storage"

// general instance for API server of REST application
type Config struct {
	//port
	BindAddr string `toml:"bind_addr"`
	//Logger Level
	LoggerLevel string `toml:"logger_level"`
	//Storage configs
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
