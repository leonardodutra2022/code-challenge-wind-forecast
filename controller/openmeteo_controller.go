package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/adapter"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/output_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
)

/*
Estrutura de dados para tratar informações obtidas Query Param
*/
type QueryString struct {
	Latitude  float64 `form:"latitude"`
	Longitude float64 `form:"longitude"`
}

/*
Função controller para a rota forecast da API local
*/
func GetForecast(c *gin.Context) {
	var queryParams QueryString
	if c.ShouldBindQuery(&queryParams) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Parametro não informado",
			"latitude_param":  queryParams.Latitude,
			"longitude_param": queryParams.Longitude,
		})
		return
	}
	var forecastOutput output_data.ForecastOutput
	statusCode, forecastInput, err := service.GetForecastApi(queryParams.Latitude, queryParams.Longitude)
	forecastOutput = adapter.ForecastInputToOutput(forecastInput)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}
	queryApiLastDate, err := service.FindLastQueryApi()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	forecastOutput.DateLastQueryApi = queryApiLastDate.DateLastQueryApi.Format(time.RFC822Z)
	forecastOutput.DateTime = utils.DateStringToTime(forecastOutput.DateTime).Format(time.RFC822Z)
	c.JSON(statusCode,
		forecastOutput,
	)
}

func GetAlerts(c *gin.Context) {
	forecastAlerts, err := service.GetForecastAlerts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,
		forecastAlerts,
	)
}
