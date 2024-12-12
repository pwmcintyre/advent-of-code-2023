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
	tests := []struct {
		input  string
		output int
	}{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 1},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 2},
		{input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", output: 0},
		{input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", output: 0},
		{input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", output: 5},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := part1(tt.input); got != tt.output {
				t.Errorf("part1() = %v, output %v", got, tt.output)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", output: 48},
		{input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", output: 12},
		{input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", output: 1560},
		{input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", output: 630},
		{input: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", output: 36},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := part2(tt.input); got != tt.output {
				t.Errorf("part2() = %v, output %v", got, tt.output)
			}
		})
	}
}
