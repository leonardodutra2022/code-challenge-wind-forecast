package server

import (
	"github.com/gin-gonic/gin"
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
	return Server{
		port:   "9000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	router.Run(":" + s.port)
}
