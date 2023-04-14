package service

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/adapter"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/api/openmeteo_api"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
)

/*
Função responsável pela requisição a API externa, no endpoint forecast, passando alguns parâmetros
*/
func GetForecastApi(latitude float64, longitude float64) (int, input_data.ForecastInput, error) {
	var forecastInput input_data.ForecastInput
	currentWeather := true
	endpoint := "forecast"
	reqParam := strings.Join([]string{"?latitude=", utils.Float64ToString(latitude), "&longitude=", utils.Float64ToString(longitude), "&current_weather=", utils.BoolToString(currentWeather)}, "")
	dataApi, statusCode, err := openmeteo_api.GetDataRequest(endpoint + reqParam)
	errConvert := utils.ByteToJson(dataApi, &forecastInput)
	if errConvert != nil {
		statusCode = 500
		err = errors.New("erro ao tratar dados da API externa")
	}
	return statusCode, forecastInput, err
}

func CheckForecast() {
	cfg := config.Config{}
	env.Parse(&cfg)
	for {
		_, forecastInput, err := GetForecastApi(cfg.LatitudeMonitor, cfg.LongitudeMonitor)
		if err == nil && isThereAlert(forecastInput) {
			if addForecast(forecastInput) != nil {
				log.Fatalln(err.Error())
			}
		}
		time.Sleep(5 * time.Minute)
	}
}

func isThereAlert(forecast input_data.ForecastInput) bool {
	return (forecast.CurrentWeather.WindDirection >= 130 && forecast.CurrentWeather.WindDirection <= 260 && forecast.CurrentWeather.WindSpeed > 9)
}

func addForecast(forecastInput input_data.ForecastInput) error {
	db := database.GetDatabase()
	forecast := adapter.ForecastInputToForecast(forecastInput)
	if err := db.Create(&forecast).Error; err != nil {
		return errors.New("erro ao cadastrar alerta de tempestade")
	}
	return nil
}
