package adapter

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/output_data"
)

func ForecastInputToOutput(forecastInput input_data.ForecastInput) output_data.ForecastOutput {
	var forecastOutput output_data.ForecastOutput
	forecastOutput.DateTime = forecastInput.CurrentWeather.DateTime
	forecastOutput.WindDirection = forecastInput.CurrentWeather.WindDirection
	forecastOutput.WindSpeed = forecastInput.CurrentWeather.WindSpeed
	return forecastOutput
}

func ForecastInputToForecast(forecastInput input_data.ForecastInput) model.Forecast {
	var forecast model.Forecast
	forecast.Dir = forecastInput.CurrentWeather.WindDirection
	forecast.Vel = forecastInput.CurrentWeather.WindSpeed
	return forecast
}
