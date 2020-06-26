package utils

import "strconv"

func StringToInt(value string) (result int64) {

	result, _ = strconv.ParseInt(value, 10, 64)
	return result
}

func StringToFloat64(value string) (result float64) {

	result, _ = strconv.ParseFloat(value, 64)
	return result
}

func Float64ToString(value float64) (result string) {

	result = strconv.FormatFloat(value, 'f', 2, 64)
	return result
}
