package db

import (
	"log"
	"time"
	"github.com/Asrez/WeatherAPIGo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitDbMysql(cfg *config.Config) error {
	var err error
	dsn := "root:root@tcp(localhost:9910)/?charset=utf8&parseTime=True&loc=Local"
	dbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Db connection established")
	return nil
}

func GetDbMysql() *gorm.DB {
	return dbClient
}

func CloseDbMysql() {
	con, _ := dbClient.DB()
	con.Close()
}