package main

import (
	"os"
	"testing"
)

func parseInput(input string) []int {
	diskMap := make([]int, len(input))

	for i := range input {
		diskMap[i] = int(input[i] - 48)
	}

	return diskMap
}

func getDiskMapSize(diskMap []int) uint {
	var size uint = 0

	for _, blockSize := range diskMap {
		size += uint(blockSize)
	}

	return size
}

func getIndividualBlocks(diskMap []int, size uint) []int {
	individualBlocks := make([]int, size)

	pointer := 0

	for i, blockSize := range diskMap {
		for j := 0; j < blockSize; j++ {
			if i%2 == 0 { // file
				individualBlocks[pointer] = i / 2
			} else { // free space
				individualBlocks[pointer] = -1
			}
			pointer++
		}
	}

	return individualBlocks
}

func moveBlocks(diskMap []int, individualBlocks []int) {
	left := diskMap[0]
	right := len(individualBlocks) - 1

	for {
		for left < len(individualBlocks) && individualBlocks[left] != -1 {
			left++
		}

		for right >= 0 && individualBlocks[right] == -1 {
			right--
		}

		if left > right {
			break
		}

		individualBlocks[left] = individualBlocks[right]
		individualBlocks[right] = -1
		left++
		right--
	}
}

func getChecksum(individualBlocks []int) int64 {
	var sum int64 = 0

	for i, block := range individualBlocks {
		if block == -1 {
			break
		}
		sum += int64(i * int(block))
	}

	return sum
}

func solution(input string) int64 {
	diskMap := parseInput(input)
	diskMapSize := getDiskMapSize(diskMap)
	individualBlocks := getIndividualBlocks(diskMap, diskMapSize)
	moveBlocks(diskMap, individualBlocks)

	return getChecksum(individualBlocks)
}

func TestSolutionSample1(t *testing.T) {
	input := "2333133121414131402"

	var expected int64 = 1928
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "231"

	var expected int64 = 2
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample3(t *testing.T) {
	input := "2323"

	var expected int64 = 5
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSionSample4(t *testing.T) {
	input := "2020"

	var expected int64 = 5
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/09.txt")
	t.Log(solution(string(input)))
}
