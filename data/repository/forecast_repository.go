package repository

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"gorm.io/gorm"
)

type IForecastRepository interface {
	IRepository
}

type Repository struct {
	DBGo *gorm.DB
}

func (f Repository) Create(forecast *model.Forecast) error {
	return f.DBGo.Create(forecast).Error
}

func (f Repository) GetAll() (*[]model.Forecast, error) {
	forecasts := []model.Forecast{}
	err := f.DBGo.Find(&forecasts).Error
	return &forecasts, err
}
