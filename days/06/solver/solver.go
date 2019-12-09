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

// SolvePartTwo is the main solver of the day 06 part 2 puzzle
func SolvePartTwo(strs []string) int {
	m := map[string]string{}

	for i := 0; i < len(strs); i++ {
		data := strings.Split(strs[i], ")")

		m[data[1]] = data[0]
	}

	val := m["YOU"]
	var list1 []string
Loop1:
	for {
		if val == "COM" {
			break Loop1
		}

		list1 = append(list1, val)
		val = m[val]
	}

	val = m["SAN"]
	var list2 []string
Loop2:
	for {
		if val == "COM" {
			break Loop2
		}

		list2 = append(list2, val)
		val = m[val]
	}

	i := 0
	j := 0
Loop3:
	for i = 0; i < len(list1); i++ {
		for j = 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				break Loop3
			}
		}
	}

	return i + j
}
