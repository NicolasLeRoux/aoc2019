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
	resultPartOne := solver.SolvePartOne(strs) // 3229 (158.584898ms)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)
}
