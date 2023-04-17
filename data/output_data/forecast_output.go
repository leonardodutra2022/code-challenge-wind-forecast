package output_data

import "time"

/*
Estrutura de dados que representa um conjunto de dados de saída para o usuário num enpoint.
*/
type ForecastOutput struct {
	WindSpeed        float64   `json:"windspeed"`      // Velocidade do Vento
	WindDirection    float64   `json:"winddirection"`  // Direção do Vento
	DateTime         string    `json:"date_time"`      // Data e Hora
	DateLastQueryApi time.Time `json:"date_lastquery"` // data e hora da ultima consulta a API
}
