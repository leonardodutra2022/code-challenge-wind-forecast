package service_test

import (
	"testing"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
	"github.com/stretchr/testify/assert"
)

func TestIsThereAlert(t *testing.T) {
	obj := input_data.Hourly{
		Windspeed180m:     []float64{17.5, 18.0},
		Winddirection180m: []float64{161.1, 160.0},
	}
	isAlert, windSpeed := service.IsThereAlert(obj)
	assert.Equal(t, true, isAlert, "deve ser verdadeiro, considerando os parâmetros da função e considerando o ultimo valor dos slices do objeto (sendo o mais atual)")
	assert.Equal(t, 18.0, windSpeed, "deve ser 18.0 com base no último valor do slice Wind Speed 180m")
}

func TestGetForecastApi(t *testing.T) {
	latitudeTest := 20.0
	longitudeTest := 60.0
	objForecastTest := input_data.ForecastInput{
		Latitude:  latitudeTest,
		Longitude: longitudeTest,
	}
	statusCode, responseForecast, err := service.GetForecastApi(latitudeTest, longitudeTest)
	assert.Nil(t, err, "deve ocorrer a requisição sem erro")
	assert.Equal(t, 200, statusCode, "deve retornar OK (200) a resposta da requisição a API")
	assert.Equal(t, objForecastTest.Latitude, responseForecast.Latitude, "deve ser igual o valor da latitude, indicando retorno do ojeto com dados")
	assert.Equal(t, objForecastTest.Longitude, responseForecast.Longitude, "deve ser igual o valor da longitude, indicando retorno do ojeto com dados")

	statusCode2, _, _ := service.GetForecastApi(-10000000, 150000000)
	assert.Equal(t, 400, statusCode2, "deve retornar status code 400")

}

func TestCheckForecast(t *testing.T) {
	obj := input_data.Hourly{
		Windspeed180m:     []float64{17.5, 18.0},
		Winddirection180m: []float64{161.1, 160.0},
	}
	errTestCheckForecast := service.CheckForecast(obj)
	assert.Nil(t, errTestCheckForecast, "deve ocorrer a checkagem de alerta sem erro")

	checkForecastAlert := service.CheckForecast(obj)
	assert.Nil(t, checkForecastAlert, "deve ocorrer adição de alerta sem erro")
}

func TestAddForecast(t *testing.T) {
	objTest := input_data.Hourly{
		Windspeed180m:     []float64{10.0, 18.0, 22.0},
		Winddirection180m: []float64{180.0, 175.0, 190.0},
		Time:              []string{"2023-01-01T05:00", "2023-01-01T06:00", "2023-01-01T07:00"},
	}

	err := service.AddForecast(objTest, 21)
	assert.Nil(t, err, "deve ocorrer registro do alerta sem erro")
}

func TestGetForecastAlerts(t *testing.T) {
	list, err := service.GetForecastAlerts()
	assert.Nil(t, err, "deve obter a lista sem erro")
	assert.Len(t, list, 2, "deve conter 2 itens na lista")
}

func TestGetAllForecast(t *testing.T) {
	list, err := service.GetAllForecast()
	assert.Nil(t, err, "deve obter a lista sem erro")
	assert.Len(t, list, 2, "deve conter 2 itens na lista")
}
