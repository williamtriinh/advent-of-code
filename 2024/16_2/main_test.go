package main

import (
	"os"
	"testing"

	"github.com/emirpasic/gods/queues/priorityqueue"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var directions = [4]imath.Vec2{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

type Node struct {
	pos    imath.Vec2
	parent *Node
	score  int
}

func findStartAndEnd(grid [][]byte) (imath.Vec2, imath.Vec2) {
	var start, end imath.Vec2

	for row := range grid {
		for col := range grid[row] {
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

func getAdjacents(grid [][]byte, node *Node, facing imath.Vec2) []*Node {
	adjacents := []*Node{}

	for _, dir := range directions {
		adjPos := imath.Vec2{X: node.pos.X + dir.X, Y: node.pos.Y + dir.Y}
		adjCell := grid[adjPos.Y][adjPos.X]

		if adjCell == '.' || adjCell == 'E' {
			score := 1

			if facing.Dot(adjPos.Subtract(node.pos)) == 0 {
				score = 1001
			}

			adjacents = append(adjacents, &Node{
				pos:    adjPos,
				score:  score,
				parent: node,
			})
		}
	}

	return adjacents
}

func dijkstra(grid [][]byte, start, end imath.Vec2) *Node {
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.(*Node).score - b.(*Node).score
	})

	distances := map[imath.Vec2]int{}

	visited := map[imath.Vec2]struct{}{}

	startNode := &Node{
		pos:   start,
		score: 0,
	}

	facing := directions[0] // Initially facing right

	pq.Enqueue(startNode)

	distances[startNode.pos] = 0

	for !pq.Empty() {
		element, _ := pq.Dequeue()
		node := element.(*Node)

		if _, exists := visited[node.pos]; exists {
			continue
		}

		if node.pos.Equals(end) {
			return node
		}

		if node.parent != nil {
			facing = node.parent.pos.Subtract(node.pos)
		}

		for _, adjacent := range getAdjacents(grid, node, facing) {
			if _, exists := visited[adjacent.pos]; exists {
				continue
			}

			alternativeScore := distances[node.pos] + adjacent.score

			if _, exists := distances[adjacent.pos]; !exists || alternativeScore < distances[adjacent.pos] {
				distances[adjacent.pos] = alternativeScore
				adjacent.score = alternativeScore
				pq.Enqueue(adjacent)
			}
		}
	}

	return nil
}

func parseInput(input string) [][]byte {
	lines := aoc.SplitAndTrim(input, "\n")
	grid := make([][]byte, len(lines))

	for i, line := range lines {
		grid[i] = []byte(line)
	}

	return grid
}

func solution(input string) int {
	grid := parseInput(input)

	gridWithPaths := aoc.CloneGrid(grid)

	start, end := findStartAndEnd(grid)
	node := dijkstra(grid, start, end)

	optimalScore := node.score

	for node != nil {
		gridWithPaths[node.pos.Y][node.pos.X] = 'O'

		if node.parent == nil {
			break
		}

		gridCopy := aoc.CloneGrid(grid)
		gridCopy[node.parent.pos.Y][node.parent.pos.X] = '#'

		otherNode := dijkstra(gridCopy, start, end)
		if otherNode != nil && otherNode.score == optimalScore {
			for otherNode != nil {
				gridWithPaths[otherNode.pos.Y][otherNode.pos.X] = 'O'
				otherNode = otherNode.parent
			}
		}

		node = node.parent
	}

	tiles := 0

	for row := range gridWithPaths {
		for col := range gridWithPaths[row] {
			if gridWithPaths[row][col] == 'O' {
				tiles++
			}
		}
	}

	aoc.PrintGrid(gridWithPaths)

	return tiles
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`###############
		 #.......#....E#
		 #.#.###.#.###.#
		 #.....#.#...#.#
		 #.###.#####.#.#
		 #.#.#.......#.#
		 #.#.#####.###.#
		 #...........#.#
		 ###.#.#####.#.#
		 #...#.....#.#.#
		 #.#.#.###.#.#.#
		 #.....#...#.#.#
		 #.###.#.#.#.#.#
		 #S..#.....#...#
		 ###############`

	expected := 45
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input :=
		`#################
		 #...#...#...#..E#
		 #.#.#.#.#.#.#.#.#
		 #.#.#.#...#...#.#
		 #.#.#.#.###.#.#.#
		 #...#.#.#.....#.#
		 #.#.#.#.#.#####.#
		 #.#...#.#.#.....#
		 #.#.#####.#.###.#
		 #.#.#.......#...#
		 #.#.###.#####.###
		 #.#.#...#.....#.#
		 #.#.#.#####.###.#
		 #.#.#.........#.#
		 #.#.#.#########.#
		 #S#.............#
		 #################`

	expected := 64
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/16.txt")
	t.Log(solution(string(input)))
}
