package main

import "github.com/leonardodutra2022/code-challenge-wind-forecast/server"

func main() {
	server := server.NewServer()
	server.Run()
}
