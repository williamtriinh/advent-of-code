package main

import (
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var A_TOKEN int64 = 3
var B_TOKEN int64 = 1

var xyRegex = regexp.MustCompile(`(\d+)`)

func parseXY(line string) [2]int64 {
	match := xyRegex.FindAllStringSubmatch(line, -1)

	x, _ := strconv.Atoi(match[0][1])
	y, _ := strconv.Atoi(match[1][1])

	return [2]int64{int64(x), int64(y)}
}

func solution(input string) int64 {
	lines := aoc.SplitAndTrim(input, "\n")

	var tokens int64 = 0

	for i := 0; i < len(lines); i += 4 {
		buttonA := parseXY(lines[i])   // 3 tokens
		buttonB := parseXY(lines[i+1]) // 1 token
		prize := parseXY(lines[i+2])
		prize[0] += 10000000000000
		prize[1] += 10000000000000

		determinant := buttonA[0]*buttonB[1] - buttonA[1]*buttonB[0]

		pressesA := (prize[0]*buttonB[1] - prize[1]*buttonB[0]) / determinant
		pressesB := (buttonA[0]*prize[1] - buttonA[1]*prize[0]) / determinant

		if buttonA[0]*pressesA+buttonB[0]*pressesB == prize[0] && buttonA[1]*pressesA+buttonB[1]*pressesB == prize[1] {
			tokens += A_TOKEN*pressesA + B_TOKEN*pressesB
		}

	}

	return tokens
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/13.txt")
	t.Log(solution(string(input)))
}
