package solver

import (
	"reflect"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	suites := []struct {
		description string
		input       []string
		expected    int
	}{
		{"TODO", []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 6},
		{"TODO", []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
		{"TODO", []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
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

func TestParseWirePath(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected    map[string]int
	}{
		{"Should go to 'right'.", "R2", map[string]int{"0,0": 0, "1,0": 1, "2,0": 2}},
		{"Should go to 'left'.", "L1", map[string]int{"0,0": 0, "-1,0": 1}},
		{"Should go to 'up'.", "U1", map[string]int{"0,0": 0, "0,-1": 1}},
		{"Should go to 'down'.", "D2", map[string]int{"0,0": 0, "0,1": 1, "0,2": 2}},
		{"Should go to 'up' then 'right'.", "U1,R2", map[string]int{"0,0": 0, "0,-1": 1, "1,-1": 2, "2,-1": 3}},
		{"Should go to 'down' then 'left'.", "D1,L2", map[string]int{"0,0": 0, "0,1": 1, "-1,1": 2, "-2,1": 3}},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := ParseWirePath(suite.input[:])

			if !reflect.DeepEqual(output, suite.expected) {
				t.Errorf("ParseWirePath(%s) got %v but expect %v",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}

func TestCalculManhattanDistance(t *testing.T) {
	suites := []struct {
		description string
		input       string
		expected    int
	}{
		{"Should ", "3,3", 6},
		{"Should ", "-3,-3", 6},
		{"Should ", "15,5", 20},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := calculManhattanDistance(suite.input[:])

			if output != suite.expected {
				t.Errorf("ParseWirePath(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}

func TestGetMinNumber(t *testing.T) {
	suites := []struct {
		description string
		input       []int
		expected    int
	}{
		{"Should ", []int{3, 4, 5}, 3},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := getMinNumber(suite.input[:])

			if output != suite.expected {
				t.Errorf("ParseWirePath(%d) got %d but expect %d",
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
		{"TODO", []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 30},
		{"TODO", []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
		{"TODO", []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
	}

	for _, suite := range suites {
		t.Run(suite.description, func(t *testing.T) {
			output := SolvePartTwo(suite.input[:])

			if output != suite.expected {
				t.Errorf("TestSolvePartTwo(%s) got %d but expect %d",
					suite.input,
					output,
					suite.expected)
			}
		})
	}
}
