package input_data

/*
Estrutura de dados para um conjunto específico de informações sobre o tempo na posição geográfica indicada advindo de uma fonte de dados json. Esses dados projetam uma previsão do tempo para 7 dias
*/
type Hourly struct {
	Time              []string  `json:"time"`               // Data e Hora em tempo real - lista
	Windspeed180m     []float64 `json:"windspeed_180m"`     // Velocidade do Vento em tempo real - lista
	Winddirection180m []float64 `json:"winddirection_180m"` // Direção do Vento em tempo real - lista
}
