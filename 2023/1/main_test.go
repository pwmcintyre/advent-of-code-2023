package main

import "testing"

func Test_part1(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{input: "1abc2", output: 12},
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
		{input: "two1nine", output: 29},
		{input: "eightwothree", output: 83},
		{input: "abcone2threexyz", output: 13},
		{input: "xtwone3four", output: 24},
		{input: "4nineeightseven2", output: 42},
		{input: "zoneight234", output: 14},
		{input: "7pqrstsixteen", output: 76},
		{input: "oneight", output: 18},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := part2(tt.input); got != tt.output {
				t.Errorf("part2() = %v, output %v", got, tt.output)
			}
		})
	}
}
