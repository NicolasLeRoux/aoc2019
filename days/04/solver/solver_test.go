package solver

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input1      int
		input2      int
		expected    int
	}{
		{"TODO", 111111, 111111, 1},
		{"TODO", 111111, 111113, 3},
		{"TODO", 123789, 123789, 0},
		{"TODO", 223450, 223454, 0},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartOne(suite.input1, suite.input2)

			if output != suite.expected {
				t.Errorf("TestSolvePartOne(%d, %d) got %d but expect %d",
					suite.input1,
					suite.input2,
					output,
					suite.expected)
			}
		})
	}
}
