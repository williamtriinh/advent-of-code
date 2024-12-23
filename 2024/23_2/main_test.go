package main

import (
	"maps"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/williamtriinh/advent-of-code/pkg/aoc"
)

var largestClique = map[string]struct{}{}

// Using Bron-Kerbosch without pivoting
func bronKerbosch(adjacencyList map[string]map[string]struct{}, R map[string]struct{}, P map[string]struct{}, X map[string]struct{}) {
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(largestClique) {
			largestClique = R
		}
		return
	}

	for vertex := range P {
		newR := maps.Clone(R)
		newR[vertex] = struct{}{}

		newP := map[string]struct{}{}
		for otherVertex := range P {
			if _, exists := adjacencyList[vertex][otherVertex]; exists {
				newP[otherVertex] = struct{}{}
			}
		}

		newX := map[string]struct{}{}
		for otherVertex := range X {
			if _, exists := adjacencyList[vertex][otherVertex]; exists {
				newX[otherVertex] = struct{}{}
			}
		}

		bronKerbosch(adjacencyList, newR, newP, newX)

		delete(P, vertex)
		X[vertex] = struct{}{}
	}
}

func solution(input string) string {
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

	allVertices := map[string]struct{}{}

	for vertex := range adjacencyList {
		allVertices[vertex] = struct{}{}
	}

	bronKerbosch(adjacencyList, map[string]struct{}{}, allVertices, map[string]struct{}{})
	set := slices.Collect(maps.Keys(largestClique))
	slices.Sort(set)

	return strings.Join(set, ",")
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

	expected := "co,de,ka,ta"
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/23.txt")
	t.Log(solution(string(input)))
}
