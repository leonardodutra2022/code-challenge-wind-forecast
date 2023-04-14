package utils_test

import (
	"testing"
	"time"

	"github.com/leonardodutra2022/code-challenge-wind-forecast/utils"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func TestFloat64ToString(t *testing.T) {
	numberFloat64 := 100.50
	assert.Equal(t, "100.50", utils.Float64ToString(numberFloat64), "o número em formato float informado deve ser convertido ao equivalente em string")
}

func TestBoolToString(t *testing.T) {
	boolValue := true
	assert.Equal(t, "true", utils.BoolToString(boolValue), "o valor booleano informado deve ser convertido ao equivalente em string")
}

func TestIntToString(t *testing.T) {
	numberInt := 100
	assert.Equal(t, "100", utils.IntToString(numberInt), "o número inteiro informado deve ser convertido ao equivalente em string")
}

func TestByteToJson(t *testing.T) {
	jsonBytes := []byte(
		`{
			"id":1000,
			"message":"test"
		}`,
	)
	var jsonResult TestStruct
	err := utils.ByteToJson(jsonBytes, &jsonResult)
	assert.Nil(t, err, "a conversão de objeto json em bytes para json deve ocorrer sem erros")
}

func TestByteToJsonValues(t *testing.T) {
	jsonBytes := []byte(
		`{
			"id":1000,
			"message":"test"
		}`,
	)
	var jsonResult TestStruct
	utils.ByteToJson(jsonBytes, &jsonResult)
	assert.Equal(t, 1000, jsonResult.ID, "os valores da conversão entre objeto json em array de bytes para json deve bater/ser equivalente (bind)")
	assert.Equal(t, "test", jsonResult.Message, "os valores da conversão entre objeto json em array de bytes para json deve bater/ser equivalente (bind)")
}

func TestDateStringToTime(t *testing.T) {
	dateStr := "2023-04-14T18:00"
	timeLocal, _ := time.LoadLocation("America/Sao_Paulo")
	dateTest := time.Date(2023, time.April, 14, 18, 0, 0, 0, timeLocal)
	assert.Equal(t, dateTest, utils.DateStringToTime(dateStr), "a data convertida a partir de uma string para time.Time deve ser equivalente")
}
