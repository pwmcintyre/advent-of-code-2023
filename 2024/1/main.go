package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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
	// parse parts as integers
	var left []int = make([]int, len(input))
	var right []int = make([]int, len(input))
	for i, line := range input {
		fmt.Sscanf(line, "%d %d", &left[i], &right[i])
	}

	// sort left and right
	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	// process
	for i := 0; i < len(left); i++ {
		if right[i] < left[i] {
			result += left[i] - right[i]
		} else {
			result += right[i] - left[i]
		}
	}

	return
}

func part2(input []string) (result int) {
	// parse parts as integers
	var left []int = make([]int, len(input))
	var right []int = make([]int, len(input))
	for i, line := range input {
		fmt.Sscanf(line, "%d %d", &left[i], &right[i])
	}

	// count occurances of each number in right
	var counts map[int]int = make(map[int]int)
	for _, r := range right {
		counts[r]++
	}

	// process left
	for _, r := range left {
		result += r * counts[r]
	}

	return
}
