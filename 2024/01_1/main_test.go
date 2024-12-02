package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/2024/imath"
)

func solution(input string) int {
	lines := strings.Split(input, "\n")

	numsLeft := make([]int, len(lines))
	numsRight := make([]int, len(lines))

	for i, line := range lines {
		values := strings.Split(strings.TrimSpace(line), "   ")
		numsLeft[i], _ = strconv.Atoi(values[0])
		numsRight[i], _ = strconv.Atoi(values[1])
	}

	slices.Sort(numsLeft)
	slices.Sort(numsRight)

	sum := 0
	for i := 0; i < len(numsLeft); i++ {
		sum += imath.Abs(numsLeft[i] - numsRight[i])
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

	expected := 11
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile(".../inputs/01.txt")
	t.Log(solution(string(input)))
}
