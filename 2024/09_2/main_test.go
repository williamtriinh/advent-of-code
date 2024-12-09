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

func moveBlocks(individualBlocks []int) {
	right := len(individualBlocks) - 1

	for right > 0 {
		for i := right; i > 0; i++ {
			if individualBlocks[right] != -1 {
				break
			}
			right--
		}

		fileSize := 0
		for i := right; i > 0; i-- {
			if individualBlocks[i] != individualBlocks[right] {
				break
			}
			fileSize++
		}

		for left := 0; left < right; {
			if individualBlocks[left] != -1 {
				left++
				continue
			}

			freeSpace := 0
			for i := left; i < right; i++ {
				if individualBlocks[i] != individualBlocks[left] {
					break
				}
				freeSpace++
			}

			if fileSize <= freeSpace {
				for i := 0; i < fileSize; i++ {
					individualBlocks[left+i] = individualBlocks[right+i+1-fileSize]
					individualBlocks[right+i+1-fileSize] = -1
				}
			}

			left += freeSpace
		}

		right -= fileSize
	}
}

func getChecksum(individualBlocks []int) int64 {
	var sum int64 = 0

	for i, block := range individualBlocks {
		if block != -1 {
			sum += int64(i * int(block))
		}
	}

	return sum
}

func solution(input string) int64 {
	diskMap := parseInput(input)
	diskMapSize := getDiskMapSize(diskMap)
	individualBlocks := getIndividualBlocks(diskMap, diskMapSize)
	moveBlocks(individualBlocks)

	return getChecksum(individualBlocks)
}

func TestSolutionSample1(t *testing.T) {
	input := "2333133121414131402"

	var expected int64 = 2858
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolutionSample2(t *testing.T) {
	input := "1313165"

	var expected int64 = 169
	received := solution(input)

	if received != expected {
		t.Errorf("Received %v but expected %v", received, expected)
	}
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/09.txt")
	t.Log(solution(string(input)))
}
