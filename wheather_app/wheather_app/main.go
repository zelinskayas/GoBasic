// main.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Передача города в запросе
	e.GET("/weather/:city", getWeather)

	// Старт сервера
	e.Start(":8080")
}

func getWeather(c echo.Context) error {
	city := c.Param("city")

	// TODO: Используйте город что бы отправить запрос в АПИ погоды и отобразить в ответе

	// TODO: Верните ответ в формате JSON
	return c.JSON(http.StatusOK, map[string]interface{}{
		"city": city,
	})
}
