package compute

import "testing"

func TestStandard(t *testing.T) {
	input := []float64{9, 56, 89, 90}
	output := Standard(input)
	expected := float64(32.992423372647245)

	if output != expected {
		t.Errorf("Standard deviation test failed:\n Got: %v\n Expected: %v\n", output, expected)
	}
}
