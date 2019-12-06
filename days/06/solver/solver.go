package solver

import (
	"strings"
)

// SolvePartOne is the main solver of the day 06 part 1 puzzle
func SolvePartOne(strs []string) int {
	m := map[string]string{}

	for i := 0; i < len(strs); i++ {
		data := strings.Split(strs[i], ")")

		m[data[1]] = data[0]
	}

	cpt := 0

	for _, val := range m {
	Loop:
		for {
			if val == "COM" {
				break Loop
			}

			val = m[val]

			cpt++
		}

		cpt++
	}

	return cpt
}
