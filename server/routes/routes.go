package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/controller"
)

/*
Função responsável pela configuração de Cors, grupos de Api, middleware, segurança da API, contexto, headers e métodos gerais permitidos
*/
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "authorization", "x-requested-with"},
		AllowCredentials: true,
	}))
	routes := router.Group("api")
	{
		routes.GET("/forecast", controller.GetForecast)
	}

	return router
}
