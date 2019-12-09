package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"./solver"
)

func main() {
	dat, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(dat), "\n")

	start := time.Now()
	resultPartOne := solver.SolvePartOne(strs[0], 1) // 15097178 (56.201µs)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

	start = time.Now()
	resultPartTwo := solver.SolvePartTwo(strs[0], 5) // 1558663 (75.592µs)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
