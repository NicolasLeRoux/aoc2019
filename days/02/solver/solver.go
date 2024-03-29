package solver

import (
    "strings"
    "strconv"
)

// SolvePartOne is the main solver of the day 02 part 1 puzzle
func SolvePartOne(str string) int {
    slice := ParseIntCode(str)

    slice[1] = 12
    slice[2] = 2

    processed := ProcessIntCode(slice)

    return processed[0]
}

// ProcessIntCode is a function that compute IntCode
func ProcessIntCode(ints []int) []int {

Loop:
    for i := 0; i < len(ints); i = i + 4 {
        opcode := ints[i]

        if opcode == 99 {
            break Loop
        }

        leftOperandIndex := ints[i + 1]
        leftOperand := ints[leftOperandIndex]
        rightOperandIndex := ints[i + 2]
        rightOperand := ints[rightOperandIndex]
        storeIndex := ints[i + 3]

        if opcode == 1 {
            result := leftOperand + rightOperand
            ints[storeIndex] = result
        } else if opcode == 2 {
            result := leftOperand * rightOperand
            ints[storeIndex] = result
        } else {
            panic("Something went wrong !")
        }
    }

    return ints
}

// ParseIntCode is a function that parse the input string
func ParseIntCode(str string) []int {
    var slice []int
    splitted := strings.Split(str, ",")

    for i := 0; i < len(splitted); i++ {
        nb, err := strconv.Atoi(splitted[i])

        if err == nil {
            slice = append(slice, nb)
        } else {
            panic("Cannot parse given number")
        }
    }

    return slice
}

// SolvePartTwo is the main solver of the day 02 part 2 puzzle
func SolvePartTwo(str string) int {
    slice := ParseIntCode(str)
    noun := 0
    verb := 0

Loop:
    for noun = 0; noun <= 99; noun++ {
        for verb = 0; verb <= 99; verb++ {
            sliceCopy := make([]int, len(slice))
            copy(sliceCopy, slice)

            sliceCopy[1] = noun
            sliceCopy[2] = verb

            processed := ProcessIntCode(sliceCopy)

            if processed[0] == 19690720 {
                break Loop
            }
        }
    }

    return 100 * noun + verb
}
