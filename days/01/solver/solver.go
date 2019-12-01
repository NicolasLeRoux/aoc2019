package solver

import (
	"math"
	"strconv"
)

// SolvePartOne is the main solver of the day 01 part 1 puzzle
func SolvePartOne(nbrs []string) int {
	var result int = 0

	for i := 0; i < len(nbrs); i++ {
		nb, err := strconv.Atoi(nbrs[i])

		if err == nil {
			result += int(math.Floor(float64(nb/3))) - 2
		} else {
			panic(err)
		}
	}

	return result
}
