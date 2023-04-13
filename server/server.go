package server

import (
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

/*
Inicializando servidor web para api rest
*/
func NewServer() Server {
	cfg := config.Config{}
	env.Parse(&cfg)
	return Server{
		port:   cfg.PortApi,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Run(":" + s.port)
}
