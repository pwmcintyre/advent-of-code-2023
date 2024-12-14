package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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

	// parse reports
	reports := make([][]int, len(input))
	for r, line := range input {

		// parse level (as numbers)
		parts := strings.Fields(line)
		reports[r] = make([]int, len(parts))
		for l, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				os.Exit(1)
			}
			reports[r][l] = num
		}
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
	for i := 0; i < len(report)-1; i++ {

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
