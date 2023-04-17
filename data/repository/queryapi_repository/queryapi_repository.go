package queryapi_repository

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
Função responsável por criar ou obter o registro no banco de dados do registro data e hora de consulta da API
*/
func (f Repository) FirstOrCreate(qrApi *model.QueryApi) (*model.QueryApi, error) {
	qrFirst := model.QueryApi{}
	err := f.DBGo.FirstOrCreate(qrApi).First(&qrFirst).Error
	return &qrFirst, err
}

/*
Função responsável por obter o registro em banco de dados de data e hora da consulta a API
*/
func (f Repository) GetOne() (*model.QueryApi, error) {
	qrApi := model.QueryApi{}
	err := f.DBGo.Find(&qrApi).Error
	return &qrApi, err
}

/*
Função responsável por atualizar o registro em banco de dados de data e hora da consulta a API
*/
func (f Repository) Updates(qrApi *model.QueryApi) (*model.QueryApi, error) {
	qrFirst := model.QueryApi{}
	err := f.DBGo.Updates(qrApi).First(&qrFirst).Error
	return &qrFirst, err
}
