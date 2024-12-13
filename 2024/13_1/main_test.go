package main

import (
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var A_TOKEN int = 3
var B_TOKEN int = 1

var xyRegex = regexp.MustCompile(`(\d+)`)

func parseXY(line string) [2]int {
	match := xyRegex.FindAllStringSubmatch(line, -1)

	x, _ := strconv.Atoi(match[0][1])
	y, _ := strconv.Atoi(match[1][1])

	return [2]int{x, y}
}

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")

	tokens := 0

	for i := 0; i < len(lines); i += 4 {
		buttonA := parseXY(lines[i])   // 3 tokens
		buttonB := parseXY(lines[i+1]) // 1 token
		prize := parseXY(lines[i+2])

		determinant := buttonA[0]*buttonB[1] - buttonA[1]*buttonB[0]

		pressesA := (prize[0]*buttonB[1] - prize[1]*buttonB[0]) / determinant
		pressesB := (buttonA[0]*prize[1] - buttonA[1]*prize[0]) / determinant

		if buttonA[0]*pressesA+buttonB[0]*pressesB == prize[0] && buttonA[1]*pressesA+buttonB[1]*pressesB == prize[1] {
			tokens += A_TOKEN*pressesA + B_TOKEN*pressesB
		}

	}

	return tokens
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`Button A: X+94, Y+34
		 Button B: X+22, Y+67
		 Prize: X=8400, Y=5400
		
		 Button A: X+26, Y+66
		 Button B: X+67, Y+21
		 Prize: X=12748, Y=12176
		
		 Button A: X+17, Y+86
		 Button B: X+84, Y+37
		 Prize: X=7870, Y=6450
		
		 Button A: X+69, Y+23
		 Button B: X+27, Y+71
		 Prize: X=18641, Y=10279`

	expected := 480
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/13.txt")
	t.Log(solution(string(input)))
}
