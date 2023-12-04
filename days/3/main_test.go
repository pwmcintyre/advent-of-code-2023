package main

import (
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		t.Errorf("main() = %v", err)
	}
	source = strings.NewReader(string(b))
	main()
}

func Test_part1(t *testing.T) {
	var input []string = []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	want := 4361
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

// some more examples from reddit:
// https://www.reddit.com/r/adventofcode/comments/189q9wv/2023_day_3_another_sample_grid_to_use/
func Test_part1b(t *testing.T) {
	var input []string = []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78..........",
		".......23...",
		"....90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1.......56",
	}
	want := 413
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part1c(t *testing.T) {
	var input []string = []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56",
	}
	want := 925
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
