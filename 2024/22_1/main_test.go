package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func parseInput(input string) []uint64 {
	lines := aoc.SplitAndTrim(input, "\n")
	numbers := make([]uint64, len(lines))

	for i, line := range lines {
		numbers[i], _ = strconv.ParseUint(line, 10, 64)
	}

	return numbers
}

func solution(input string) uint64 {
	initialSecretNumbers := parseInput(input)

	var sum uint64 = 0

	for _, secret := range initialSecretNumbers {
		for i := 0; i < 2000; i++ {
			secret = (secret ^ (secret * 64)) % 16777216
			secret = (secret ^ (secret / 32)) % 16777216
			secret = (secret ^ (secret * 2048)) % 16777216
		}
		sum += secret
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`1
		10
		100
		2024`

	var expected uint64 = 37327623
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/22.txt")
	t.Log(solution(string(input)))
}
