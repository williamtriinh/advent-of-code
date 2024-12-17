package aoc

import (
	"fmt"
	"slices"
	"strings"
)

func PrintGrid(grid [][]byte) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Printf("%c", grid[row][col])
		}
		fmt.Println()
	}
}

func CloneGrid(grid [][]byte) [][]byte {
	copy := make([][]byte, len(grid))

	for row := range grid {
		copy[row] = slices.Clone(grid[row])
	}

	return copy
}

func SplitAndTrim(input, separation string) []string {
	lines := strings.Split(input, separation)

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}
