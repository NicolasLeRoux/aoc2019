package solver

import (
	"strconv"
)

// SolvePartOne is the main solver of the day 04 part 1 puzzle
func SolvePartOne(min int, max int) int {
	var list []int

	for i := min; i <= max; i++ {
		nb := strconv.Itoa(i)
		digit1, err1 := strconv.Atoi(string(nb[0]))
		digit2, err2 := strconv.Atoi(string(nb[1]))
		digit3, err3 := strconv.Atoi(string(nb[2]))
		digit4, err4 := strconv.Atoi(string(nb[3]))
		digit5, err5 := strconv.Atoi(string(nb[4]))
		digit6, err6 := strconv.Atoi(string(nb[5]))

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
			panic("Cannot parse number !")
		}

		if digit1 <= digit2 && digit2 <= digit3 && digit3 <= digit4 && digit4 <= digit5 && digit5 <= digit6 && (digit1 == digit2 || digit2 == digit3 || digit3 == digit4 || digit4 == digit5 || digit5 == digit6) {
			list = append(list, i)
		}
	}

	return len(list)
}
