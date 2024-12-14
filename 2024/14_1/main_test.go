package main

import (
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

func solution(width, height int, input string) int {
	lines := aoc.SplitAndTrim(input, "\n")

	robots := make([][2][2]int, len(lines))
	for i, line := range lines {
		p, v := parseLine(line)
		robots[i] = [2][2]int{p, v}
	}

	// Simulate robots
	for i, robot := range robots {
		p := robot[0]
		v := robot[1]

		for j := 0; j < 100; j++ {
			p[0] = (p[0] + v[0]) % width

			if p[0] < 0 {
				p[0] += width
			}

			p[1] = (p[1] + v[1]) % height

			if p[1] < 0 {
				p[1] += height
			}
		}

		robots[i][0] = p
	}

	// Count robots in each quadrant
	quadrants := make([]int, 4)

	for _, robot := range robots {
		p := robot[0]

		if p[0] < width/2 && p[1] < height/2 {
			quadrants[0]++
		} else if p[0] > width/2 && p[1] < height/2 {
			quadrants[1]++
		} else if p[0] < width/2 && p[1] > height/2 {
			quadrants[2]++
		} else if p[0] > width/2 && p[1] > height/2 {
			quadrants[3]++
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`p=0,4 v=3,-3
		 p=6,3 v=-1,-3
		 p=10,3 v=-1,2
		 p=2,0 v=2,-1
		 p=0,0 v=1,3
		 p=3,0 v=-2,-2
		 p=7,6 v=-1,-3
		 p=3,0 v=-1,-2
		 p=9,3 v=2,3
		 p=7,3 v=-1,2
		 p=2,4 v=2,-3
		 p=9,5 v=-3,-3`

	expected := 12
	received := solution(11, 7, input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/14.txt")
	t.Log(solution(101, 103, string(input)))
}
