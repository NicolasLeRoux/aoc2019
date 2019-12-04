package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	boundary := strings.Split(strs[0], "-")

	min, errMin := strconv.Atoi(boundary[0])
	max, errMax := strconv.Atoi(boundary[1])

	if errMin != nil || errMax != nil {
		panic("Cannot parse boundary !")
	}

	start := time.Now()
	resultPartOne := solver.SolvePartOne(min, max) // 481 (67.635753ms)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Answer part one: %d (%+v)\n", resultPartOne, elapsed)
}
