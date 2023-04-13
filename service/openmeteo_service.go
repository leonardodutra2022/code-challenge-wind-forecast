package service

import (
	"strings"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/api/openmeteo_api"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
)

func GetForecast(latitude float64, longitude float64) (int, input_data.ForecastInput, string, error) {
	var forecastInput input_data.ForecastInput
	endpoint := "forecast"
	reqParam := strings.Join([]string{"?latitude=", utils.Float64ToString(latitude), "&longitude=", utils.Float64ToString(longitude)}, "")
	dataApi, statusCode, errorMessage, err := openmeteo_api.GetDataRequest(endpoint + reqParam)
	errConvert := utils.ByteToJson(dataApi, &forecastInput)
	if errConvert != nil {
		statusCode = 500
		errorMessage = "Erro ao trater dados da API externa"
	}
	return statusCode, forecastInput, errorMessage, err
}
