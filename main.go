package main

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/input_data"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
)

func main() {
	go database.StartDB(false, config.Config{})
	go service.CheckForecast(false, input_data.Hourly{})
	server := server.NewServer()
	server.Run()
}
