package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func parseInput(input string) (map[byte][][2]int, [2]int) {
	lines := aoc.SplitAndTrim(input, "\n")

	antennas := make(map[byte][][2]int)

	re := regexp.MustCompile(`\w`)

	for row, line := range lines {
		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			col := match[0]
			antennas[line[col]] = append(antennas[line[col]], [2]int{row, col})
		}
	}

	return antennas, [2]int{len(lines[0]), len(lines)}
}

func isWithinMap(position [2]int, dimensions [2]int) bool {
	return position[0] >= 0 && position[0] < dimensions[0] && position[1] >= 0 && position[1] < dimensions[1]
}

func solution(input string) int {
	antennas, dimensions := parseInput(input)

	antinodes := make(map[string]struct{})

	for frequency := range antennas {
		for i := range antennas[frequency] {
			for j := range antennas[frequency] {
				if i == j {
					continue // skip itself
				}

				antenna1 := antennas[frequency][i]
				antenna2 := antennas[frequency][j]

				difference := [2]int{
					antenna2[0] - antenna1[0],
					antenna2[1] - antenna1[1],
				}

				antinode1 := [2]int{
					antenna1[0] - difference[0],
					antenna1[1] - difference[1],
				}

				antinode2 := [2]int{
					antenna2[0] + difference[0],
					antenna2[1] + difference[1],
				}

				if isWithinMap(antinode1, dimensions) {
					antinodes[fmt.Sprint(antinode1)] = struct{}{}
				}

				if isWithinMap(antinode2, dimensions) {
					antinodes[fmt.Sprint(antinode2)] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func TestSolutionSample(t *testing.T) {
	input :=
		`............
		 ........0..
		 .....0......
		 .......0....
		 ....0.......
		 ......A.....
		 ............
		 ............
		 ........A...
		 .........A..
		 ............
		 ............`

	expected := 14
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/08.txt")
	t.Log(solution(string(input)))
}
