package migrations

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Forecast{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Forecast{})
}
