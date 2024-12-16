package main

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.sample.txt
var sample string

func Test_sample(t *testing.T) {
	source := strings.NewReader(sample)
	target := new(bytes.Buffer)
	// want := "2" // part 1
	want := "4"
	run(source, target)
	if got := target.String(); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

//go:embed input.txt
var actual string

func Test_actual(t *testing.T) {
	source := strings.NewReader(actual)
	target := new(bytes.Buffer)
	// want := "470" // too low
	// want := "488" // too low
	// want := "497" // too low
	// want := "504" // not right
	// want := "515" // not right!
	// want := "547" // not right!
	// want := "507" // not right!
	// want := "532" // not right!
	// want := "561" // not right!
	// want := "503" // found the bug 🔥😅 (and still not right)
	want := "503" //
	run(source, target)
	if got := target.String(); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
