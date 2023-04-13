package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/adapter"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/output_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
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
		c.JSON(400, gin.H{
			"error":           "Parametro não informado",
			"latitude_param":  queryParams.Latitude,
			"longitude_param": queryParams.Longitude,
		})
		return
	}

	var forecastOutput output_data.ForecastOutput
	statusCode, forecastInput, err := service.GetForecast(queryParams.Latitude, queryParams.Longitude)
	forecastOutput = adapter.ForecastInputToOutput(forecastInput)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(statusCode,
		forecastOutput,
	)
}
