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
			grid[i][j] = int(b - 48)
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
				visited := make([][]bool, len(grid))
				for i := range visited {
					visited[i] = make([]bool, len(grid[i]))
				}

				queue := list.New()
				queue.PushBack([2]int{x, y})
				visited[y][x] = true

				for queue.Len() > 0 {
					position := queue.Remove(queue.Front()).([2]int)
					positionHeight := grid[position[1]][position[0]]

					if positionHeight == 9 {
						sum++
					}

					visited[position[1]][position[0]] = true

					for _, direction := range directions {
						adjacentPosition := [2]int{position[0] + direction[0], position[1] + direction[1]}

						if adjacentPosition[0] < 0 || adjacentPosition[0] >= len(grid[position[1]]) || adjacentPosition[1] < 0 || adjacentPosition[1] >= len(grid) {
							continue
						}

						if grid[adjacentPosition[1]][adjacentPosition[0]] == positionHeight+1 && !visited[adjacentPosition[1]][adjacentPosition[0]] {
							queue.PushBack(adjacentPosition)
							visited[adjacentPosition[1]][adjacentPosition[0]] = true
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
		`0123
		 1234
		 8765
		 9876`

	expected := 1
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input :=
		`89010123
	   78121874
	   87430965
	   96549874
	   45678903
	   32019012
	   01329801
	   10456732`

	expected := 36
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/10.txt")
	t.Log(solution(string(input)))
}
