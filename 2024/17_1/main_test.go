package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

func parseInput(input string) ([3]int, []int) {
	lines := aoc.SplitAndTrim(input, "\n")

	regex := regexp.MustCompile(`[\d,]+`)

	a, _ := strconv.Atoi(regex.FindAllString(lines[0], -1)[0])
	b, _ := strconv.Atoi(regex.FindAllString(lines[1], -1)[0])
	c, _ := strconv.Atoi(regex.FindAllString(lines[2], -1)[0])

	instructions := strings.Split(regex.FindAllString(lines[4], -1)[0], ",")
	program := make([]int, len(instructions))

	for i, instruction := range instructions {
		program[i], _ = strconv.Atoi(instruction)
	}

	return [3]int{a, b, c}, program
}

// Returns the combo operand
func combo(registers [3]int, operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return registers[0]
	} else if operand == 5 {
		return registers[1]
	} else {
		return registers[2]
	}
}

func solution(input string) (string, [3]int) {
	registers, program := parseInput(input)

	output := []string{}

	instructionPtr := 0

	for instructionPtr < len(program) {
		opcode := program[instructionPtr]
		operand := program[instructionPtr+1]

		switch opcode {
		case 0:
			registers[0] = registers[0] / imath.Pow(2, combo(registers, operand))
		case 1:
			registers[1] = registers[1] ^ operand
		case 2:
			registers[1] = combo(registers, operand) % 8
		case 3:
			if registers[0] == 0 {
				break
			}
			instructionPtr = operand
			continue
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			output = append(output, fmt.Sprintf("%d", combo(registers, operand)%8))
		case 6:
			registers[1] = registers[0] / imath.Pow(2, combo(registers, operand))
		case 7:
			registers[2] = registers[0] / imath.Pow(2, combo(registers, operand))
		}

		instructionPtr += 2
	}

	return strings.Join(output, ","), registers
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`Register A: 729
		 Register B: 0
		 Register C: 0

		 Program: 0,1,5,4,3,0`

	expected := "4,6,3,5,6,3,5,2,1,0"
	received, _ := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/17.txt")
	t.Log(solution(string(input)))
}
