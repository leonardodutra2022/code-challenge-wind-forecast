package routes

import (
	"github.com/caarlos0/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/controller"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
)

/*
Função responsável pela configuração de Cors, grupos de Api, middleware, segurança da API, contexto, headers e métodos gerais permitidos
*/
func ConfigRoutes() *gin.Engine {
	cfg := config.Config{}
	env.Parse(&cfg)
	if cfg.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "authorization", "x-requested-with"},
		AllowCredentials: true,
	}))
	routes := router.Group("api")
	{
		routes.GET("/previsao", controller.GetForecast)
		routes.GET("/checkHealth", controller.CheckHealth)
	}

	return router
}
