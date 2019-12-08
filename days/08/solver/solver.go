package solver

import (
	"math"
)

// SolvePartOne is the main solver of the day 08 part 1 puzzle
func SolvePartOne(width int, height int, str string) int {
	layers := extractLayers(width, height, str)
	storedNbDigit0 := math.MaxInt32
	storedNbDigit1 := 0
	storedNbDigit2 := 0

	for i := 0; i < len(layers); i++ {
		nbDigit0, nbDigit1, nbDigit2 := extractDigitsFromLayer(layers[i])

		if storedNbDigit0 > nbDigit0 {
			storedNbDigit0 = nbDigit0
			storedNbDigit1 = nbDigit1
			storedNbDigit2 = nbDigit2
		}
	}

	return storedNbDigit1 * storedNbDigit2
}

func extractLayers(width int, height int, str string) []string {
	var layers []string
	startingIndex := 0
	increment := width * height

	for i := increment; i <= len(str); i = i + increment {
		layers = append(layers, str[startingIndex:i])
		startingIndex = i
	}

	return layers
}

func extractDigitsFromLayer(layer string) (int, int, int) {
	nbDigit0 := 0
	nbDigit1 := 0
	nbDigit2 := 0

	for i := 0; i < len(layer); i++ {
		switch string(layer[i]) {
		case "0":
			nbDigit0++
		case "1":
			nbDigit1++
		case "2":
			nbDigit2++
		}
	}

	return nbDigit0, nbDigit1, nbDigit2
}
