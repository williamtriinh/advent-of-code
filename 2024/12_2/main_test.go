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
				corners := 0 // # of corners = # of sides

				for queue.Len() > 0 {
					pos := queue.Remove(queue.Front()).([2]int)

					area++

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
						}
					}

					// Find corners
					a := 0
					b := 1

					for i := 0; i < len(directions); i++ {
						dirA := directions[a]
						dirB := directions[b]

						adjAPos := [2]int{pos[0] + dirA[0], pos[1] + dirA[1]}
						adjBPos := [2]int{pos[0] + dirB[0], pos[1] + dirB[1]}

						var adjA, adjB byte

						if adjAPos[0] >= 0 && adjAPos[0] < len(grid[0]) && adjAPos[1] >= 0 && adjAPos[1] < len(grid) {
							adjA = grid[adjAPos[1]][adjAPos[0]]
						}

						if adjBPos[0] >= 0 && adjBPos[0] < len(grid[0]) && adjBPos[1] >= 0 && adjBPos[1] < len(grid) {
							adjB = grid[adjBPos[1]][adjBPos[0]]
						}

						// Outside corner
						if adjA != regionType && adjB != regionType {
							corners++
						}

						// Inside corner
						if adjA == regionType && adjB == regionType && grid[adjAPos[1]+dirB[1]][adjAPos[0]+dirB[0]] != regionType {
							corners++
						}

						a = (a + 1) % len(directions)
						b = (b + 1) % len(directions)
					}
				}

				answer += area * corners
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

	expected := 80
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

	expected := 436
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input :=
		`EEEEE
		 EXXXX
		 EEEEE
		 EXXXX
		 EEEEE`

	expected := 236
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample4(t *testing.T) {
	input :=
		`AAAAAA
		 AAABBA
		 AAABBA
		 ABBAAA
		 ABBAAA
		 AAAAAA`

	expected := 368
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample5(t *testing.T) {
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

	expected := 1206
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/12.txt")
	t.Log(solution(string(input)))
}
