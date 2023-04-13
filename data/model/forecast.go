package model

import "time"

type Forecast struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Vel       float64   `json:"vel"`        // Velocidade
	Dir       float64   `json:"dir"`        // Direção em graus
	Alerta    bool      `json:"alerta"`     // Alerta de velocidade do vento, quando superior a 20 km/h
	Data      time.Time `json:"data"`       // Data de registro da informação da API
	UpdatedAt time.Time `json:"updated_at"` // Data da atualização do registro localmente
}
