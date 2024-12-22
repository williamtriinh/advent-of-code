package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

const PRUNE = 16777216

func parseInput(input string) []int {
	lines := aoc.SplitAndTrim(input, "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	return numbers
}

func getLastDigit(secret int) int {
	str := fmt.Sprint(secret)
	digit, _ := strconv.Atoi(str[len(str)-1:])
	return digit
}

func nextSecret(secret int) int {
	secret = (secret ^ (secret * 64)) % PRUNE
	secret = (secret ^ (secret / 32)) % PRUNE
	return (secret ^ (secret * 2048)) % PRUNE
}

func solution(input string, iterations int) int {
	initialSecretNumbers := parseInput(input)

	sequencesToBananas := map[[4]int]int{}

	for _, secret := range initialSecretNumbers {
		sequences := map[[4]int]int{}

		prices := make([]int, iterations+1)
		prices[0] = getLastDigit(secret)

		// Create a slice of digits for each iteration of the secret
		for i := 0; i < iterations; i++ {
			secret = nextSecret(secret)
			prices[i+1] = getLastDigit(secret)
		}

		// Add the first sequence of numbers to a map of sequence to price
		for i := 0; i <= iterations-4; i++ {
			sequence := [4]int{}

			for j := 0; j < 4; j++ {
				sequence[j] = prices[i+1+j] - prices[i+j]
			}

			if _, exists := sequences[sequence]; !exists {
				sequences[sequence] = prices[i+4]
			}
		}

		// Sum the prices of the same sequence together
		for sequence, price := range sequences {
			sequencesToBananas[sequence] += price
		}
	}

	return slices.Max(slices.Collect(maps.Values(sequencesToBananas)))
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`1
		2
		3
		2024`

	expected := 23
	received := solution(input, 2000)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "123"

	var expected int = 6
	received := solution(input, 9)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/22.txt")
	t.Log(solution(string(input), 2000))
}
