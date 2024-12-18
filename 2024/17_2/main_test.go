package main

import (
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"github.com/williamtriinh/advent-of-code/pkg/imath"
)

var instructions []int

func parseInput(input string) []int {
	lines := aoc.SplitAndTrim(input, "\n")

	regex := regexp.MustCompile(`[\d,]+`)

	instructions := strings.Split(regex.FindAllString(lines[4], -1)[0], ",")
	program := make([]int, len(instructions))

	for i, instruction := range instructions {
		program[i], _ = strconv.Atoi(instruction)
	}

	return program
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

func program(a int) string {
	registers := [3]int{a, 0, 0}

	output := []string{}

	instructionPtr := 0

	for instructionPtr < len(instructions) {
		opcode := instructions[instructionPtr]
		operand := instructions[instructionPtr+1]

		switch opcode {
		case 0: // adv
			registers[0] /= imath.Pow(2, combo(registers, operand))
		case 1: // bxl
			registers[1] ^= operand
		case 2: // bst
			registers[1] = combo(registers, operand) % 8
		case 3: // jnz
			if registers[0] == 0 {
				break
			}
			instructionPtr = operand
			continue
		case 4: // bxc
			registers[1] ^= registers[2]
		case 5: // out
			output = append(output, fmt.Sprintf("%d", combo(registers, operand)%8))
		case 6: // bdv
			registers[1] = registers[0] / imath.Pow(2, combo(registers, operand))
		case 7: // cdv
			registers[2] = registers[0] / imath.Pow(2, combo(registers, operand))
		}

		instructionPtr += 2
	}

	return strings.Join(output, ",")
}

// Use BFS to find an input for register A. Start by finding the inputs that
// will give you the last value in the program instructions and add them to the
// queue.
func solution() int {
	queue := list.New()
	queue.PushBack(0)

	count := 1 // Keep track of which instruction we want to match with

	for queue.Len() > 0 && count <= len(instructions) {
		n := queue.Len()

		// Convert the instructions to a string of comma-separated values
		desired := strings.Trim(strings.Join(strings.Split(fmt.Sprint(instructions[len(instructions)-count:]), " "), ","), "[]")

		for i := 0; i < n; i++ {
			num := queue.Remove(queue.Front()).(int)

			for j := 0; j <= 7; j++ {
				variant := num<<3 + j // The higher-order bits are preserved in each iteration of 'count'
				output := program(variant)

				if output == desired {
					queue.PushBack(variant)
				}
			}
		}

		count++
	}

	return queue.Front().Value.(int)
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/17.txt")
	instructions = parseInput(string(input))
	fmt.Println(solution())
}
