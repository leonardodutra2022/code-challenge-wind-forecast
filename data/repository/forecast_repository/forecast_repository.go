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

/*
Função responsável por criar um registro em banco de dados para Previsão do tempo
*/
func (f Repository) Create(forecast *model.Forecast) error {
	return f.DBGo.Create(forecast).Error
}

/*
Função responsável por obter todos os registros em banco de dados sobre a previsão do tempo
*/
func (f Repository) GetAll() (*[]model.Forecast, error) {
	forecasts := []model.Forecast{}
	err := f.DBGo.Find(&forecasts).Error
	return &forecasts, err
}

/*
Função responsável por obter alertas por status em banco de dados sobre a previsão do tempo
*/
func (f Repository) GetAlertByStatus(status bool) (*[]model.Forecast, error) {
	forecasts := []model.Forecast{}
	err := f.DBGo.Where(&model.Forecast{Alerta: status}).Find(&forecasts).Error
	return &forecasts, err
}
