package input_data

type ForecastInput struct {
	Latitude  float64 `json:"latitude"`   // Latitude (localização geográfica)
	Longitude float64 `json:"longitude"`  // Longitude (localização geográfica)
	Elevation float64 `json:"elevantion"` // Nível do mar
}
