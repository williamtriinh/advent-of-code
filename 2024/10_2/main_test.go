package main

import (
	"container/list"
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var directions = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func parseInput(input string) [][]int {
	lines := aoc.SplitAndTrim(input, "\n")

	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))

		for j, b := range line {
			if b == '.' {
				grid[i][j] = -1
			} else {
				grid[i][j] = int(b - 48)
			}
		}
	}

	return grid
}

func solution(input string) int {
	grid := parseInput(input)
	sum := 0

	// Look for trailheads (0)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				stack := list.New()
				stack.PushFront([2]int{x, y})

				for stack.Len() > 0 {
					position := stack.Remove(stack.Front()).([2]int)
					positionHeight := grid[position[1]][position[0]]

					if positionHeight == 9 {
						sum++
						continue
					}

					for _, direction := range directions {
						adjacentPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}

						if adjacentPosition[0] < 0 || adjacentPosition[0] >= len(grid[position[1]]) || adjacentPosition[1] < 0 || adjacentPosition[1] >= len(grid) || grid[adjacentPosition[1]][adjacentPosition[0]] == -1 {
							continue
						}

						if grid[adjacentPosition[1]][adjacentPosition[0]] == positionHeight+1 {
							stack.PushFront(adjacentPosition)
						}
					}
				}
			}
		}
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`.....0.
		 ..4321.
		 ..5..2.
		 ..6543.
		 ..7..4.
	 	 ..8765.
		 ..9....`

	expected := 3
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input :=
		`..90..9
		 ...1.98
		 ...2..7
		 6543456
		 765.987
		 876....
		 987....`

	expected := 13
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input :=
		`012345
		 123456
		 234567
		 345678
		 4.6789
		 56789.`

	expected := 227
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample4(t *testing.T) {
	input :=
		`89010123
		 78121874
		 87430965
		 96549874
		 45678903
		 32019012
		 01329801
		 10456732`

	expected := 81
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/10.txt")
	t.Log(solution(string(input)))
}
