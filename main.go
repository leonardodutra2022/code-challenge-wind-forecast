package main

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server"
)

func main() {
	go database.StartDB()
	server := server.NewServer()
	server.Run()
}
