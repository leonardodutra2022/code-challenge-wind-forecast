package output_data

/*
Estrutura de dados que representa um conjunto de dados de saída para o usuário num enpoint.
*/
type ForecastOutput struct {
	Latitude      float64 `json:"latitude"`      // Latitude (localização geográfica)
	Longitude     float64 `json:"longitude"`     // Longitude (localização geográfica)
	Elevation     float64 `json:"elevation"`     // Nível do mar
	Temperature   float64 `json:"temperature"`   // Temperatura
	WindSpeed     float64 `json:"windspeed"`     // Velocidade do Vento
	WindDirection float64 `json:"winddirection"` // Direção do Vento
	DateTime      string  `json:"date_time"`     // Data e Hora
}
