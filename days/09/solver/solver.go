package solver

import (
	"strconv"
	"strings"
)

// SolvePartOne is the main solver of the day 08 part 1 puzzle
func SolvePartOne(str string) int {
	return 0
}

func parseIntCode(str string) map[int]int {
	m := map[int]int{}
	splitted := strings.Split(str, ",")

	for i := 0; i < len(splitted); i++ {
		nb, err := strconv.Atoi(splitted[i])

		if err == nil {
			m[i] = nb
		} else {
			panic("Cannot parse given number")
		}
	}

	return m
}
