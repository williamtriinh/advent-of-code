package main

import (
	"os"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var directionToVec2 = map[byte]imath.Vec2{
	'^': aoc.Up,
	'>': aoc.Right,
	'v': aoc.Down,
	'<': aoc.Left,
}

func solution(input string) int {
	houses := map[imath.Vec2]int{}

	santaPosition := imath.Vec2{}
	robotPosition := imath.Vec2{}

	houses[santaPosition]++
	houses[robotPosition]++

	for i, char := range input {
		if i%2 == 0 {
			santaPosition = santaPosition.Add(directionToVec2[byte(char)])
			houses[santaPosition]++
		} else {
			robotPosition = robotPosition.Add(directionToVec2[byte(char)])
			houses[robotPosition]++
		}
	}

	return len(houses)
}

func TestSolutionSample1(t *testing.T) {
	input := "^v"

	expected := 3
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "^>v<"

	expected := 3
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := "^v^v^v^v^v"

	expected := 11
	received := solution(input)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/03.txt")
	t.Log(solution(string(input)))
}
