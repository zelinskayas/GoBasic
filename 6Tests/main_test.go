package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	InputData int //то что на вход подается
	Answer    int //то что вернет функция
	Expected  int //то что ожидаем получить
}

var Cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 3,
		Expected:  6,
	},
	{
		InputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("test case %d failed: input %v result %v expected %v", id, test.InputData, test.Answer, test.Expected)
		}
	}

}

type HttpTestCase struct {
	Name     string //имя теста
	Numeric  int    //значение которое будет передаваться в hhtp запрос
	Expected []byte //http response, который ожидаем увидеть
}

// тестовый сценарий для http POST запроса
// для каждого запроса делается свой сценарий
var HttpCases = []HttpTestCase{
	{
		Name:     "first test",
		Numeric:  1,
		Expected: []byte("1"),
	},
	{
		Name:     "second test",
		Numeric:  3,
		Expected: []byte("6"),
	}, {
		Name:     "third test",
		Numeric:  5,
		Expected: []byte("120"),
	},
}

func TestHandleFactorial(t *testing.T) {
	for _, test := range HttpCases {
		//под тест (суб тест)
		t.Run(test.Name, func(t *testing.T) {
			handler := http.HandlerFunc(HandlerFactorial)
			recorder := httptest.NewRecorder()
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil) //какой будет запрос
			//data := io.Reader([]byte(`{"num" :6}`))
			//request, err := http.Post("http://localhost:8080/factorial?num=6", "application/json", data)
			if err != nil {
				t.Error(err)
			}
			handler.ServeHTTP(recorder, request) //выполняем запрос и ответ записываем в recorder
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input: %v! result: %v expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected),
				)
			}
		}) //под-тестовый раннер
	}
}
