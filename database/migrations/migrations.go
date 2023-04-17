package migrations

import (
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/model"
	"gorm.io/gorm"
)

/*
Função para criação de tabelas, atualização de estrutura de tabela, ou pré SQL a serem executados no banco de dados
*/
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.Forecast{}, &model.QueryApi{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Forecast{}, &model.QueryApi{})
}
