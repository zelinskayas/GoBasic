package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zelinskayas/GoBasic/9TaskSW_5/storage"
	"log"
	"net/http"
)

// base API server instance description
type API struct {
	//UNEXPORTED FIELD!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//добавление поля для работы с хранилищем
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// start http server/configure Loggers, router, database connection and etc...
func (api *API) Start() error {
	log.Println("start 1: ")
	//trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//подтверждение того что логгер сконфигурирован
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	//кофигурируем марщрутизатор
	api.configureRouterField()

	//конфигурируем хранилище
	if err := api.configureStorageField(); err != nil {
		return err
	}

	//на этапе валидного завершения стартуем http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
