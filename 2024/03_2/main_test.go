package main

import (
	"os"
	"regexp"
	"strconv"
	"testing"
)

func solution(input string) int {
	regex := regexp.MustCompile(`(mul\((\d+),(\d+)\))|(don't\(\))|(do\(\))`)
	matches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	skipping := false

	for _, match := range matches {
		if match[0] == "don't()" {
			skipping = true
			continue
		}

		if match[0] == "do()" {
			skipping = false
			continue
		}

		if !skipping {
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			sum += x * y
		}
	}

	return sum
}

func TestSolutionSample(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	expected := 48
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/03.txt")
	t.Log(solution(string(input)))
}
