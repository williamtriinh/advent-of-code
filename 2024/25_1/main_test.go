package main

import (
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

// The grids are 5x7
func parseInput(input string) (map[[5]int]struct{}, map[[5]int]struct{}) {
	lines := aoc.SplitAndTrim(input, "\n")

	locks := map[[5]int]struct{}{}
	keys := map[[5]int]struct{}{}

	for i := 0; i < len(lines); i += 8 {
		height := [5]int{}

		// Ignore the top and bottom rows (hence the range [1,5]) as they don't count
		// towards the height of each column
		for j := 1; j < 6; j++ {
			line := lines[i+j]

			for k, char := range line {
				if char == '#' {
					height[k]++
				}
			}
		}

		if lines[i][0] == '#' {
			locks[height] = struct{}{}
		} else {
			keys[height] = struct{}{}
		}
	}

	return locks, keys
}

func solution(input string) int {
	locks, keys := parseInput(input)

	answer := 0

	for lock := range locks {
		for key := range keys {
			fits := true

			for i := range lock {
				if key[i]+lock[i] > 5 {
					fits = false
					break
				}
			}

			if fits {
				answer++
			}
		}
	}

	return answer
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`#####
		.####
		.####
		.####
		.#.#.
		.#...
		.....
		
		#####
		##.##
		.#.##
		...##
		...#.
		...#.
		.....
		
		.....
		#....
		#....
		#...#
		#.#.#
		#.###
		#####
		
		.....
		.....
		#.#..
		###..
		###.#
		###.#
		#####
		
		.....
		.....
		.....
		#....
		#.#..
		#.#.#
		#####`

	expected := 3
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/25.txt")
	t.Log(solution(string(input)))
}
