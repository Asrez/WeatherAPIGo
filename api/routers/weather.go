package routers

import (
	handlers "github.com/Asrez/WeatherAPIGo/api/handler"
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/gin-gonic/gin"
)


func Weather(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewWeatherHandler(cfg)
	router.GET("/current" , h.Current)
	router.GET("/forecast",h.Forecast)
}