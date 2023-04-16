package main

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/service"
)

func main() {
	go database.StartDB()
	go service.CheckForecast(false)
	server := server.NewServer()
	server.Run()
}
