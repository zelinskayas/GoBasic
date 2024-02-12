package storage

type Config struct {
	//строка подключения к бд
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
