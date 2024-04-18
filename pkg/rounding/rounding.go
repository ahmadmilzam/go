package rounding

import "math"

func RoundFloat(v float64, prec uint) float64 {
	ratio := math.Pow(10, float64(prec))
	return math.Round(v*ratio) / ratio
}
