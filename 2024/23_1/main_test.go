package main

import (
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var regex = regexp.MustCompile(`t\w+`)

func solution(input string) int {
	lines := aoc.SplitAndTrim(input, "\n")

	adjacencyList := map[string]map[string]struct{}{}

	for _, line := range lines {
		vertices := strings.Split(line, "-")

		if _, exists := adjacencyList[vertices[0]]; !exists {
			adjacencyList[vertices[0]] = map[string]struct{}{}
		}

		if _, exists := adjacencyList[vertices[1]]; !exists {
			adjacencyList[vertices[1]] = map[string]struct{}{}
		}

		adjacencyList[vertices[0]][vertices[1]] = struct{}{}
		adjacencyList[vertices[1]][vertices[0]] = struct{}{}
	}

	sets := map[string]struct{}{}

	for vertexA := range adjacencyList {
		for vertexB := range adjacencyList[vertexA] {
			for vertexC := range adjacencyList[vertexB] {
				if _, exists := adjacencyList[vertexC][vertexA]; exists {
					set := []string{vertexA, vertexB, vertexC}
					slices.Sort(set)
					sets[strings.Join(set, ",")] = struct{}{}
				}
			}
		}
	}

	sum := 0

	for set := range sets {
		if regex.MatchString(set) {
			sum++
		}
	}

	return sum
}

func TestSolutionSample1(t *testing.T) {
	input :=
		`kh-tc
		qp-kh
		de-cg
		ka-co
		yn-aq
		qp-ub
		cg-tb
		vc-aq
		tb-ka
		wh-tc
		yn-cg
		kh-ub
		ta-co
		de-co
		tc-td
		tb-wq
		wh-td
		ta-ka
		td-qp
		aq-cg
		wq-ub
		ub-vc
		de-ta
		wq-aq
		wq-vc
		wh-yn
		ka-de
		kh-ta
		co-tc
		wh-qp
		tb-vc
		td-yn`

	expected := 7
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/23.txt")
	t.Log(solution(string(input)))
}
