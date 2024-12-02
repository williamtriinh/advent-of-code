package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func solution(input string) int {
	lines := strings.Split(input, "\n")

	numsLeft := make(map[int]int, len(lines))
	numsRight := make(map[int]int, len(lines))

	for _, line := range lines {
		values := strings.Split(strings.TrimSpace(line), "   ")

		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])

		numsLeft[left]++
		numsRight[right]++
	}

	sum := 0
	for key, frequency1 := range numsLeft {
		if frequency2, exists := numsRight[key]; exists {
			sum += key * frequency2 * frequency1
		}
	}

	return sum
}

func TestSolutionSample(t *testing.T) {
	input :=
		`3   4
		4   3
		2   5
		1   3
		3   9
		3   3`

	expected := 31
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/01.txt")
	t.Log(solution(string(input)))
}
