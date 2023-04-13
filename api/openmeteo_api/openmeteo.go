package openmeteo_api

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

var API string = "https://api.open-meteo.com/v1/"

/*
Método privado dedicado a fazer chamada a API vai http e obter os dados em formato binário, para posterior tratamento
*/
func request(method string, endpoint string) (*http.Response, error) {
	clientHttp := http.Client{}
	req, err := http.NewRequest(method, strings.Join([]string{API, endpoint}, "/"), nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		err = errors.New("erro ao criar requisição")
		return nil, err
	}
	res, err := clientHttp.Do(req)
	if err != nil {
		err = errors.New("erro ao conectar-se com a API Open-Meteo")
		return nil, err
	}

	return res, nil
}

/*
Função responsável por fazer requisição a API
*/
func GetDataRequest(endpoint string) ([]byte, int, error) {
	res, err := request(http.MethodGet, endpoint)
	dataApi, _ := io.ReadAll(res.Body)
	return dataApi, res.StatusCode, err
}
