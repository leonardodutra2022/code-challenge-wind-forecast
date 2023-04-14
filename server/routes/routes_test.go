package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/server/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetForecast(t *testing.T) {
	router := routes.ConfigRoutes()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/previsao", nil)
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "deve retornar http status code 200, indicando que a requisição ocorreu tudo bem")
	assert.Contains(t, res.Body.String(), "windspeed", "deve conter o atributo 'Wind Speed' no endpoint")
	assert.Contains(t, res.Body.String(), "winddirection", "deve conter o atributo 'Wind Direction' no endpoint")
	assert.Contains(t, res.Body.String(), "date_time", "deve conter o atributo 'DateTime' no endpoint")
}
