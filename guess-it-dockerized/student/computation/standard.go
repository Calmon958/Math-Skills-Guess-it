package compute

import "math"

//Func Standard looks for the standard deviation of a given dataset
func Standard(floatNum []float64) float64 {
	variance := Variance(floatNum)
	StandardDeviation := math.Sqrt(variance)
	return StandardDeviation
}
