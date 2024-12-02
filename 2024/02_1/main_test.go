package main

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/2024/imath"
)

func areLevelsSafe(levels []string) bool {
	lastDiff := 0

	for i := 0; i < len(levels)-1; i++ {
		a, _ := strconv.Atoi(levels[i])
		b, _ := strconv.Atoi(levels[i+1])

		diff := b - a

		if imath.Abs(diff) < 1 || imath.Abs(diff) > 3 || lastDiff > 0 && diff < 0 || lastDiff < 0 && diff > 0 {
			return false
		}

		lastDiff = diff
	}

	return true
}

func solution(input string) int {
	lines := strings.Split(input, "\n")
	answer := 0

	for _, line := range lines {
		levels := strings.Split(strings.TrimSpace(line), " ")

		if areLevelsSafe(levels) {
			answer++
		}
	}

	return answer
}

func TestSolutionSample(t *testing.T) {
	input :=
		`7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9`

	want := 2
	got := solution(input)

	if got != want {
		t.Errorf("Got %v expected %v", got, want)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/02.txt")
	t.Log(solution(string(input)))
}
