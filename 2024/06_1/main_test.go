package main

import (
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func getGuardInitialPosition(grid []string) (int, int) {
	for y, row := range grid {
		for x, col := range row {
			if col == '^' {
				return x, y
			}
		}
	}

	return -1, -1
}

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")
	height := len(lines)
	width := len(lines[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	x, y := getGuardInitialPosition(lines)
	facing := 0

	for x >= 0 && x < width && y >= 0 && y < height {
		visited[y][x] = true

		dx := x + directions[facing][0]
		dy := y + directions[facing][1]

		if dx < 0 || dx >= width || dy < 0 || dy >= height {
			break
		}

		if lines[dy][dx] == '#' {
			facing = (facing + 1) % len(directions)
		}

		x += directions[facing][0]
		y += directions[facing][1]
	}

	sum := 0
	for _, row := range visited {
		for _, col := range row {
			if col {
				sum++
			}
		}
	}

	return sum
}

func TestSolutionSample(t *testing.T) {
	input :=
		`....#.....
		 .........#
		 ..........
		 ..#.......
		 .......#..
		 ..........
		 .#..^.....
		 ........#.
		 #.........
		 ......#...`

	expected := 41
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/06.txt")
	t.Log(solution(string(input)))
}
