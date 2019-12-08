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
	resultPartOne := solver.SolvePartOne(25, 6, strs[0]) // 2016 (164.31µs)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

	start = time.Now()
	resultPartTwo := solver.SolvePartTwo(25, 6, strs[0]) // 100...100 (142.463µs)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Answer part two: %s (%+v)\n", resultPartTwo, elapsed)
	solver.FormatImg(25, 6, resultPartTwo)
}
