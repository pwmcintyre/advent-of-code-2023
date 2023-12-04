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
		string(`467..114..`),
		string(`...*......`),
		string(`..35..633.`),
		string(`......#...`),
		string(`617*......`),
		string(`.....+.58.`),
		string(`..592.....`),
		string(`......755.`),
		string(`...$.*....`),
		string(`.664.598..`),
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
		string("12.......*.."),
		string("+.........34"),
		string(".......-12.."),
		string("..78........"),
		string("..*....60..."),
		string("78.........."),
		string(".......23..."),
		string("....90*12..."),
		string("............"),
		string("2.2......12."),
		string(".*.........*"),
		string("1.1.......56"),
	}
	want := 413
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part1c(t *testing.T) {
	var input []string = []string{
		string("12.......*.."),
		string("+.........34"),
		string(".......-12.."),
		string("..78........"),
		string("..*....60..."),
		string("78.........9"),
		string(".5.....23..$"),
		string("8...90*12..."),
		string("............"),
		string("2.2......12."),
		string(".*.........*"),
		string("1.1..503+.56"),
	}
	want := 925
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

// 273951 x
// 481451 x
// 483386 x
// 483386 x
// 483386
// 509115
