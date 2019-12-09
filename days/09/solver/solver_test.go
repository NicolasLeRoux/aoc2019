package solver

import (
	"reflect"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected    int
	}{
		{"TODO", "104,1125899906842624,99", 1125899906842624},
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

func TestParseIntCode(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected    map[int]int
	}{
		{"TODO", "1,0,0,0,99", map[int]int{0: 1, 1: 0, 2: 0, 3: 0, 4: 99}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := parseIntCode(suite.input[:])

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestProcessIntCode(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}
