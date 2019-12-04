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

func TestSolvePartTwo(t *testing.T) {
	suites := []struct {
		description string
		input1      int
		input2      int
		expected    int
	}{
		{"TODO", 112233, 112233, 1},
		{"TODO", 123444, 123444, 0},
		{"TODO", 122444, 122444, 1},
		{"TODO", 111122, 111122, 1},
		{"TODO", 111223, 111223, 1},
		{"TODO", 112223, 112223, 1},
		{"TODO", 122234, 122234, 0},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartTwo(suite.input1, suite.input2)

			if output != suite.expected {
				t.Errorf("TestSolvePartTwo(%d, %d) got %d but expect %d",
					suite.input1,
					suite.input2,
					output,
					suite.expected)
			}
		})
	}
}
