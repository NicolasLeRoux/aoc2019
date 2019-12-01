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
			result += calculFuelForGivenMass(nb)
		} else {
			panic(err)
		}
	}

	return result
}

// SolvePartTwo is the main solver of the day 01 part 2 puzzle
func SolvePartTwo(nbrs []string) int {
	var result int = 0

	for i := 0; i < len(nbrs); i++ {
		nb, err := strconv.Atoi(nbrs[i])

		if err == nil {
			fuel := calculFuelForGivenMass(nb)
			additionnal := fuel

		Loop:
			for {
				additionnal = calculFuelForGivenMass(additionnal)

				if additionnal <= 0 {
					break Loop
				}

				fuel += additionnal
			}

			result += fuel
		} else {
			panic(err)
		}
	}

	return result
}

func calculFuelForGivenMass(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}
