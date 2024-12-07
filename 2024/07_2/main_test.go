package main

import (
	"fmt"
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

	operations := make([]byte, n-1)

	for i := 0; i < imath.Pow(3, n-1); i++ { // Every combination
		var value int64 = numbers[0]

		for j, operation := range operations {
			if operation == 0 { // addition
				value = value + numbers[j+1]
			} else if operation == 1 { // multiplication
				value = value * numbers[j+1]
			} else { // concat
				num, _ := strconv.Atoi(fmt.Sprintf("%v%v", value, numbers[j+1]))
				value = int64(num)
			}
		}

		if value == testValue {
			return true
		}

		for j := range operations {
			operations[j]++
			if operations[j] >= 3 {
				operations[j] = 0
				continue
			}
			break
		}
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

	var expected int64 = 11387
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := `192: 17 8 14`

	var expected int64 = 192
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := `161011: 16 10 13`

	var expected int64 = 0
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/07.txt")
	t.Log(solution(string(input)))
}
