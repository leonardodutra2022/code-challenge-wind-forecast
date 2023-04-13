package input_data

/*
Estrutura de dados de entrada de dados para um conjunto advindo de uma fonte de dados json.
*/
type ForecastInput struct {
	Latitude       float64        `json:"latitude"`        // Latitude (localização geográfica)
	Longitude      float64        `json:"longitude"`       // Longitude (localização geográfica)
	Elevation      float64        `json:"elevation"`       // Nível do mar
	CurrentWeather CurrentWeather `json:"current_weather"` // Informações do tempo em geral
}
