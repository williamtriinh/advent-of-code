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

func isWithinBounds(pos imath.Vec2, grid aoc.Grid) bool {
	return pos.X > 0 && pos.X < grid.Width()-1 && pos.Y > 0 && pos.Y < grid.Height()-1
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

func bfs(start imath.Vec2, end imath.Vec2, grid aoc.Grid) ([]*Step, map[imath.Vec2]*Step) {
	queue := list.New()
	visited := make(map[imath.Vec2]struct{}, imath.Pow(len(grid), 2))

	path := []*Step{}
	stepMap := map[imath.Vec2]*Step{}

	queue.PushBack(start)
	visited[start] = struct{}{}

	for queue.Len() > 0 {
		pos := queue.Remove(queue.Front()).(imath.Vec2)

		path = append(path, &Step{pos: pos, index: len(path)})
		stepMap[pos] = path[len(path)-1]

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

	return path, stepMap
}

func solution(input string) int {
	grid := parseInput(input)
	start, end := getStartAndEndPos(grid)

	path, stepMap := bfs(start, end, grid)
	answer := 0

	for _, step := range path {
		for _, direction := range directions {
			// Check for wall first
			otherPos := step.pos.Add(direction)

			if !isWithinBounds(otherPos, grid) {
				continue
			}

			if grid[otherPos.Y][otherPos.X] != '#' {
				continue
			}

			// Check for an empty cell on the other side of the wall
			otherPos = otherPos.Add(direction)

			if !isWithinBounds(otherPos, grid) {
				continue
			}

			if grid[otherPos.Y][otherPos.X] != '.' {
				continue
			}

			if stepMap[otherPos].index-step.index-1 >= 100 {
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
