package database

import (
	"log"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database/migrations"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	cfg := config.Config{}
	env.Parse(&cfg)
	strConnection := strings.Join([]string{string("host=" + cfg.HostDB), string("port=" + utils.IntToString(cfg.PortDB)), string("user=" + cfg.UserDB), string("dbname=" + cfg.DatabaseName), string("sslmode=" + cfg.SSLMode), string("password=" + cfg.PassDB)}, " ")
	database, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error in database connection: ", err.Error())
	}

	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
