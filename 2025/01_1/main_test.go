package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

const MAX int = 99

func solution(input string) int {
	lines := strings.Split(input, "\n")

	pointing_at_zero_count := 0
	pointing_at := 50

	for _, line := range lines {
		line = strings.TrimSpace(line)
		direction := line[0]

		distance, _ := strconv.Atoi(line[1:])
		distance = distance % (MAX + 1)

		if direction == 'L' {
			pointing_at -= distance
			if pointing_at < 0 {
				pointing_at += MAX + 1
			}
		} else {
			pointing_at = (pointing_at + distance) % (MAX + 1)
		}

		if pointing_at == 0 {
			pointing_at_zero_count += 1
		}
	}

	return pointing_at_zero_count
}

func TestSolutionSample(t *testing.T) {
	input :=
		`L68
		L30
		R48
		L5
		R60
		L55
		L1
		L99
		R14
		L82`

	expected := 3
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../input.txt")
	t.Log(solution(string(input)))
}
