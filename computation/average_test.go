package compute

import "testing"

func TestAverage(t *testing.T) {
	input := []float64{9, 56, 89, 90}
	output := Average(input)
	expected := float64(61)

	if output != expected {
		t.Errorf("Average test failed:\n Got: %v\n Expected: %v\n", output, expected)
	}
}
