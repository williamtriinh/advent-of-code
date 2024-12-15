package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var directions = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func parseInput(input string) ([][]byte, string) {
	lines := aoc.SplitAndTrim(input, "\n")

	grid := [][]byte{}
	builder := strings.Builder{}

	buildingGrid := true

	for _, line := range lines {
		if line == "" {
			buildingGrid = false
			continue
		}

		if buildingGrid {
			wideLine := make([]byte, len(line)*2)
			for i, r := range line {
				if r == 'O' {
					wideLine[i*2] = '['
					wideLine[i*2+1] = ']'
				} else if r == '@' {
					wideLine[i*2] = '@'
					wideLine[i*2+1] = '.'
				} else {
					wideLine[i*2] = byte(r)
					wideLine[i*2+1] = byte(r)
				}
			}
			grid = append(grid, wideLine)
		} else {
			builder.WriteString(line)
		}
	}

	return grid, builder.String()
}

func getRobotPos(grid [][]byte) [2]int {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '@' {
				return [2]int{col, row}
			}
		}
	}

	return [2]int{-1, -1}
}

func convertMoveToDirection(move rune) [2]int {
	if move == '>' {
		return directions[0]
	}

	if move == 'v' {
		return directions[1]
	}

	if move == '<' {
		return directions[2]
	}

	return directions[3]
}

func moveAndPush(pos [2]int, dir [2]int, grid [][]byte, commit bool) bool {
	nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

	if grid[nextPos[1]][nextPos[0]] == '#' {
		return false
	}

	if dir[1] == 0 {
		if grid[nextPos[1]][nextPos[0]] == '.' || moveAndPush(nextPos, dir, grid, commit) {
			if commit {
				grid[nextPos[1]][nextPos[0]] = grid[pos[1]][pos[0]]
				grid[pos[1]][pos[0]] = '.'
			}
			return true
		}
	} else {
		if grid[nextPos[1]][nextPos[0]] == ']' {
			otherPos := [2]int{nextPos[0] - 1, nextPos[1]}
			if moveAndPush(nextPos, dir, grid, commit) && moveAndPush(otherPos, dir, grid, commit) {
				if commit {
					grid[nextPos[1]][nextPos[0]] = grid[pos[1]][pos[0]]
					grid[pos[1]][pos[0]] = '.'
					grid[otherPos[1]][otherPos[0]] = '.'
				}
				return true
			}
			return false
		}

		if grid[nextPos[1]][nextPos[0]] == '[' {
			otherPos := [2]int{nextPos[0] + 1, nextPos[1]}
			if moveAndPush(nextPos, dir, grid, commit) && moveAndPush(otherPos, dir, grid, commit) {
				if commit {
					grid[nextPos[1]][nextPos[0]] = grid[pos[1]][pos[0]]
					grid[pos[1]][pos[0]] = '.'
					grid[otherPos[1]][otherPos[0]] = '.'
				}
				return true
			}
			return false
		}

		if commit {
			grid[nextPos[1]][nextPos[0]] = grid[pos[1]][pos[0]]
			grid[pos[1]][pos[0]] = '.'
		}
		return true
	}

	return false
}

func sumBoxCoordinates(grid [][]byte) int {
	sum := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '[' {
				sum += 100*row + col
			}
		}
	}

	return sum
}

func printGrid(grid [][]byte) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Printf("%c", grid[row][col])
		}
		fmt.Println()
	}
}

func solution(input string) int {
	grid, moves := parseInput(input)
	robotPos := getRobotPos(grid)

	for _, move := range moves {
		dir := convertMoveToDirection(move)
		if moveAndPush(robotPos, dir, grid, false) {
			moveAndPush(robotPos, dir, grid, true)
			robotPos[0] += dir[0]
			robotPos[1] += dir[1]
		}
	}

	printGrid(grid)

	return sumBoxCoordinates(grid)
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`##########
		 #..O..O.O#
		 #......O.#
		 #.OO..O.O#
		 #..O@..O.#
		 #O#..O...#
		 #O..O..O.#
		 #.OO.O.OO#
		 #....O...#
		 ##########
		
		 <vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
		 vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
		 ><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
		 <<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
		 ^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
		 ^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
		 >^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
		 <><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
		 ^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
		 v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	expected := 9021
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/15.txt")
	t.Log(solution(string(input)))
}
