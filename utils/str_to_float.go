package utils

import "strconv"

func Float64ToString(number float64) string {
	return strconv.FormatFloat(number, 'f', 2, 64)
}
