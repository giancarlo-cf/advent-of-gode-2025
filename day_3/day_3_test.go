package main

import (
	"os"
	"strings"
	"testing"

	"github.com/giancarlo-cf/advent-of-gode-2025/utils"
)

func setup() string {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "day_3") {
		os.Chdir("..")
	}
	input := utils.FetchTestInput(3)

	return input
}

func TestPartOne(t *testing.T) {
	input := setup()
	total := PartOne(input)
	expected := 357

	if total != expected {
		t.Errorf("PartOne() = %d; want %d", total, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := setup()
	total := PartTwo(input)
	expected := 3121910778619

	if total != expected {
		t.Errorf("PartTwo() = %d; want %d", total, expected)
	}
}
