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
		{"TODO", []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, 42},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartOne(suite.input)

			if output != suite.expected {
				t.Errorf("TestSolvePartOne(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	suites := []struct {
		description string
		input       []string
		expected    int
	}{
		{"TODO", []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}, 4},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartTwo(suite.input)

			if output != suite.expected {
				t.Errorf("TestSolvePartTwo(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}
