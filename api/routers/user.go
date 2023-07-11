package routers

import (
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/Asrez/WeatherAPIGo/api/handler"
	"github.com/gin-gonic/gin"
)


func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	router.POST("/login", h.LoginByUsername)
	router.POST("/register", h.RegisterByUsername)
}