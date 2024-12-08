package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
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

	antinodes := make(map[string]struct{}, imath.Pow(dimensions[0], 2))

	for frequency := range antennas {
		for i := range antennas[frequency] {
			for j := range antennas[frequency] {
				if i == j {
					continue // skip itself
				}

				node1 := antennas[frequency][i]
				node2 := antennas[frequency][j]

				// Add the antennas to the set of antinodes
				antinodes[fmt.Sprint(node1)] = struct{}{}
				antinodes[fmt.Sprint(node2)] = struct{}{}

				difference := [2]int{
					node2[0] - node1[0],
					node2[1] - node1[1],
				}

				otherNode1 := [2]int{
					node1[0] - difference[0],
					node1[1] - difference[1],
				}

				otherNode2 := [2]int{
					node2[0] + difference[0],
					node2[1] + difference[1],
				}

				for isWithinMap(otherNode1, dimensions) {
					antinodes[fmt.Sprint(otherNode1)] = struct{}{}
					node1 = otherNode1
					otherNode1 = [2]int{
						otherNode1[0] - difference[0],
						otherNode1[1] - difference[1],
					}
				}

				for isWithinMap(otherNode2, dimensions) {
					antinodes[fmt.Sprint(otherNode2)] = struct{}{}
					node2 = otherNode2
					otherNode2 = [2]int{
						otherNode2[0] + difference[0],
						otherNode2[1] + difference[1],
					}
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

	expected := 34
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/08.txt")
	t.Log(solution(string(input)))
}
