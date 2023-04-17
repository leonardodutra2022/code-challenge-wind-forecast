package forecast_repository

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/repository"
	"gorm.io/gorm"
)

type IForecastRepository interface {
	repository.IRepository
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

func (f Repository) GetAlertByStatus(status bool) (*[]model.Forecast, error) {
	forecasts := []model.Forecast{}
	err := f.DBGo.Where(&model.Forecast{Alerta: status}).Find(&forecasts).Error
	return &forecasts, err
}
