package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	run(os.Stdin, os.Stdout)
	fmt.Fprintf(os.Stdout, "\n")
}

func run(source io.Reader, target io.Writer) {
	bytes, _ := io.ReadAll(source)
	lines := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	fmt.Fprint(target, part1(lines))
	// fmt.Fprint(target, part2(lines))
}

func part1(input []string) (result int) {

	// parse levels
	var reports [][]int = make([][]int, len(input))
	for l, line := range input {
		reports[l] = make([]int, 5)
		fmt.Sscanf(line, "%d %d %d %d %d", &reports[l][0], &reports[l][1], &reports[l][2], &reports[l][3], &reports[l][4])
	}

	// valid reports
	for _, report := range reports {
		if isSafe(report) {
			result++
		}
	}

	return
}

func isSafe(report []int) bool {
	var increasing = report[0] < report[1]
	for i := 0; i < 4; i++ {

		// The levels are either all increasing or all decreasing.
		if increasing != (report[i] < report[i+1]) {
			return false
		}

		// Any two adjacent levels differ by at least one and at most three.
		var delta = report[i] - report[i+1]
		if delta < 0 {
			delta = -delta
		}
		if delta < 1 || delta > 3 {
			return false
		}
	}
	return true
}

func part2(input []string) (result int) {
	return
}
