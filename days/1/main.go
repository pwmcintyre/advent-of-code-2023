package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		line  string
		read  int
		err   error
		sumP1 int
		sumP2 int
	)
	for {
		// read
		read, err = fmt.Scan(&line)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "read error: %v", err)
			return
		}

		// check
		if read == 0 {
			break
		}

		// process
		sumP1 += part1(line)
		sumP2 += part2(line)
	}
	fmt.Fprintln(os.Stdout, sumP1)
	fmt.Fprintln(os.Stdout, sumP2)
}

func part1(line string) int {

	var first, last int

	// check each letter if it is a number
	for i := 0; i < len(line); i++ {
		letter := line[i]
		if letter >= '0' && letter <= '9' {
			first = int(letter - '0')
			break
		}
	}

	// reverse the line
	for i := len(line) - 1; i >= 0; i-- {
		letter := line[i]
		if letter >= '0' && letter <= '9' {
			last = int(letter - '0')
			break
		}
	}

	return first*10 + last

}

var numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part2(line string) int {

	var found []int = []int{}

	// go through the line letter by letter
	// checking each time to see if a number is found
	// if found; add to the found list
	for line != "" {
		for key := range numbers {
			if strings.HasPrefix(line, key) {
				n := numbers[key]
				found = append(found, n)
				break
			}
		}
		line = line[1:]
	}

	// get the first and last
	var (
		first = found[0]
		last  = found[len(found)-1]
	)

	return first*10 + last

}
