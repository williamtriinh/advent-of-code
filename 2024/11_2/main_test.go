package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func blink(value string, i int, cache map[string]int64) int64 {
	if i == 0 {
		return 1
	}

	key := fmt.Sprintf("%v:%d", value, i)

	if _, exists := cache[key]; exists {
		return cache[key]
	}

	if value == "0" {
		cache[key] = blink("1", i-1, cache)
	} else if len(value)%2 == 0 {
		left := value[:len(value)/2]
		right := value[len(value)/2:]

		// Remove zero padding
		leftNum, _ := strconv.Atoi(left)
		rightNum, _ := strconv.Atoi(right)

		left = fmt.Sprint(leftNum)
		right = fmt.Sprint(rightNum)

		cache[key] = blink(left, i-1, cache) + blink(right, i-1, cache)
	} else {
		x, _ := strconv.Atoi(value)
		cache[key] = blink(fmt.Sprint(x*2024), i-1, cache)
	}

	return cache[key]
}

func solution(input string, blinks int) int64 {
	numbers := aoc.SplitAndTrim(input, " ")
	cache := map[string]int64{}

	var sum int64 = 0
	for _, num := range numbers {
		sum += blink(num, blinks, cache)
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input := "0 1 10 99 999"

	var expected int64 = 7
	received := solution(input, 1)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := "125 17"

	var expected int64 = 22
	received := solution(input, 6)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution1(t *testing.T) {
	input, _ := os.ReadFile("../inputs/11.txt")

	var expected int64 = 203228
	received := solution(string(input), 25)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution2(t *testing.T) {
	input, _ := os.ReadFile("../inputs/11.txt")
	t.Log(solution(string(input), 75))
}
