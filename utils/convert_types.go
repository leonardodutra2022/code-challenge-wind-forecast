package utils

import (
	"encoding/json"
	"strconv"
)

func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}

func BoolToString(value bool) string {
	return strconv.FormatBool(value)
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func ByteToJson(bytes []byte, obj interface{}) error {
	return json.Unmarshal(bytes, &obj)
}
