package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	compute "guess/computation"
)

func main() {
	var numbers []float64
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter numbers one by one (Ctrl+D to end input):")

	for scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer")
			continue
		}

		numbers = append(numbers, float64(number))
		if len(numbers) > 1 {
			mean := compute.Average(numbers)
			stddev := compute.Standard(numbers)
			rangeLower, rangeUpper := predictRange(mean, stddev)
			fmt.Printf("%d %d\n", rangeLower, rangeUpper)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func predictRange(mean, stddev float64) (int, int) {
	deviation := 1.2 * stddev // Using 1.5 standard deviations as range
	return int(mean - deviation), int(mean + deviation)
}
