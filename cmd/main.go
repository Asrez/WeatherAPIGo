package main

import (
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/Asrez/WeatherAPIGo/data/db"
	"github.com/Asrez/WeatherAPIGo/api"
	"github.com/Asrez/WeatherAPIGo/data/db/migration"
	"github.com/Asrez/WeatherAPIGo/packge/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up()

	api.InitServer(cfg)
}