package compute

import (
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{9, 56, 89, 90}
	output := Variance(input)
	expected := float64(1088.5)

	if output != expected {
		t.Errorf("Variance test failed: Got: %v\n Expected: %v\n", output, expected)
	}
}
