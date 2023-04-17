package adapter_test

import (
	"testing"
	"time"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/adapter"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/stretchr/testify/assert"
)

func TestForecastInputToOutput(t *testing.T) {
	forecastInput := input_data.ForecastInput{
		Latitude:  100.0,
		Longitude: 200.0,
		Elevation: 10.0,
		Hourly: input_data.Hourly{
			Time:              []string{"2023-01-01"},
			Windspeed180m:     []float64{10.0},
			Winddirection180m: []float64{100.0},
		},
	}
	forecastOutput := adapter.ForecastInputToOutput(forecastInput)
	assert.Equal(t, 10.0, forecastOutput.WindSpeed, "deve ser equivalente o valor de 'velocidade do vento' já convertido/adaptado")
	assert.Equal(t, 100.0, forecastOutput.WindDirection, "deve ser equivalente o valor de 'direção do vento' já convertido/adaptado")
	assert.Equal(t, "2023-01-01", forecastOutput.DateTime, "deve ser equivalente o valor de 'data e hora' já convertido/adaptado")
}

func TestForecastToForecastAlertOutput(t *testing.T) {
	forecast := model.Forecast{
		Vel:       10.0,
		Dir:       100.0,
		Alerta:    false,
		Data:      time.Now(),
		UpdatedAt: time.Now(),
	}
	forecastAlertOutput := adapter.ForecastToForecastAlertOutput(forecast)
	assert.Equal(t, 10.0, forecastAlertOutput.WindSpeed, "deve ser equivalente o valor de 'velocidade do vento' já convertido/adaptado")
	assert.Equal(t, 100.0, forecastAlertOutput.WindDirection, "deve ser equivalente o valor de 'direção do vento' já convertido/adaptado")
	assert.Equal(t, false, forecastAlertOutput.Alert, "deve ser equivalente o valor de 'alerta' em boleano já convertido/adaptado")
}
