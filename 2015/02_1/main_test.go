package main

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")

	sum := 0

	for _, line := range lines {
		values := strings.Split(line, "x")

		l, _ := strconv.Atoi(values[0])
		w, _ := strconv.Atoi(values[1])
		h, _ := strconv.Atoi(values[2])

		a := l * w
		b := w * h
		c := h * l

		sum += 2*(a+b+c) + imath.Min(a, b, c)
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input := "2x3x4"

	expected := 58
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "1x1x10"

	expected := 43
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/02.txt")
	t.Log(solution(string(input)))
}
