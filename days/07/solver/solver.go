package solver

import (
	"strconv"
	"strings"
)

// SolvePartOne is the main solver of the day 06 part 1 puzzle
func SolvePartOne(str string) (int, []int) {
	settings := generateEveryPhaseSetting([]int{0, 1, 2, 3, 4})
	maxThrusterSignal := 0
	var savedSetting []int

	for i := 0; i < len(settings); i++ {
		lastOutput := 0
		setting := settings[i]

		for j := 0; j < len(setting); j++ {
			outputs := processIntcode(str, []int{setting[j], lastOutput})

			lastOutput = outputs[len(outputs)-1]
		}

		if maxThrusterSignal < lastOutput {
			maxThrusterSignal = lastOutput
			savedSetting = setting

			//fmt.Printf("%d -> %v\n", maxThrusterSignal, setting)
		}
	}

	return maxThrusterSignal, savedSetting
}

func processIntcode(str string, inputs []int) []int {
	intcodes := parseIntCode(str)
	var outputs []int
	inputIndex := 0

Loop:
	for i := 0; i < len(intcodes); {
		instruction, modeParamNb1, modeParamNb2, modeParamNb3 := parseOpcode(intcodes[i])

		if instruction == 1 || instruction == 2 {
			var leftOperand int

			if modeParamNb1 == 0 {
				leftOperandIndex := intcodes[i+1]
				leftOperand = intcodes[leftOperandIndex]
			} else {
				leftOperand = intcodes[i+1]
			}

			var rightOperand int

			if modeParamNb2 == 0 {
				rightOperandIndex := intcodes[i+2]
				rightOperand = intcodes[rightOperandIndex]
			} else {
				rightOperand = intcodes[i+2]
			}

			if modeParamNb3 == 0 {
				storeIndex := intcodes[i+3]

				if instruction == 1 {
					intcodes[storeIndex] = leftOperand + rightOperand
				} else {
					intcodes[storeIndex] = leftOperand * rightOperand
				}
			} else {
				if instruction == 1 {
					intcodes[i+3] = leftOperand + rightOperand
				} else {
					intcodes[i+3] = leftOperand * rightOperand
				}
			}
		} else if instruction == 3 {
			if modeParamNb1 == 0 {
				parameterIndex := intcodes[i+1]
				intcodes[parameterIndex] = inputs[inputIndex]
			} else {
				intcodes[i+1] = inputs[inputIndex]
			}
			inputIndex++
		} else if instruction == 4 {
			if modeParamNb1 == 0 {
				parameterIndex := intcodes[i+1]
				outputs = append(outputs, intcodes[parameterIndex])
			} else {
				outputs = append(outputs, intcodes[i+1])
			}
		} else if instruction == 5 || instruction == 6 {
			var leftOperand int

			if modeParamNb1 == 0 {
				leftOperandIndex := intcodes[i+1]
				leftOperand = intcodes[leftOperandIndex]
			} else {
				leftOperand = intcodes[i+1]
			}

			var rightOperand int

			if modeParamNb2 == 0 {
				rightOperandIndex := intcodes[i+2]
				rightOperand = intcodes[rightOperandIndex]
			} else {
				rightOperand = intcodes[i+2]
			}

			if leftOperand != 0 && instruction == 5 || leftOperand == 0 && instruction == 6 {
				i = rightOperand
				continue
			}
		} else if instruction == 7 || instruction == 8 {
			var param1 int

			if modeParamNb1 == 0 {
				param1Index := intcodes[i+1]
				param1 = intcodes[param1Index]
			} else {
				param1 = intcodes[i+1]
			}

			var param2 int

			if modeParamNb2 == 0 {
				param2Index := intcodes[i+2]
				param2 = intcodes[param2Index]
			} else {
				param2 = intcodes[i+2]
			}

			var param3 int

			if modeParamNb3 == 0 {
				//param3Index := intcodes[i+3]
				//param3 = intcodes[param3Index]
				param3 = intcodes[i+3]
			} else {
				param3 = intcodes[i+3]
			}

			if param1 < param2 && instruction == 7 || param1 == param2 && instruction == 8 {
				intcodes[param3] = 1
			} else {
				intcodes[param3] = 0
			}
		} else if instruction == 99 {
			break Loop
		} else {
			panic("Unknow instruction!")
		}

		i = i + getIncrementForInstruction(instruction)
	}

	return outputs
}

func parseIntCode(str string) []int {
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

func parseOpcode(opcode int) (int, int, int, int) {
	opcodeStr := strconv.Itoa(opcode)
	length := len(opcodeStr)

	var instruction int
	var instructionErr error

	if length == 1 {
		instruction = opcode
	} else {
		instruction, instructionErr = strconv.Atoi(opcodeStr[length-2:])
	}

	if instructionErr != nil {
		panic("Cannot parse opcode instruction!")
	}

	modeParamNb1 := 0
	var modeParamNb1Err error

	if length >= 3 {
		modeParamNb1, modeParamNb1Err = strconv.Atoi(opcodeStr[length-3 : length-2])
	}

	if modeParamNb1Err != nil {
		panic("Cannot parse mode of 1st parameter!")
	}

	modeParamNb2 := 0
	var modeParamNb2Err error

	if length >= 4 {
		modeParamNb2, modeParamNb2Err = strconv.Atoi(opcodeStr[length-4 : length-3])
	}

	if modeParamNb2Err != nil {
		panic("Cannot parse mode of 2nd parameter!")
	}

	modeParamNb3 := 0
	var modeParamNb3Err error

	if length >= 5 {
		modeParamNb3, modeParamNb3Err = strconv.Atoi(opcodeStr[length-5 : length-4])
	}

	if modeParamNb3Err != nil {
		panic("Cannot parse mode of 3rd parameter!")
	}

	return instruction, modeParamNb1, modeParamNb2, modeParamNb3
}

func getIncrementForInstruction(instruction int) int {
	increment := 0

	if instruction == 1 || instruction == 2 || instruction == 7 || instruction == 8 {
		increment = 4
	} else if instruction == 3 || instruction == 4 {
		increment = 2
	} else if instruction == 5 || instruction == 6 {
		increment = 3
	} else if instruction == 99 {
		increment = 1
	} else {
		panic("Unknow instruction!")
	}

	return increment
}

func generateEveryPhaseSetting(nbs []int) [][]int {
	if len(nbs) == 2 {
		return [][]int{[]int{nbs[0], nbs[1]}, []int{nbs[1], nbs[0]}}
	}

	var result [][]int

	for i := 0; i < len(nbs); i++ {
		slice := make([]int, len(nbs))
		copy(slice, nbs)
		slice = append(slice[:i], slice[i+1:]...)

		lowerSetting := generateEveryPhaseSetting(slice)

		for j := 0; j < len(lowerSetting); j++ {
			result = append(result, append([]int{nbs[i]}, lowerSetting[j]...))
		}
	}

	return result
}
