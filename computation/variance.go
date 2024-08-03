package compute

import (
	"math"
)

//Func Variance performs variance operation
func Variance(floatNum []float64) float64 {
	average := Average(floatNum)
	// fmt.Println(average)
	sum := 0.0
	for _, digit := range floatNum {
		digit = math.Pow((digit - average), 2)
		sum += digit
	}
	variance := sum / float64(len(floatNum))
	return variance

}
