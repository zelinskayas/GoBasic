package api

import (
	"github.com/sirupsen/logrus"
	"github.com/zelinskayas/GoBasic/9TaskSW_5/internal/app/models"
	"github.com/zelinskayas/GoBasic/9TaskSW_5/storage"
)

// пытаемся откунфигурировать наш API инстанс, а конкретнее поле Logger
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// пытаемся откунфигурировать маршрутизатор, а конкретнее поле router API
func (a *API) configureRouterField() {
	a.router.HandleFunc("/solve", a.GetSolve).Methods("GET")
	a.router.HandleFunc("/grab", a.PostGrab).Methods("POST")

	//a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello! This is good rest api!!!"))
	//})
}

// пытаемся конфигурировать наше хранилище storage API
func (a *API) configureStorageField() error {
	db := models.NewEquation()
	storage := storage.New(a.config.Storage, db)
	a.storage = storage
	return nil
}
