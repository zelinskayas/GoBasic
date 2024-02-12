package api

import (
	"encoding/json"
	"github.com/zelinskayas/GoBasic/9TaskSW_5/internal/app/models"
	"log"
	"net/http"
)

// вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

// full API Handler initialization file
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// считать коэффициенты квадратного уравнения из тела запроса
func (api *API) PostGrab(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Grab POST /grab")

	var equation models.Equation
	err := json.NewDecoder(req.Body).Decode(&equation)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println(&equation, equation)
	a := api.storage.Equation().Grab(&equation)
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

// возвращает json с решением квадратного уравнения, A B C и Nroots
func (api *API) GetSolve(writer http.ResponseWriter, req *http.Request) {
	//инициализируем хедеры
	initHeaders(writer)
	//логируем момент начала обработки запроса
	api.logger.Info("Get Solve /solve")

	//пытаемся что-то получить от бд
	equation := api.storage.Equation().Solve()
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(equation)
}
