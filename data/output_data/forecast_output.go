package output_data

/*
Estrutura de dados que representa um conjunto de dados de saída para o usuário num enpoint.
*/
type ForecastOutput struct {
	WindSpeed        float64 `json:"velocidade_vento"`    // Velocidade do Vento
	WindDirection    float64 `json:"direcao_vento"`       // Direção do Vento
	DateTime         string  `json:"previsao_datahora"`   // Data e Hora
	DateLastQueryApi string  `json:"ultima_consulta_api"` // data e hora da ultima consulta a API
}
