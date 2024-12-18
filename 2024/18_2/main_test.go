package main

import (
	"container/list"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var directions = [4]imath.Vec2{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

func parseInput(input string) []imath.Vec2 {
	lines := aoc.SplitAndTrim(input, "\n")

	bytePositions := make([]imath.Vec2, len(lines))

	for i, line := range lines {
		values := strings.Split(line, ",")
		pos := imath.Vec2{}
		pos.X, _ = strconv.Atoi(values[0])
		pos.Y, _ = strconv.Atoi(values[1])
		bytePositions[i] = pos
	}

	return bytePositions
}

func createGrid(gridSize int) [][]byte {
	grid := make([][]byte, gridSize)

	for row := range grid {
		grid[row] = make([]byte, gridSize)

		for col := range grid[row] {
			grid[row][col] = '.'
		}
	}

	return grid
}

func getNeighbors(grid [][]byte, pos *imath.Vec2) []*imath.Vec2 {
	neighbors := make([]*imath.Vec2, 0, 4)

	for _, dir := range directions {
		neighborPos := pos.Add(dir)

		if neighborPos.X < 0 || neighborPos.X >= len(grid) || neighborPos.Y < 0 || neighborPos.Y >= len(grid) {
			continue
		}

		if grid[neighborPos.Y][neighborPos.X] != '.' {
			continue
		}

		neighbors = append(neighbors, &neighborPos)
	}

	return neighbors
}

func createPath(cameFrom map[imath.Vec2]*imath.Vec2, endPos *imath.Vec2) []*imath.Vec2 {
	pos := endPos
	path := []*imath.Vec2{}

	for pos != nil {
		path = append(path, pos)
		pos = cameFrom[*pos]
	}

	slices.Reverse(path)

	return path
}

func dijkstras(grid [][]byte) []*imath.Vec2 {
	queue := list.New()
	cameFrom := make(map[imath.Vec2]*imath.Vec2, len(grid))

	startPos := &imath.Vec2{X: 0, Y: 0}
	endPos := &imath.Vec2{X: len(grid) - 1, Y: len(grid) - 1}

	queue.PushBack(startPos)
	cameFrom[*startPos] = nil

	for queue.Len() > 0 {
		pos := queue.Remove(queue.Front()).(*imath.Vec2)

		if pos.Equals(*endPos) {
			return createPath(cameFrom, endPos)
		}

		for _, neighbor := range getNeighbors(grid, pos) {
			if _, exists := cameFrom[*neighbor]; !exists {
				queue.PushBack(neighbor)
				cameFrom[*neighbor] = pos
			}
		}
	}

	return []*imath.Vec2{}
}

func solution(gridSize int, input string) string {
	bytePositions := parseInput(input)
	grid := createGrid(gridSize)

	for _, bytePos := range bytePositions {
		grid[bytePos.Y][bytePos.X] = '#'
		path := dijkstras(grid)

		if len(path) == 0 {
			return fmt.Sprintf("%v,%v", bytePos.X, bytePos.Y)
		}
	}

	return "FAIL"
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`5,4
		4,2
		4,5
		3,0
		2,1
		6,3
		2,4
		1,5
		0,6
		3,3
		2,6
		5,1
		1,2
		5,5
		2,5
		6,5
		1,4
		0,4
		6,4
		1,1
		6,1
		1,0
		0,5
		1,6
		2,0`

	expected := "6,1"
	received := solution(7, input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/18.txt")
	t.Log(solution(71, string(input)))
}
