package solver

import (
	"reflect"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input1      int
		input2      int
		input3      string
		expected    int
	}{
		{"Should return 4 as value.", 2, 2, "001100221122", 4},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartOne(suite.input1, suite.input2, suite.input3)

			if output != suite.expected {
				t.Errorf("TestSolvePartOne(%d, %d, %s) got %d but expect %d",
					suite.input1,
					suite.input2,
					suite.input3,
					output,
					suite.expected)
			}
		})
	}
}

func TestExtractLayers(t *testing.T) {
	suites := []struct {
		description string
		input1      int
		input2      int
		input3      string
		expected    []string
	}{
		{"Should split the data image in 2 layers.", 3, 2, "123456789012", []string{"123456", "789012"}},
		{"Should split the data image in 3 layers.", 2, 2, "123456789012", []string{"1234", "5678", "9012"}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := extractLayers(suite.input1, suite.input2, suite.input3)

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("TestExtractLayers(%d, %d, %s) got %v but expect %v",
					suite.input1,
					suite.input2,
					suite.input3,
					output,
					suite.expected)
			}
		})
	}
}

func TestExtractDigitsFromLayer(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected1   int
		expected2   int
		expected3   int
	}{
		{"Should return 1 for each output.", "012", 1, 1, 1},
		{"Should return 3 zero, 2 one and one 2.", "010210", 3, 2, 1},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output1, output2, output3 := extractDigitsFromLayer(suite.input)

			if output1 != suite.expected1 || output2 != suite.expected2 || output3 != suite.expected3 {
				t.Errorf("TestExtractDigitsFromLayer(%s) got %d, %d, %d but expect %d, %d, %d",
					suite.input,
					output1,
					output2,
					output3,
					suite.expected1,
					suite.expected2,
					suite.expected3)
			}
		})
	}
}
