package main

import (
	"os"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

func parseInput(input string) [][]int64 {
	lines := aoc.SplitAndTrim(input, "\n")

	lineNums := make([][]int64, len(lines))

	for i, line := range lines {
		a := aoc.SplitAndTrim(line, ":")
		b := aoc.SplitAndTrim(a[1], " ")

		testValue, _ := strconv.Atoi(a[0])
		lineNums[i] = append(lineNums[i], int64(testValue))

		for _, value := range b {
			num, _ := strconv.Atoi(value)
			lineNums[i] = append(lineNums[i], int64(num))
		}
	}

	return lineNums
}

func hasCombination(testValue int64, numbers []int64) bool {
	n := len(numbers)

	combination := 0

	for i := 0; i < imath.Pow(2, n-1); i++ { // Every combination
		var value int64 = numbers[0]

		for k := 0; k < n-1; k++ {
			bitmask := 1 << k
			bit := (combination & bitmask) >> k

			if bit == 0 { // addition
				value = value + numbers[k+1]
			} else { // multiplication
				value = value * numbers[k+1]
			}
		}

		if value == testValue {
			return true
		}

		combination++
	}

	return false
}

func solution(input string) int64 {
	lines := parseInput(input)

	var sum int64 = 0

	for _, line := range lines {
		testValue, numbers := line[0], line[1:]
		if hasCombination(testValue, numbers) {
			sum += testValue
		}
	}

	return sum
}

func TestSolutionSample(t *testing.T) {
	input :=
		`190: 10 19
		 3267: 81 40 27
		 83: 17 5
		 156: 15 6
		 7290: 6 8 6 15
		 161011: 16 10 13
		 192: 17 8 14
		 21037: 9 7 18 13
		 292: 11 6 16 20`

	var expected int64 = 3749
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/07.txt")
	t.Log(solution(string(input)))
}
