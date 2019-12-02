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
	resultPartOne := solver.SolvePartOne(strs[0]) // 3654868 (22.908µs)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)

	start = time.Now()
	resultPartTwo := solver.SolvePartTwo(strs[0]) // 7014 (5.406472ms)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartTwo, elapsed)
}
