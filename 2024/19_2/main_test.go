package main

import (
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var patterns map[string]struct{}
var largestPattern int

func parseInput(input string) []string {
	lines := aoc.SplitAndTrim(input, "\n")

	patternsSlice := aoc.SplitAndTrim(lines[0], ",")
	patterns = make(map[string]struct{}, len(patternsSlice))

	largestPattern = 1

	for _, pattern := range patternsSlice {
		patterns[pattern] = struct{}{}

		if len(pattern) > largestPattern {
			largestPattern = len(pattern)
		}
	}

	designs := lines[2:]

	return designs
}

func recursion(design string, cache map[string]int) int {
	if design == "" {
		return 1
	}

	if _, exists := cache[design]; exists {
		return cache[design]
	}

	count := 0

	for i := 1; i <= imath.Min(len(design), largestPattern); i++ {
		window := design[:i]

		if _, exists := patterns[window]; exists {
			count += recursion(design[i:], cache)
		}
	}

	cache[design] = count
	return count
}

func solution(input string) int {
	designs := parseInput(input)
	cache := map[string]int{}

	answer := 0

	for _, design := range designs {
		answer += recursion(design, cache)
	}

	return answer
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`r, wr, b, g, bwu, rb, gb, br

		brwrr
		bggr
		gbbr
		rrbgbr
		ubwu
		bwurrg
		brgr
		bbrgwb`

	expected := 16
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/19.txt")
	t.Log(solution(string(input)))
}
