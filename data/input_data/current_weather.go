package input_data

/*
Estrutura de dados para um conjunto específico de informações sobre o tempo na posição geográfica indicada advindo de uma fonte de dados json.
*/
type CurrentWeather struct {
	Temperature   float64 `json:"temperature"`   // Temperatura
	WindSpeed     float64 `json:"windspeed"`     // Velocidade do Vento
	WindDirection float64 `json:"winddirection"` // Direção do Vento
	DateTime      string  `json:"time"`          // Data e Hora
}
