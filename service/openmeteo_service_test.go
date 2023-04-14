package service_test

import (
	"testing"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
	"github.com/stretchr/testify/assert"
)

func TestIsThereAlert(t *testing.T) {
	obj := input_data.Hourly{
		Windspeed180m:     []float64{17.5, 18.0},
		Winddirection180m: []float64{161.1, 160.0},
	}
	isAlert, windSpeed := service.IsThereAlert(obj)
	assert.Equal(t, true, isAlert, "deve ser verdadeiro, considerando os parâmetros da função e considerando o ultimo valor dos slices do objeto (sendo o mais atual)")
	assert.Equal(t, 18.0, windSpeed, "deve ser 18.0 com base no último valor do slice Wind Speed 180m")
}
