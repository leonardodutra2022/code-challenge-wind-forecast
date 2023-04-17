package model

import "time"

/*
Estrutura de dados para persistência em banco de dados, sobre o registro de data e hora de quando foi consultado a API
*/
type QueryApi struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	DateLastQueryApi time.Time `json:"last_queryapi"` // Data e hora da ultima consulta a API open-meteo.com
}
