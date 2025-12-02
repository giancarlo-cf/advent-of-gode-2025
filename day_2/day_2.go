package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/giancarlo-cf/advent-of-gode-2025/utils"
	"github.com/giancarlo-cf/advent-of-gode-utils/parsers"
)

func main() {
	input := utils.FetchInput(2025, 2)

	fmt.Printf("Part One: %d\n", PartOne(input))
	fmt.Printf("Part Two: %d\n", PartTwo(input))
}

func process(input string) [][]int {
	parser := parsers.RangeIntParser{BaseParser: parsers.BaseParser{Input: input}}
	return parser.GetInts(",")
}

func isInvalid(id int) bool {
	idString := strconv.Itoa(id)
	idStringLength := len(idString)
	idStringHalfLength := idStringLength / 2
	idFirstHalf, idSecondHalf := idString[:idStringHalfLength], idString[idStringHalfLength:]

	return idFirstHalf == idSecondHalf
}

func isReallyInvalid(id int) bool {
	idString := strconv.Itoa(id)
	idStringLength := len(idString)

	for i := 1; i <= idStringLength; i++ {
		re := regexp.MustCompile(idString[:i])
		occurrences := len(re.FindAllString(idString, -1))
		if occurrences*i == idStringLength && occurrences > 1 {
			return true
		}
	}

	return false
}

func PartOne(input string) int {
	total := 0
	ranges := process(input)

	for _, pair := range ranges {
		l, r := pair[0], pair[1]
		for n := l; n <= r; n++ {
			if isInvalid(n) {
				total += n
			}
		}
	}

	return total
}

func PartTwo(input string) int {
	total := 0
	ranges := process(input)

	for _, pair := range ranges {
		l, r := pair[0], pair[1]
		fmt.Printf("Checking: %d - %d\n", l, r)
		for n := l; n <= r; n++ {
			if isReallyInvalid(n) {
				total += n
			}
		}
	}

	return total
}
