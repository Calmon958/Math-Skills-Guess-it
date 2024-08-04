package compute

import "testing"

func TestMedian(t *testing.T) {
	input := []float64{56, 47, 65}
	output := Median(input)
	expected := float64(56)

	if output != expected {
		t.Errorf("Median odd test failed:\n expected: %v\n got %v\n", expected, output)
	}

	input1 := []float64{70, 75, 122, 123}
	output1 := Median(input1)
	expected1 := float64(98.5)

	if output1 != expected1 {
		t.Errorf("Median even test failed:\n expected: %v\n got %v\n", expected1, output1)
	}
}

func TestMedianRev(t *testing.T) {
	input := []float64{56, 47, 65, 39}
	output := MedianRev(input)
	expected := float64(51.5)

	if output != expected {
		t.Errorf("Median test failed:\n expected: %v\n got %v\n", expected, output)
	}
}
