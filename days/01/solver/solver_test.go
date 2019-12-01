package solver

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input       []string
		expected    int
	}{
		{"For a mass of 12", []string{"12"}, 2},
		{"For a mass of 14", []string{"14"}, 2},
		{"For a mass of 1969", []string{"1969"}, 654},
		{"For a mass of 100756", []string{"100756"}, 33583},
		{"For a mass of 12, 14, 1969 and 100756", []string{"12", "14", "1969", "100756"}, 34241},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartOne(suite.input[:])

			if output != suite.expected {
				t.Errorf("TestSolvePartOne(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}
