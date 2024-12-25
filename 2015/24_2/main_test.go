package main

import (
	"os"
	"slices"
	"strconv"
	"testing"

	"github.com/thoas/go-funk"
	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"gonum.org/v1/gonum/stat/combin"
)

func parseInput(input string) []int {
	lines := aoc.SplitAndTrim(input, "\n")
	weights := make([]int, len(lines))

	for i, line := range lines {
		weights[i], _ = strconv.Atoi(line)
	}

	return weights
}

func solution(input string) int {
	weights := parseInput(input)

	totalWeight := funk.Reduce(weights, func(accumulator, current int) int {
		return accumulator + current
	}, 0).(int)

	groupWeight := totalWeight / 4

	validFirstGroupCombinations := [][]int{}

	// Iterate over every combination of packages for the first group. Start with
	// the smallest subset of elements to get the fewest packages as possible.
	for i := 1; len(validFirstGroupCombinations) == 0; i++ {
		// Returns the indices of the combinations
		combinationIndices := combin.Combinations(len(weights), i)

		// Builds the combination using the actual weights instead of the indices
		combinations := funk.Map(combinationIndices, func(indices []int) []int {
			return funk.Map(indices, func(index int) int {
				return weights[index]
			}).([]int)
		}).([][]int)

		// For each combination, if it is equal to the group weight, then add it to
		// the list of valid first group combinations
		for _, combination := range combinations {
			sum := funk.Reduce(combination, func(accumulator, current int) int {
				return accumulator + current
			}, 0).(int)

			if sum == groupWeight {
				validFirstGroupCombinations = append(validFirstGroupCombinations, combination)
			}
		}
	}

	// Calculate the quantum entanglement for every valid combination, sort it and
	// take the smallest one
	quantumEntabglements := funk.Map(validFirstGroupCombinations, func(combination []int) int {
		return funk.Reduce(combination, func(accumulator, current int) int {
			return accumulator * current
		}, 1).(int)
	}).([]int)

	slices.Sort(quantumEntabglements)

	return quantumEntabglements[0]
}

func TestSolutionSample(t *testing.T) {
	input :=
		`1
		2
		3
		4
		5
		7
		8
		9
		10
		11`

	expected := 44
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/24.txt")
	t.Log(solution(string(input)))
}
