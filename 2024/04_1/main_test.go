package main

import (
	"os"
	"strings"
	"testing"
)

const WORD string = "XMAS"

func solution(input string) int {
	lines := strings.Split(input, "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	occurrences := 0

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != 'X' {
				continue
			}

			directions := [][]int{
				{1, 0},
				{1, -1},
				{0, -1},
				{-1, -1},
				{-1, 0},
				{-1, 1},
				{0, 1},
				{1, 1},
			}

			for _, direction := range directions {
				var j int

				// Begin at index 1 to search for "MAS"
				for j = 1; j < len(WORD); j++ {
					x := col + direction[0]*j
					y := row + direction[1]*j

					if x < 0 || x >= len(lines[row]) || y < 0 || y >= len(lines) {
						break
					}

					if lines[y][x] != WORD[j] {
						break
					}
				}

				if j == len(WORD) {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func TestSolutionSample(t *testing.T) {
	input :=
		`MMMSXXMASM
		MSAMXMSMSA
		AMXSXMAAMM
		MSAMASMSMX
		XMASAMXAMM
		XXAMMXXAMA
		SMSMSASXSS
		SAXAMASAAA
		MAMMMXMMMM
		MXMXAXMASX`

	expected := 18
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/04.txt")
	t.Log(solution(string(input)))
}
