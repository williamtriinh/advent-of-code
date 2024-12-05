package main

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func smallestPerimeter(l, w, h int) int {
	sum := 0

	if l < w {
		sum += 2 * l
		if w < h {
			sum += 2 * w
		} else {
			sum += 2 * h
		}
	} else {
		sum += 2 * w
		if l < h {
			sum += 2 * l
		} else {
			sum += 2 * h
		}
	}

	return sum
}

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")

	sum := 0

	for _, line := range lines {
		values := strings.Split(line, "x")

		l, _ := strconv.Atoi(values[0])
		w, _ := strconv.Atoi(values[1])
		h, _ := strconv.Atoi(values[2])

		perimeter := smallestPerimeter(l, w, h)
		volume := l * w * h

		sum += perimeter + volume
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input := "2x3x4"

	expected := 34
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "1x1x10"

	expected := 14
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/02.txt")
	t.Log(solution(string(input)))
}
