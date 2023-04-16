package service

import (
	"errors"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/api/openmeteo_api"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
)

/*
Função responsável pela requisição a API externa, no endpoint forecast, passando alguns parâmetros
*/
func GetForecastApi(latitude float64, longitude float64) (int, input_data.ForecastInput, error) {
	var forecastInput input_data.ForecastInput
	endpoint := "forecast"
	reqParam := strings.Join([]string{"?latitude=", utils.Float64ToString(latitude), "&longitude=", utils.Float64ToString(longitude), "&hourly=", "windspeed_180m,winddirection_180m"}, "")
	dataApi, statusCode, err := openmeteo_api.GetDataRequest(endpoint + reqParam)
	errConvert := utils.ByteToJson(dataApi, &forecastInput)
	if errConvert != nil {
		statusCode = 500
		err = errors.New("erro ao tratar dados da API externa")
	}
	return statusCode, forecastInput, err
}

/*
Função executada em background em que frequentemente executa a cada 5 minutos para requisitar API externa e fazer verificações para persistir em banco de dados
*/
func CheckForecast(testing bool) error {
	cfg := config.Config{}
	env.Parse(&cfg)
	for {
		_, forecastInput, err := GetForecastApi(cfg.LatitudeMonitor, cfg.LongitudeMonitor)
		isAlert, windSpeedForecast := IsThereAlert(forecastInput.Hourly)
		if err == nil {
			if isAlert {
				if AddForecast(forecastInput.Hourly, windSpeedForecast) != nil {
					return errors.New("erro ao registrar informação de alerta no banco de dados")
				}
			}
		} else {
			return errors.New("erro ao requisitar informações da API")
		}
		if testing {
			time.Sleep(time.Duration(5 * 1e9)) // alterando para 5 segundos o tempo periódico para requisitar API
			break
		} else {
			time.Sleep(time.Duration(cfg.CheckTimeInSeconds * 1e9))
		}
	}
	return nil
}

/*
Função que obtem os últimos registros (em 7 dias a frente como previsão) e com base em alguns parâmetros retorna se os dados indicam uma situação de alerta para a previsão obtida
*/
func IsThereAlert(fc input_data.Hourly) (bool, float64) {
	windDirLen := len(fc.Winddirection180m)
	windSpLen := len(fc.Windspeed180m)
	return (fc.Winddirection180m[windDirLen-1] >= 130 && fc.Winddirection180m[windDirLen-1] <= 260 && fc.Windspeed180m[windSpLen-1] > 15), fc.Windspeed180m[windSpLen-1]
}

/*
Função para realizar o registro em banco de dados como alerta indicado pela função isThereAlert
*/
func AddForecast(forecastHourly input_data.Hourly, windSpeedForecast float64) error {
	db := database.GetDatabase()
	forecast := model.Forecast{}
	forecast.Alerta = windSpeedForecast > 20
	forecast.Data = utils.DateStringToTime(forecastHourly.Time[len(forecastHourly.Time)-1])
	forecast.Dir = forecastHourly.Winddirection180m[len(forecastHourly.Winddirection180m)-1]
	forecast.Vel = forecastHourly.Windspeed180m[len(forecastHourly.Windspeed180m)-1]
	if err := db.Create(&forecast).Error; err != nil {
		return errors.New("erro ao cadastrar alerta de tempestade")
	}
	return nil
}
