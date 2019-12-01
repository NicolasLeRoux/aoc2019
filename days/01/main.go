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

	modules := strings.Split(string(dat), "\n")

	start := time.Now()
	resultPartOne := solver.SolvePartOne(modules) // 3267638 (3.508µs)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

	start = time.Now()
	resultPartTwo := solver.SolvePartTwo(modules) // 4898585 (8.564µs)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Answer part two: %d (%+v)\n", resultPartTwo, elapsed)
}
