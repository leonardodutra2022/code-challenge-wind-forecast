package controller

import (
	"net/http"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/adapter"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/output_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
)

/*
Estrutura de dados para saída de dados
*/
type ResponseForecastOutput struct {
	DateLastQueryApi string                       `json:"ultima_consulta_api"`
	Forecasts        []output_data.ForecastOutput `json:"previsoes"`
}

type ResponseAlertsOutput struct {
	DateLastQueryApi string                            `json:"ultima_consulta_api"`
	Alerts           []output_data.ForecastAlertOutput `json:"alertas"`
}

/*
Função controller para a rota forecast (previsão) da API local
*/
func GetForecast(c *gin.Context) {
	cfg := config.Config{}
	env.Parse(&cfg)
	listAllForecasts, err := service.GetAllForecast()
	forecastsOutput := adapter.ForecastsToForecastOutput(listAllForecasts)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var queryApiLastDate model.QueryApi
	if cfg.IsTestMode {
		queryApiLastDate = model.QueryApi{
			ID:               1,
			DateLastQueryApi: time.Now(),
		}
	} else {
		queryApiLastDate, err = service.FindLastQueryApi()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var response ResponseForecastOutput
	response.DateLastQueryApi = queryApiLastDate.DateLastQueryApi.Format(time.RFC822Z)
	response.Forecasts = forecastsOutput
	c.JSON(http.StatusOK,
		response,
	)
}

/*
Função controller para a rota de alertas de previsão de tempestade
*/
func GetAlerts(c *gin.Context) {
	cfg := config.Config{}
	env.Parse(&cfg)
	var forecastAlerts []model.Forecast
	var err error
	if !cfg.IsTestMode {
		forecastAlerts, err = service.GetForecastAlerts()
	} else {
		forecastAlerts = []model.Forecast{
			{
				ID:     1,
				Vel:    10.,
				Dir:    100.0,
				Alerta: false,
				Data:   time.Now(),
			},
			{
				ID:     2,
				Vel:    20.,
				Dir:    200.0,
				Alerta: true,
				Data:   time.Now(),
			},
		}
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var queryApiLastDate model.QueryApi
	if cfg.IsTestMode {
		queryApiLastDate = model.QueryApi{
			ID:               1,
			DateLastQueryApi: time.Now(),
		}
	} else {
		queryApiLastDate, err = service.FindLastQueryApi()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	forecastAlertsOutputs := adapter.ForecastsToForecastAlertOutput(forecastAlerts)
	var response ResponseAlertsOutput
	response.DateLastQueryApi = queryApiLastDate.DateLastQueryApi.Format(time.RFC822Z)
	response.Alerts = forecastAlertsOutputs
	c.JSON(http.StatusOK,
		response,
	)
}
