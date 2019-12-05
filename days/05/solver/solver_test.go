package solver

import (
	"reflect"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input1      string
		input2      int
		expected    []int
	}{
		{"Simply output the input value.", "3,0,4,0,99", 10, []int{10}},
		{"Simply output the input value bis.", "3,0,4,0,99", 33, []int{33}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartOne(suite.input1, suite.input2)

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestSolvePartOne(%s, %d) got %d but expect %d",
					suite.input1,
					suite.input2,
					output,
					suite.expected)
			}
		})
	}
}

func TestParseOpcode(t *testing.T) {
	suites := []struct {
		description string
		input       int
		expected1   int
		expected2   int
		expected3   int
		expected4   int
	}{
		{"TODO", 1002, 2, 0, 1, 0},
		{"TODO", 99, 99, 0, 0, 0},
		{"TODO", 1101, 1, 1, 1, 0},
		{"TODO", 10004, 4, 0, 0, 1},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output1, output2, output3, output4 := parseOpcode(suite.input)

			if output1 != suite.expected1 || output2 != suite.expected2 || output3 != suite.expected3 || output4 != suite.expected4 {
				t.Errorf("TestParseOpcode(%d) got %d, %d, %d, %d but expect %d, %d, %d, %d",
					suite.input,
					output1,
					output2,
					output3,
					output4,
					suite.expected1,
					suite.expected2,
					suite.expected3,
					suite.expected4)
			}
		})
	}
}

func TestGetIncrementForInstruction(t *testing.T) {
	suites := []struct {
		description string
		input       int
		expected    int
	}{
		{"Should increment by 4 given the instruction 1.", 1, 4},
		{"Should increment by 4 given the instruction 2.", 2, 4},
		{"Should increment by 2 given the instruction 3.", 3, 2},
		{"Should increment by 2 given the instruction 4.", 4, 2},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := getIncrementForInstruction(suite.input)

			if output != suite.expected {
				t.Errorf("TestGetIncrementForInstruction(%d) got %d but expect %d",
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
