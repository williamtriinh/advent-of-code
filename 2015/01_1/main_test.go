package main

import (
	"os"
	"testing"
)

func solution(input string) int {
	floor := 0

	for _, x := range input {
		if x == '(' {
			floor++
		} else {
			floor--
		}
	}

	return floor
}

func TestSolutionSample1(t *testing.T) {
	input := "(())"

	expected := 0
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "()()"

	expected := 0
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := "(()(()("

	expected := 3
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample4(t *testing.T) {
	input := ")())())"

	expected := -3
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/01.txt")
	t.Log(solution(string(input)))
}
