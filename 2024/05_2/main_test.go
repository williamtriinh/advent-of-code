package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

func parseSections(input string) ([]string, []string) {
	lines := aoc.SplitAndTrim(input, "\n")

	for i := len(lines) - 1; i > 0; i-- {
		if lines[i] == "" {
			return lines[:i], lines[i+1:]
		}
	}

	return []string{}, []string{}
}

func solution(input string) int {
	firstSection, secondSection := parseSections(input)

	adjacencyList := make(map[string][]string, len(firstSection))
	sum := 0

	for _, line := range firstSection {
		values := strings.Split(line, "|")
		adjacencyList[values[0]] = append(adjacencyList[values[0]], values[1])
	}

	for _, line := range secondSection {
		values := strings.Split(line, ",")

		correctOrder := true
		currentlyCorrect := false

		// Keep swapping until the pages are in the correct order
		for !currentlyCorrect {
			currentlyCorrect = true

			for i := 0; i < len(values)-1; i++ {
				if !slices.Contains(adjacencyList[values[i]], values[i+1]) {
					correctOrder = false
					currentlyCorrect = false

					// Swap adjacent values
					values[i], values[i+1] = values[i+1], values[i]
				}
			}
		}

		if !correctOrder {
			num, _ := strconv.Atoi(values[len(values)/2])
			sum += num
		}
	}

	return sum
}

func TestSolutionSample(t *testing.T) {
	input :=
		`47|53
		97|13
		97|61
		97|47
		75|29
		61|13
		75|53
		29|13
		97|29
		53|29
		61|53
		97|53
		61|29
		47|13
		75|47
		97|75
		47|61
		75|61
		47|29
		75|13
		53|13
		
		75,47,61,53,29
		97,61,53,29,13
		75,29,13
		75,97,47,61,53
		61,13,29
		97,13,75,29,47`

	expected := 123
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/05.txt")
	t.Log(solution(string(input)))
}
