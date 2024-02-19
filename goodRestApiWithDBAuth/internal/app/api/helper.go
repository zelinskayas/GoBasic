package api

import (
	"github.com/sirupsen/logrus"
	"github.com/zelinskayas/GoBasic/goodRestApiWithDBAuth/internal/app/middleware"
	"github.com/zelinskayas/GoBasic/goodRestApiWithDBAuth/storage"
	"net/http"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")

	//было до jwt
	//a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	//теперь требует наличие jwt
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods("GET")

	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")

	//new pair for auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostToAuth).Methods("POST")
	/*заглушка
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is good rest api!!!"))
	})
	*/
}

// пытаемся конфигурировать наше хранилище storage API
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//пытаемся установить соединение если невозможно - возвращаем ошибку
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}

//  migrate -path C:/Users/Мир/GolandProjects/github.com/zelinskayas/GoBasic/goodRestApiWithDB/migrations -database "sqlserver://localhost:1433" up
