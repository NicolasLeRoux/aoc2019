package solver

import (
	"testing"
	"reflect"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected    int
	}{
		{"TODO", "1,0,0,0,99,0,0,0,0,0,0,0,10", 12},
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

func TestProcessIntCode(t *testing.T) {
	suites := []struct {
		description string
		input       []int
		expected    []int
	}{
		{"TODO", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"TODO", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"TODO", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"TODO", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := ProcessIntCode(suite.input[:])

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestProcessIntCode(%d) got %d but expect %d",
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
		expected    []int
	}{
		{"TODO", "1,0,0,0,99", []int{1, 0, 0, 0, 99}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := ParseIntCode(suite.input[:])

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestProcessIntCode(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}
