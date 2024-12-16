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
	// fmt.Fprint(target, part1(lines))
	fmt.Fprint(target, part2(lines))
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
		if isSafe(report, 0) {
			result++
		}
	}

	return
}

func isSafe(report []int, tolerance int) bool {

	var errors = 0
	var increasing = report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		var delta = report[i+1] - report[i]

		if !check(increasing, delta) {
			errors++

			if errors > tolerance {
				return false
			}

			// attempt to remove a level
			var t = tolerance - 1
			for j := 0; j < len(report)-1; j++ {
				r := report[:j:j]
				r = append(r, report[j+1:]...)
				if isSafe(r, t) {
					return true
				}
			}
			return false

		}
	}

	return true
}

func check(increasing bool, delta int) bool {

	// The levels are either all increasing or all decreasing.
	if increasing == (delta < 0) {
		return false
	}

	// Any two adjacent levels differ by at least one and at most three.
	if delta < -3 || delta > 3 {
		return false
	}
	if delta > -1 && delta < 1 {
		return false
	}

	return true
}

func part2(input []string) (result int) {

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
		if isSafe(report, 1) {
			result++
		}
	}

	return
}
