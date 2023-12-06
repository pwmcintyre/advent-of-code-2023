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
	want := "0"
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
	want := "0"
	run(source, target)
	if got := target.String(); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
