package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetForecast(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.ConfigRoutes()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/previsao", nil)
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "deve retornar http status code 200, indicando que a requisição ocorreu tudo bem")
	assert.Contains(t, res.Body.String(), "velocidade_vento", "deve conter o atributo 'velocidade_vento' no endpoint")
	assert.Contains(t, res.Body.String(), "direcao_vento", "deve conter o atributo 'direcao_vento' no endpoint")
	assert.Contains(t, res.Body.String(), "previsao_datahora", "deve conter o atributo 'previsao_datahora' no endpoint")
}
