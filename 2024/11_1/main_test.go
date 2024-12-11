package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func solution(input string, blinks int) string {
	numbers := aoc.SplitAndTrim(input, " ")

	for i := 0; i < blinks; i++ {
		newNumbers := make([]string, 0, len(numbers)*2)

		for _, num := range numbers {
			if num == "0" {
				newNumbers = append(newNumbers, "1")
			} else if len(num)%2 == 0 {
				str := fmt.Sprint(num)
				leftStr, rightStr := str[:len(str)/2], str[len(str)/2:]
				leftNum, _ := strconv.Atoi(leftStr)
				rightNum, _ := strconv.Atoi(rightStr)
				newNumbers = append(newNumbers, fmt.Sprint(leftNum), fmt.Sprint(rightNum))
			} else {
				x, _ := strconv.Atoi(num)
				newNumbers = append(newNumbers, fmt.Sprint(x*2024))
			}
		}

		numbers = newNumbers
	}

	return strings.Join(numbers, " ")
}

func TestSolutionSample1(t *testing.T) {
	input := "0 1 10 99 999"

	expected := "1 2024 1 0 9 9 2021976"
	received := solution(input, 1)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "125 17"

	expected := "2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2"
	received := solution(input, 6)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := "17639"
	t.Log(len(aoc.SplitAndTrim(solution(input, 30), " ")))
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/11.txt")
	t.Log(len(aoc.SplitAndTrim(solution(string(input), 25), " ")))
}
