package main

import (
	"fmt"
	"strconv"

	"github.com/giancarlo-cf/advent-of-gode-2025/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2025, 1)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func process(input string) []int {
	parser := parsers.BaseParser{Input: input}
	lines := parser.GetLines()

	numbers := make([]int, len(lines))

	for index, line := range lines {
		direction := line[0]

		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			number *= -1
		}

		numbers[index] = number
	}

	return numbers
}

func PartOne(input string) int {
	total := 0

	numbers := process(input)
	current := 50

	for _, number := range numbers {
		current += number

		for current > 99 || current < 0 {
			switch {
			case current < 0:
				current += 100
			case current > 99:
				current -= 100
			}
		}

		if current == 0 {
			total++
		}
	}

	return total
}

func PartTwo(input string) int {
	total := 0

	numbers := process(input)
	current := 50


	for _, number := range numbers {
		ticks := 0

		if current == 0 && number < 0 {
			ticks--
		}

		current += number

		if current == 0 {
			total++
		}

		if !(current > 99 || current < 0) {
			continue
		}
		
		for current > 99 || current < 0 {
			
			switch {
			case current < 0:
				current += 100
			case current > 99:
				current -= 100
			}
			
			ticks++
		}

		if current == 0 && number < 0 {
			ticks++
		}

		total += ticks
	}

	return total
}
