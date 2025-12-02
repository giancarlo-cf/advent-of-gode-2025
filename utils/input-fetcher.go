package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/giancarlo-cf/advent-of-gode-utils"
	"github.com/joho/godotenv"
)

func FetchInput(year int, day int) string {

	inputFile := fmt.Sprintf("inputs/%d.txt", day)
	if data, err := os.ReadFile(inputFile); err == nil {
		return string(data)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	sessionCookie := os.Getenv("SESSION_COOKIE")

	inputFetcher := advent_of_gode_utils.InputFetcher{}
	input, err := inputFetcher.FetchInput(year, day, sessionCookie)
	if err != nil {
		log.Fatalf("Error fetching input: %v", err)
	}

	if err := os.MkdirAll("inputs", 0755); err != nil {
		log.Fatalf("Error creating inputs directory: %v", err)
	}
	err = os.WriteFile(fmt.Sprintf("inputs/%d.txt", day), []byte(input), 0644)
	if err != nil {
		log.Fatalf("Error saving input file: %v", err)
	}

	return input
}

func FetchTestInput(day int) string {
	inputFile := fmt.Sprintf("inputs/%d_test.txt", day)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading test input file: %v", err)
	}
	return string(data)
}

func FetchSpecialTestInput(day int) string {
	inputFile := fmt.Sprintf("inputs/%d_test_2.txt", day)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading test input file: %v", err)
	}
	return string(data)
}
