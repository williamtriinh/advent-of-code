package main

import (
	"testing"
)

func solution(row, col int) int64 {
	var code int64 = 20151125

	currentRow := 1
	currentCol := 1

	for !(currentRow == row && currentCol == col) {
		if currentRow <= 1 {
			currentRow = currentCol + 1
			currentCol = 1
		} else {
			currentRow--
			currentCol++
		}

		code = (code * 252533) % 33554393
	}

	return code
}

func TestSolutionSample1(t *testing.T) {
	var expected int64 = 31916031
	received := solution(2, 1)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolutionSample2(t *testing.T) {
	var expected int64 = 27995004
	received := solution(6, 6)

	if expected != received {
		t.Errorf("Expected %v but received %v", expected, received)
	}
}

func TestSolution(t *testing.T) {
	t.Log(solution(2981, 3075))
}
