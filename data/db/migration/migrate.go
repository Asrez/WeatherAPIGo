package migration

import (
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/Asrez/WeatherAPIGo/constants"
	"github.com/Asrez/WeatherAPIGo/data/db"
	"github.com/Asrez/WeatherAPIGo/data/models"
	"github.com/Asrez/WeatherAPIGo/packge/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up() {
	database := db.GetDb()
	createTables(database)
	createDefaultUserInformation(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// User
	tables = addNewTable(database, models.User{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {
	u := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "09111112222", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u)

}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
	}
}

func Down() {

}