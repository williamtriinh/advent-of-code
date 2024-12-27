package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func parseInput(input string) [][]string {
	lines := aoc.SplitAndTrim(input, "\n")
	instructions := make([][]string, len(lines))

	for i, line := range lines {
		instructions[i] = []string{line[:3]}
		instructions[i] = slices.Concat(instructions[i], strings.Split(line[4:], ", "))
	}

	return instructions
}

func parseOffset(offset string) int {
	amount, _ := strconv.Atoi(offset[1:])

	if offset[0] == '+' {
		return amount
	}

	return -amount
}

func solution(input string) map[string]int {
	instructions := parseInput(input)

	registers := map[string]int{}
	pointer := 0

	for pointer < len(instructions) {
		instruction := instructions[pointer]

		switch instruction[0] {
		case "hlf":
			registers[instruction[1]] /= 2
		case "tpl":
			registers[instruction[1]] *= 3
		case "inc":
			registers[instruction[1]]++
		case "jmp":
			pointer += parseOffset(instruction[1])
			continue
		case "jie":
			if registers[instruction[1]]%2 == 0 {
				pointer += parseOffset(instruction[2])
				continue
			}
		case "jio":
			if registers[instruction[1]] == 1 {
				pointer += parseOffset(instruction[2])
				continue
			}
		}

		pointer++
	}

	return registers
}

func TestSolutionSample(t *testing.T) {
	input :=
		`inc a
		jio a, +2
		tpl a
		inc a`

	expected := 2
	received := solution(input)

	if expected != received["a"] {
		t.Errorf("Expected %v but received %v", expected, received["a"])
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/23.txt")
	t.Log(solution(string(input))["b"])
}
