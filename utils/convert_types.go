package utils

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/caarlos0/env"
	"github.com/leonardodutra2022/code-challenge-wind-forecast/data/config"
)

/*
Função utilitária para converter um valor float para string
*/
func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

/*
Função utilitária para converter um valor bool para string
*/
func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

/*
Função utilitária para converter um valor int para string
*/
func IntToString(number int) string {
	return strconv.Itoa(number)
}

/*
Função utilitária para converter um array de bytes para json, quando o input é um json
*/
func ByteToJson(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, &obj)
}

/*
Função utilitária para converter um valor de data como string para tipo time.Time
*/
func DateStringToTime(dateString string) time.Time {
	cfg := config.Config{}
	env.Parse(&cfg)
	timeLocal, _ := time.LoadLocation(cfg.TimeZone)
	date, err := time.ParseInLocation("2006-01-02T15:04", dateString, timeLocal)
	if err != nil {
		return time.Time{}
	}
	return date
}
