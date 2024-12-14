package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var regex = regexp.MustCompile(`(\d+)|(-\d+)`)

func parseLine(line string) ([2]int, [2]int) {
	match := regex.FindAllStringSubmatch(line, -1)

	px, _ := strconv.Atoi(match[0][0])
	py, _ := strconv.Atoi(match[1][0])
	vx, _ := strconv.Atoi(match[2][0])
	vy, _ := strconv.Atoi(match[3][0])

	return [2]int{px, py}, [2]int{vx, vy}
}

func solution(width, height int, input string) {
	lines := aoc.SplitAndTrim(input, "\n")

	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	robots := make([][2][2]int, len(lines))
	for i, line := range lines {
		p, v := parseLine(line)
		robots[i] = [2][2]int{p, v}
	}

	for seconds := 1; seconds <= width*height; seconds++ {
		// Simulate robots
		for i, robot := range robots {
			p := robot[0]
			v := robot[1]

			p[0] = (p[0] + v[0]) % width

			if p[0] < 0 {
				p[0] += width
			}

			p[1] = (p[1] + v[1]) % height

			if p[1] < 0 {
				p[1] += height
			}

			robots[i][0] = p
		}

		// Find robots that are located consecutively
		for _, robot := range robots {
			p := robot[0]
			grid[p[1]][p[0]]++
		}

		largestConsecutiveRobotos := 0
		consecutiveRobots := 0

		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == 0 {
					consecutiveRobots = 0
				} else {
					consecutiveRobots++
				}

				if consecutiveRobots > largestConsecutiveRobotos {
					largestConsecutiveRobotos = consecutiveRobots
				}
			}
		}

		if largestConsecutiveRobotos > 15 { // 15 is arbitrary
			fmt.Println(seconds)
			for row := range grid {
				for col := range grid[row] {
					if grid[row][col] == 0 {
						fmt.Print(".")
					} else {
						fmt.Print(grid[row][col])
					}
				}
				fmt.Println()
			}
		}

		for row := range grid {
			for col := range grid[row] {
				grid[row][col] = 0
			}
		}
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/14.txt")
	solution(101, 103, string(input))
}
