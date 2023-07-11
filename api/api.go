package api

import (
	"fmt"

	"github.com/Asrez/WeatherAPIGo/api/routers"
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/Asrez/WeatherAPIGo/packge/logging"
	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	RegisterRoutes(r,cfg)
	r.Use(gin.Logger(), gin.Recovery())

	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// User
		users := v1.Group("/users")
		routers.User(users, cfg)
	}
}