package main

import (
	"os"
	"slices"
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

func canResolveLevels(levels []string) bool {
	if areLevelsSafe(levels) {
		return true
	}

	for i := range levels {
		newLevels := slices.Concat(levels[:i], levels[i+1:])
		if areLevelsSafe(newLevels) {
			return true
		}
	}

	return false
}

func solution(input string) int {
	lines := strings.Split(input, "\n")
	answer := 0

	for _, line := range lines {
		levels := strings.Split(strings.TrimSpace(line), " ")

		if canResolveLevels(levels) {
			answer++
		}
	}

	return answer
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9`

	want := 4
	got := solution(input)

	if got != want {
		t.Errorf("Got %v expected %v", got, want)
	}
}

func TestSolutionSample2(t *testing.T) {
	input :=
		`48 46 47 49 51 54 56
		1 1 2 3 4 5
		1 2 3 4 5 5
		5 1 2 3 4 5
		1 4 3 2 1
		1 6 7 8 9
		1 2 3 4 3
		9 8 7 6 7
		7 10 8 10 11
		29 28 27 25 26 25 22 20`

	want := 10
	got := solution(input)

	if got != want {
		t.Errorf("Got %v expected %v", got, want)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/02.txt")
	t.Log(solution(string(input)))
}
