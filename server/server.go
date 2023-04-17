package server

import (
	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server/routes"
)

type Server struct {
	Port string
}

/*
Inicializando servidor web para api rest
*/
func NewServer() Server {
	cfg := config.Config{}
	env.Parse(&cfg)
	return Server{
		Port: cfg.PortApi,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes()
	router.Run(":" + s.Port)
}
