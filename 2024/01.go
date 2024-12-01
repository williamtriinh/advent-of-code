package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputByte, _ := os.ReadFile("01.txt")

	inputString := string(inputByte)
	inputString = strings.ReplaceAll(inputString, "   ", "\n")

	locationIds := strings.Split(inputString, "\n")

	// partOne(locationIds)
	partTwo(locationIds)
}

func partOne(locationIds []string) {
	list1, list2 := []int{}, []int{}

	for i := 0; i < len(locationIds); i += 2 {
		num1, _ := strconv.Atoi(locationIds[i])
		num2, _ := strconv.Atoi(locationIds[i+1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]

		if diff < 0 {
			diff = diff * -1
		}

		sum += diff
	}

	fmt.Println(sum)
}

func partTwo(locationIds []string) {
	frequencies1 := make(map[int]int, len(locationIds)/2)
	frequencies2 := make(map[int]int, len(locationIds)/2)

	for i := 0; i < len(locationIds); i += 2 {
		left, _ := strconv.Atoi(locationIds[i])
		right, _ := strconv.Atoi(locationIds[i+1])

		if _, exists := frequencies1[left]; exists {
			frequencies1[left]++
		} else {
			frequencies1[left] = 1
		}

		if _, exists := frequencies2[right]; exists {
			frequencies2[right]++
		} else {
			frequencies2[right] = 1
		}
	}

	sum := 0

	for key, frequency1 := range frequencies1 {
		if frequency2, exists := frequencies2[key]; exists {
			sum += key * frequency2 * frequency1
		}
	}

	fmt.Println(sum)
}
