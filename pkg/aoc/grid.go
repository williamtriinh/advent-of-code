package aoc

import (
	"fmt"
	"slices"

	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

type Grid [][]byte

type Cell struct {
	Position imath.Vec2
	Value    byte
}

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

func (g Grid) Cells() []Cell {
	cells := make([]Cell, g.Width()*g.Height())

	for i := 0; i < len(cells); i++ {
		position := imath.Vec2{X: i % g.Width(), Y: i / g.Height()}
		cells[i] = Cell{Position: position, Value: g[position.Y][position.X]}
	}

	return cells
}
