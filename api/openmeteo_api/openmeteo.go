package openmeteo_api

import (
	"io"
	"net/http"
	"strings"
)

var API string = "https://api.open-meteo.com/v1"

func request(method string, endpoint string) (*http.Response, string, error) {
	var errorMessage string
	clientHttp := http.Client{}
	req, err := http.NewRequest(method, strings.Join([]string{API, endpoint}, "/"), nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		errorMessage = "Erro ao criar requisição"
		return nil, errorMessage, err
	}
	res, err := clientHttp.Do(req)
	if err != nil {
		errorMessage = "Erro ao conectar-se com a API Open-Meteo"
		return nil, errorMessage, err
	}

	return res, "", nil
}

func GetDataRequest(endpoint string) ([]byte, int, string, error) {
	requestMethod := "GET"
	res, errorMessage, err := request(requestMethod, endpoint)
	dataApi, _ := io.ReadAll(res.Body)
	return dataApi, res.StatusCode, errorMessage, err
}
