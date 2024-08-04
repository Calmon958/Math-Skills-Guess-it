package compute

//Average takes the sum of all the numbers and divides it with the length of the no of values
func Average(floatNum []float64) float64 {

	var sum float64
	sum = 0.0
	for _, data := range floatNum {
		sum = sum + data
	}
	a := sum / float64(len(floatNum))

	return a
}
