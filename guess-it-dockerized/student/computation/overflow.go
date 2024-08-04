package compute

import (
	"fmt"
	"math"
)

//Func Overflow checks if the values obtain are within the range of int64 digits
func Overflow(digit float64) int64 {
	var result int64
	if digit > float64(math.MaxInt64) || digit < float64(math.MinInt64) {
		fmt.Println("Value is out of range")
	} else {
		result = int64(math.Round(digit))
	}
	return result
}
