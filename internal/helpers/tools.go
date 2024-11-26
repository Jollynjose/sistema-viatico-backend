package helpers

import "strconv"

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
