package main

import (
	"container/list"
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var directions = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func solution(input string) int {
	grid := aoc.SplitAndTrim(input, "\n")

	visited := make([][]bool, len(grid))
	for row := range visited {
		visited[row] = make([]bool, len(grid[row]))
	}

	answer := 0

	queue := list.New()

	for row := range grid {
		for col := range grid[row] {
			if !visited[row][col] {
				queue.PushBack([2]int{col, row})
				visited[row][col] = true
				regionType := grid[row][col]

				area := 0
				perimeter := 0

				for queue.Len() > 0 {
					pos := queue.Remove(queue.Front()).([2]int)

					area++

					adjacentCount := 0
					for _, dir := range directions {
						dx := pos[0] + dir[0]
						dy := pos[1] + dir[1]

						if dx < 0 || dx >= len(grid[0]) || dy < 0 || dy >= len(grid) {
							continue
						}

						if regionType == grid[dy][dx] {
							if !visited[dy][dx] {
								queue.PushBack([2]int{dx, dy})
								visited[dy][dx] = true
							}
							adjacentCount++
						}
					}

					perimeter += 4 - adjacentCount
				}

				answer += area * perimeter
			}
		}
	}

	return answer
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`AAAA
		 BBCD
		 BBCC
		 EEEC`

	expected := 140
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input :=
		`OOOOO
		 OXOXO
		 OOOOO
		 OXOXO
		 OOOOO`

	expected := 772
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input :=
		`RRRRIICCFF
		 RRRRIICCCF
		 VVRRRCCFFF
		 VVRCCCJFFF
		 VVVVCJJCFE
		 VVIVCCJJEE
		 VVIIICJJEE
		 MIIIIIJJEE
		 MIIISIJEEE
		 MMMISSJEEE`

	expected := 1930
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/12.txt")
	t.Log(solution(string(input)))
}
