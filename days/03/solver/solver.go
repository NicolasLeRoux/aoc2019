package solver

import (
	"math"
	"strconv"
	"strings"
)

// SolvePartOne is the main solver of the day 01 part 1 puzzle
func SolvePartOne(strs []string) int {
	pathWire1 := ParseWirePath(strs[0])
	pathWire2 := ParseWirePath(strs[1])
	intersections := []int{}

	for coord := range pathWire1 {
		_, ok := pathWire2[coord]

		if ok && coord != "0,0" {
			intersections = append(intersections, calculManhattanDistance(coord))
		}
	}

	return getMinNumber(intersections)
}

// ParseWirePath is a function to parse the wire's path
func ParseWirePath(str string) map[string]bool {
	paths := strings.Split(str, ",")
	x := 0
	y := 0
	m := map[string]bool{"0,0": true}

	for i := 0; i < len(paths); i++ {
		path := paths[i]
		dir := string(path[0])
		nb, err := strconv.Atoi(path[1:len(path)])

		if err != nil {
			panic("Cannot convert path to number !")
		}

		for j := 0; j < nb; j++ {
			if dir == "R" {
				x++
			} else if dir == "L" {
				x--
			} else if dir == "U" {
				y--
			} else if dir == "D" {
				y++
			} else {
				panic("Unknown direction to handle !")
			}

			key := strconv.Itoa(x) + "," + strconv.Itoa(y)
			m[key] = true
		}
	}

	return m
}

func calculManhattanDistance(str string) int {
	coords := strings.Split(str, ",")
	x, errX := strconv.Atoi(coords[0])
	y, errY := strconv.Atoi(coords[1])
	var result int

	if errX == nil && errY == nil {
		result = int(math.Abs(float64(x)) + math.Abs(float64(y)))
	}

	return result
}

func getMinNumber(nbs []int) int {
	min := math.MaxInt64

	for i := 0; i < len(nbs); i++ {
		if nbs[i] <= min {
			min = nbs[i]
		}
	}

	return min
}
