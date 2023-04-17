package adapter

import (
	"time"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/output_data"
)

/*
Adapter para duas estrutura de dados, onde uma há tipos de dados distintos para input de dados output de dados
*/
func ForecastInputToOutput(forecastInput input_data.ForecastInput) output_data.ForecastOutput {
	forecastOutput := output_data.ForecastOutput{}
	forecastHourly := forecastInput.Hourly
	forecastOutput.DateTime = forecastHourly.Time[len(forecastHourly.Time)-1]
	forecastOutput.WindDirection = forecastHourly.Winddirection180m[len(forecastHourly.Winddirection180m)-1]
	forecastOutput.WindSpeed = forecastHourly.Windspeed180m[len(forecastHourly.Windspeed180m)-1]
	return forecastOutput
}

/*
Adapter para duas estrutura de dados, onde uma há tipos de dados distintos para input de dados output de dados, nesse caso para o endpoint alerta
*/
func ForecastToForecastAlertOutput(forecastAlerts model.Forecast) output_data.ForecastAlertOutput {
	forecastOutput := output_data.ForecastAlertOutput{}
	forecastOutput.Alert = forecastAlerts.Alerta
	forecastOutput.WindDirection = forecastAlerts.Dir
	forecastOutput.WindSpeed = forecastAlerts.Vel
	forecastOutput.DateTime = forecastAlerts.Data.Format(time.RFC822Z)
	return forecastOutput
}
