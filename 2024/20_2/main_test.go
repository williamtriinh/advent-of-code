package main

import (
	"container/list"
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var directions = [4]imath.Vec2{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

type Step struct {
	pos   imath.Vec2
	index int
}

func parseInput(input string) aoc.Grid {
	lines := aoc.SplitAndTrim(input, "\n")

	grid := make(aoc.Grid, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func getStartAndEndPos(grid aoc.Grid) (imath.Vec2, imath.Vec2) {
	var start, end imath.Vec2

	for row := range grid {
		for col := range grid {
			if grid[row][col] == 'S' {
				start = imath.Vec2{X: col, Y: row}
			}

			if grid[row][col] == 'E' {
				end = imath.Vec2{X: col, Y: row}
			}

			if !start.IsZero() && !end.IsZero() {
				return start, end
			}
		}
	}

	return start, end
}

func getManhattanDistance(a, b imath.Vec2) int {
	return imath.Abs(a.X-b.X) + imath.Abs(a.Y-b.Y)
}

func getNeighbors(grid aoc.Grid, pos imath.Vec2) []imath.Vec2 {
	neighbors := make([]imath.Vec2, 0, 4)

	for _, direction := range directions {
		otherPos := pos.Add(direction)
		if grid[otherPos.Y][otherPos.X] == '.' || grid[otherPos.Y][otherPos.X] == 'E' {
			neighbors = append(neighbors, otherPos)
		}
	}

	return neighbors
}

func bfs(start imath.Vec2, end imath.Vec2, grid aoc.Grid) []*Step {
	queue := list.New()
	visited := make(map[imath.Vec2]struct{}, imath.Pow(len(grid), 2))

	path := []*Step{}

	queue.PushBack(start)
	visited[start] = struct{}{}

	for queue.Len() > 0 {
		pos := queue.Remove(queue.Front()).(imath.Vec2)

		path = append(path, &Step{pos: pos, index: len(path)})

		if pos.Equals(end) {
			break
		}

		for _, neighbor := range getNeighbors(grid, pos) {
			if _, exists := visited[neighbor]; !exists {
				queue.PushBack(neighbor)
				visited[neighbor] = struct{}{}
			}
		}
	}

	return path
}

func solution(input string) int {
	grid := parseInput(input)
	start, end := getStartAndEndPos(grid)

	path := bfs(start, end, grid)
	answer := 0

	for _, step1 := range path {
		for _, step2 := range path {
			manhattanDistance := getManhattanDistance(step1.pos, step2.pos)
			if manhattanDistance <= 20 && step2.index-step1.index-manhattanDistance >= 100 {
				answer++
			}
		}
	}

	return answer
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/20.txt")
	t.Log(solution(string(input)))
}
