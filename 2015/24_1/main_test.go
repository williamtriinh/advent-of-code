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

	groupWeight := totalWeight / 3

	validFirstGroupCombinations := [][]int{}

	for i := 1; len(validFirstGroupCombinations) == 0; i++ {
		combinationIndices := combin.Combinations(len(weights), i)
		combinations := funk.Map(combinationIndices, func(indices []int) []int {
			return funk.Map(indices, func(index int) int {
				return weights[index]
			}).([]int)
		}).([][]int)

		for _, combination := range combinations {
			sum := funk.Reduce(combination, func(accumulator, current int) int {
				return accumulator + current
			}, 0).(int)

			if sum == groupWeight {
				validFirstGroupCombinations = append(validFirstGroupCombinations, combination)
			}
		}
	}

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

	expected := 99
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/24.txt")
	t.Log(solution(string(input)))
}
