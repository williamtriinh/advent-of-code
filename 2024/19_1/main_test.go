package main

import (
	"container/list"
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

func parseInput(input string) (map[string]struct{}, []string, int) {
	lines := aoc.SplitAndTrim(input, "\n")

	patterns := aoc.SplitAndTrim(lines[0], ",")
	patternsMap := make(map[string]struct{}, len(patterns))

	largestPattern := 1

	for _, pattern := range patterns {
		patternsMap[pattern] = struct{}{}

		if len(pattern) > largestPattern {
			largestPattern = len(pattern)
		}
	}

	designs := lines[2:]

	return patternsMap, designs, largestPattern
}

// Using DFS
func solution(input string) int {
	patterns, designs, largestPattern := parseInput(input)

	answer := 0

	for _, design := range designs {
		stack := list.New()
		visited := map[int]struct{}{}

		matched := false

		// Initial stack setup
		for i := 1; i <= largestPattern; i++ {
			window := design[0:i]

			if _, exists := patterns[window]; exists {
				stack.PushFront(i)
			}
		}

		for stack.Len() > 0 {
			index := stack.Front().Value.(int)

			if _, exists := visited[index]; exists {
				stack.Remove(stack.Front())
				continue
			}

			visited[index] = struct{}{}

			// Try to match against every pattern size
			for i := imath.Min(len(design), index+largestPattern); i > index; i-- {
				window := design[index:i]

				if _, exists := patterns[window]; exists {
					if i == len(design) {
						matched = true
						break
					}
					stack.PushFront(i)
				}
			}

			if matched {
				break
			}
		}

		if matched {
			answer++
		}
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

	expected := 6
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/19.txt")
	t.Log(solution(string(input)))
}
