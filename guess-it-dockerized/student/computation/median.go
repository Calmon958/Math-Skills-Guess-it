package compute

import (
	"fmt"
	"os"
	"sort"
)

// looks for the number at the middle of a given set of numbers and returns it.
func Median(content []float64) float64 {
	dataFloat := sort.Float64Slice(content)
	dataFloat.Sort()
	result := 0.0
	if len(dataFloat) == 0 {
		fmt.Println("Empty file")
		os.Exit(1)
	}

	//number at the middle of the provided values
	middle := len(dataFloat) / 2

	if len(dataFloat)%2 == 0 {
		result = (dataFloat[middle] + dataFloat[middle-1]) / 2
	} else {
		result = dataFloat[middle]
	}
	return result
}

//MedianRev prints the revers sorting of the median value
func MedianRev(content []float64) float64 {
	dataFloat := sort.Float64Slice(content)
	sort.Sort(sort.Reverse(dataFloat))
	result := 0.0

	middle := len(dataFloat) / 2

	if len(dataFloat)%2 == 0 {
		result = (dataFloat[middle] + dataFloat[middle-1]) / 2
	} else {
		result = dataFloat[middle]
	}
	return result
}
