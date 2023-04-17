package main

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
)

func main() {
	go database.StartDB(false, config.Config{})   // inicializando conexão com DB
	go service.CheckForecast(input_data.Hourly{}) // inicializando serviço em background que verifica a API periodicamente
	server := server.NewServer()                  // definindo configuração para o serviço Rest API
	server.Run()                                  // inicializando o serviço Rest
}
