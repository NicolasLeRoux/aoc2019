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
	resultPartOne := solver.SolvePartOne(strs) // 200001 (7.062124ms)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

	start = time.Now()
	resultPartTwo := solver.SolvePartTwo(strs) // 379 (809.516µs)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
