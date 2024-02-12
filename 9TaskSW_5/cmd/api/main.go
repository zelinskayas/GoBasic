package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/zelinskayas/GoBasic/9TaskSW_5/internal/app/api"
	"log"
	"os"
)

var (
	configPath   string //= "configs/api.toml"
	configFormat string //.toml/.env
)

func init() {
	//скажем что наше приложение будет получать путь до конфиг файла из командной строки
	flag.StringVar(&configFormat, "format", ".toml", "format to config file in .toml/.env")
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml/.env format")
}

func main() {
	//в этот момент происходит инициализация переменной configPath  значением
	flag.Parse()
	log.Println("it api works")
	//server instance initialization
	config := api.NewConfig()
	//тут читаем из .toml/.env

	log.Println("configFormat:", configFormat)
	if configFormat == " .env" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("could not find .env file:", err)
		}
		config.BindAddr = os.Getenv("app_port")
		config.LoggerLevel = os.Getenv("logger_level")
	} else {
		//десериализация содержимого .toml файла
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("can not find configs file. using default values:", err)
		}
	}
	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}
