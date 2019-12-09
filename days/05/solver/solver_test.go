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
		{"Should increment by 3 given the instruction 5.", 5, 3},
		{"Should increment by 3 given the instruction 6.", 6, 3},
		{"Should increment by 4 given the instruction 7.", 7, 4},
		{"Should increment by 4 given the instruction 8.", 8, 4},
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

func TestSolvePartTwo(t *testing.T) {
	suites := []struct {
		description string
		input1      string
		input2      int
		expected    []int
	}{
		{"Simply output the input value.", "3,0,4,0,99", 10, []int{10}},
		{"Simply output the input value bis.", "3,0,4,0,99", 33, []int{33}},
		{"Should return 1 given an input equal to 8.", "3,9,8,9,10,9,4,9,99,-1,8", 8, []int{1}},
		{"Should return 0 given an input not equal to 8.", "3,9,8,9,10,9,4,9,99,-1,8", 9, []int{0}},
		{"Should return 1 given an input less than 8.", "3,9,7,9,10,9,4,9,99,-1,8", 7, []int{1}},
		{"Should return 0 given an input not less than 8.", "3,9,7,9,10,9,4,9,99,-1,8", 8, []int{0}},
		{"Should return 1 given an input equal to 8 (bis).", "3,3,1108,-1,8,3,4,3,99", 8, []int{1}},
		{"Should return 0 given an input not equal to 8 (bis).", "3,3,1108,-1,8,3,4,3,99", 9, []int{0}},
		{"Should return 1 given an input less than 8 (bis).", "3,3,1107,-1,8,3,4,3,99", 7, []int{1}},
		{"Should return 0 given an input not less than 8 (bis).", "3,3,1107,-1,8,3,4,3,99", 8, []int{0}},
		{"Should return 0 given an input equal to 0.", "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, []int{0}},
		{"Should return 1 given an input not equal to 0.", "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 2, []int{1}},
		{"Should return 0 given an input equal to 0 (bis).", "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, []int{0}},
		{"Should return 1 given an input not equal to 0 (bis).", "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 2, []int{1}},
		{"Should return 999 given an input below 8.", "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 7, []int{999}},
		{"Should return 1000 given an input equal to 8.", "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 8, []int{1000}},
		{"Should return 1001 given an input greater than 8.", "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 9, []int{1001}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartTwo(suite.input1, suite.input2)

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestSolvePartTwo(%s, %d) got %d but expect %d",
					suite.input1,
					suite.input2,
					output,
					suite.expected)
			}
		})
	}
}
