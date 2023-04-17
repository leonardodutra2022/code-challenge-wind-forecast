package service

import (
	"errors"
	"time"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/repository/queryapi_repository"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/database"
)

/*
Função responsável para consulta registro de data da última consulta a API
*/
func FindLastQueryApi() (model.QueryApi, error) {
	repo := queryapi_repository.Repository{DBGo: database.GetDatabase()}
	queryApi, err := repo.GetOne()
	if err != nil {
		return model.QueryApi{}, errors.New("erro ao obter registro em banco de dados")
	} else {
		if queryApi.ID < 1 {
			queryApi, err = repo.FirstOrCreate(&model.QueryApi{
				DateLastQueryApi: time.Now(),
			})
			return *queryApi, err
		} else {
			queryApi, err = repo.Updates(&model.QueryApi{
				ID:               queryApi.ID,
				DateLastQueryApi: time.Now(),
			})
			return *queryApi, err
		}
	}
}
