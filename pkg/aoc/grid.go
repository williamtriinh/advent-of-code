package aoc

import (
	"fmt"
	"slices"
)

type Grid [][]byte

func (g Grid) Clone() Grid {
	clone := make(Grid, len(g))

	for row := range g {
		clone[row] = slices.Clone(g[row])
	}

	return clone
}

func (g Grid) Print() {
	for row := range g {
		for col := range g[row] {
			fmt.Printf("%c", g[row][col])
		}
		fmt.Println()
	}
}

func (g Grid) Width() int {
	return len(g[0])
}

func (g Grid) Height() int {
	return len(g)
}
