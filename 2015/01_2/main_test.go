package main

import (
	"os"
	"testing"
)

func solution(input string) int {
	floor := 0

	for i, x := range input {
		if x == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}

func TestSolutionSample1(t *testing.T) {
	input := ")"

	expected := 1
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "()())"

	expected := 5
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/01.txt")
	t.Log(solution(string(input)))
}
