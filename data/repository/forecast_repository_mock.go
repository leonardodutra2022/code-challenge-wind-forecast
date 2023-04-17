package repository

import (
	"errors"
	"time"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
)

type IForecastRepositoryMock interface {
	IRepository
}

type RepositoryMock struct {
	DBGo []model.Forecast
}

var list []model.Forecast = []model.Forecast{
	{ID: 1, Vel: 30.0, Dir: 180.0, Alerta: true, Data: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Vel: 20.0, Dir: 160.0, Alerta: true, Data: time.Now(), UpdatedAt: time.Now()},
}

func (f RepositoryMock) Create(forecast model.Forecast) error {
	if len(f.DBGo) < 1 {
		f.DBGo = list
	}
	f.DBGo = append(f.DBGo, forecast)
	if len(f.DBGo) < 3 {
		return errors.New("objeto não adicionado ao slice")
	}
	return nil
}

func (f RepositoryMock) GetAll() (*[]model.Forecast, error) {
	if len(f.DBGo) < 1 {
		f.DBGo = list
	}
	if len(f.DBGo) > 0 {
		return &f.DBGo, nil
	}
	return &f.DBGo, errors.New("erro encontrado ao carregar lista")
}
