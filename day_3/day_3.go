package main

import (
	"fmt"
	"math"

	"github.com/giancarlo-cf/advent-of-gode-2025/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2025, 3)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func process(input string) [][]int {
	parser := parsers.RowIntParser{BaseParser: parsers.BaseParser{Input: input}}
	return parser.GetInts("")
}

func PartOne(input string) int {
	total := 0
	matrix := process(input)

	for _, row := range matrix {
		maxNumber := -1
		currentMax := -1

		for i := len(row) - 1; i >= 0; i-- {

			if currentMax != -1 {
				maxNumber = int(math.Max(float64(row[i]*10+currentMax), float64(maxNumber)))
			}

			currentMax = int(math.Max(float64(row[i]), float64(currentMax)))
		}

		total += maxNumber
	}

	return total
}

func PartTwo(input string) int {
	total := 0
	matrix := process(input)

	for _, row := range matrix {
		maxArray := make([]int, 12)
		copy(maxArray, row[:12])
		candidateIndex := 0
		for i := 12; i < len(row); i++ {

			for k := range 12 {
				candidateIndex = k
				if k < 11 && maxArray[k+1] > maxArray[k] {
					break
				}
			}

			new := row[i]

			if candidateIndex == 11 && new > maxArray[candidateIndex] {
				maxArray[candidateIndex] = new
			} else if candidateIndex < 11 {
				maxArray = append(maxArray[:candidateIndex], maxArray[candidateIndex+1:]...)
				maxArray = append(maxArray, new)
			}
		}
		for i := 12; i > 0; i-- {
			total += maxArray[12-i] * int(math.Pow(10, float64(i-1)))
		}
	}

	return total
}
