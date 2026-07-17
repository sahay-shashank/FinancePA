package utils

import "math"

func RoundToTwoDecimal(number float64) float64 {
	return math.Round(number*100) / 100
}
