package model

import "time"

type QueryApi struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	DateLastQueryApi time.Time `json:"last_queryapi"` // Data e hora da ultima consulta a API open-meteo.com
}
