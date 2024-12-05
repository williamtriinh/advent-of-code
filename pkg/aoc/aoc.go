package aoc

import "strings"

func SplitAndTrim(input, separation string) []string {
	lines := strings.Split(input, separation)

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}
