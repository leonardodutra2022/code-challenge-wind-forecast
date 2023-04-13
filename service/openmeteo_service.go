package service

import (
	"errors"
	"strings"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/api/openmeteo_api"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
)

/*
Função responsável pela requisição a API externa, no endpoint forecast, passando alguns parâmetros
*/
func GetForecast(latitude float64, longitude float64) (int, input_data.ForecastInput, error) {
	var forecastInput input_data.ForecastInput
	currentWeather := true
	endpoint := "forecast"
	reqParam := strings.Join([]string{"?latitude=", utils.Float64ToString(latitude), "&longitude=", utils.Float64ToString(longitude), "&current_weather=", utils.BoolToString(currentWeather)}, "")
	dataApi, statusCode, err := openmeteo_api.GetDataRequest(endpoint + reqParam)
	errConvert := utils.ByteToJson(dataApi, &forecastInput)
	if errConvert != nil {
		statusCode = 500
		err = errors.New("erro ao trater dados da API externa")
	}
	return statusCode, forecastInput, err
}
