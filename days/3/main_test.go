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
	var input [][]byte = [][]byte{
		[]byte(`467..114..`),
		[]byte(`...*......`),
		[]byte(`..35..633.`),
		[]byte(`......#...`),
		[]byte(`617*......`),
		[]byte(`.....+.58.`),
		[]byte(`..592.....`),
		[]byte(`......755.`),
		[]byte(`...$.*....`),
		[]byte(`.664.598..`),
	}
	want := 4361
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

// 273951 x
// 481451 x
// 483386 x
