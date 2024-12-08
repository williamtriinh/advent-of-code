package main

import (
	"os"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var directions = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

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

func doesLoop(lines []string, x, y int) bool {
	height := len(lines)
	width := len(lines[0])

	visited := make([][]byte, height)
	for i := range visited {
		visited[i] = make([]byte, width)
		for j := range visited[i] {
			visited[i][j] = 4
		}
	}

	facing := 0

	for x >= 0 && x < width && y >= 0 && y < height {
		if visited[y][x] == byte(facing) {
			// Looped
			return true
		}

		if visited[y][x] == 4 {
			visited[y][x] = byte(facing)
		}

		for { // Keep turning until we don't reach an obstacle or out of bounds
			dx := x + directions[facing][0]
			dy := y + directions[facing][1]

			// Guard is about to leave the map
			if dx < 0 || dx >= width || dy < 0 || dy >= height {
				return false
			}

			if lines[dy][dx] == '#' {
				facing = (facing + 1) % len(directions)
				continue
			}

			break
		}

		x += directions[facing][0]
		y += directions[facing][1]
	}

	return false
}

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")
	height := len(lines)
	width := len(lines[0])

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	startX, startY := getGuardInitialPosition(lines)
	x, y := startX, startY
	facing := 0

	for x >= 0 && x < width && y >= 0 && y < height {
		visited[y][x] = true

		dx := x + directions[facing][0]
		dy := y + directions[facing][1]

		// Guard is about to leave the map
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
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] && lines[i][j] == '.' {
				linesCopy := make([]string, len(lines))
				for i := range linesCopy {
					linesCopy[i] = strings.Clone(lines[i])
				}

				modified1 := []rune(linesCopy[i])
				modified1[j] = '#'
				linesCopy[i] = string(modified1)

				modified2 := []rune(linesCopy[startY])
				modified2[startX] = '.'
				linesCopy[startY] = string(modified2)

				if doesLoop(linesCopy, startX, startY) {
					sum++
				}
			}
		}
	}

	return sum
}

func TestDoesLoopSample1(t *testing.T) {
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
		 ......##..`

	parsedInput := aoc.SplitAndTrim(input, "\n")
	x, y := getGuardInitialPosition(parsedInput)

	expected := true
	received := doesLoop(parsedInput, x, y)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestDoesLoopSample2(t *testing.T) {
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
		 ......#..`

	parsedInput := aoc.SplitAndTrim(input, "\n")
	x, y := getGuardInitialPosition(parsedInput)

	expected := false
	received := doesLoop(parsedInput, x, y)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample1(t *testing.T) {
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

	expected := 6
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/06.txt")
	t.Log(solution(string(input)))
}
