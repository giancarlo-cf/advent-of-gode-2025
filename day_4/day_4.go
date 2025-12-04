package main

import (
	"fmt"
	"math"

	"github.com/giancarlo-cf/advent-of-gode-2025/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2025, 4)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func process(input string) [][]rune {
	parser := parsers.MatrixRuneParser{BaseParser: parsers.BaseParser{Input: input}}
	return parser.GetStrings()
}

func isAccessible(matrix [][]rune, i int, j int) bool {

	maxJ := len(matrix[i]) - 1
	rolls := 0

	for jj := math.Max(0, float64(j-1)); jj <= math.Min(float64(j+1), float64(maxJ)); jj++ {
		if i > 0 && matrix[i-1][int(jj)] == '@' {
			rolls++
		}
		if i < len(matrix)-1 && matrix[i+1][int(jj)] == '@' {
			rolls++
		}
	}
	if j > 0 && matrix[i][j-1] == '@' {
		rolls++
	}
	if j < maxJ && matrix[i][j+1] == '@' {
		rolls++
	}

	return rolls < 4
}

func PartOne(input string) int {
	total := 0
	matrix := process(input)

	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] != '@' {
				continue
			}
			if isAccessible(matrix, i, j) {
				total++
			}
		}
	}

	return total
}

func PartTwo(input string) int {
	total := 0
	matrix := process(input)

	previousRollCount := -1
	currentRollCount := 0

	for previousRollCount != currentRollCount {
		previousRollCount = currentRollCount
		currentRollCount = 0
		toRemove := make([][]int, 0)

		for i := range len(matrix) {
			for j := range len(matrix[i]) {
				if matrix[i][j] != '@' {
					continue
				}
				currentRollCount++
				if isAccessible(matrix, i, j) {
					total++
					toRemove = append(toRemove, []int{i, j})
				}
			}
		}

		for _, position := range toRemove {
			matrix[position[0]][position[1]] = '.'
		}
	}

	return total
}
