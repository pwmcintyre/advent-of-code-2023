package main

import (
	"fmt"
	"io"
	"os"
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
}

func part1(input []string) int {
	var sum int = 0
	return sum
}
