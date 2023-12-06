package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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
	var inventory map[string]map[int]int = make(map[string]map[int]int)

	// assume first line is seeds
	seeds := parseSeeds(input[0])

	// parse mappings
	var source, dest string
	for _, line := range input[1:] {
		if line == "" {
			continue
		}
		if strings.Contains(line, "map") {
			source, dest = parseMapHeader(line)
		} else {
			numbers := parseNumbers(line)
			if len(numbers) != 3 {
				panic(fmt.Sprintf("invalid input: %s", line))
			}
			// ... TODO
		}
	}

	return sum
}

func parseSeeds(line string) []int {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid seed input: %s", line))
	}
	return parseNumbers(parts[1])
}

func parseMapHeader(line string) (string, string) {
	re := regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		panic(fmt.Sprintf("invalid category line: %s", line))
	}
	return matches[1], matches[2]
}

func parseLine(line string) (int, map[int]int, map[int]int) {
	re := regexp.MustCompile(`^Card +(\d+): (.*) \| (.*)$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 4 {
		panic(fmt.Sprintf("invalid input: %s", line))
	}
	cardID, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(fmt.Sprintf("invalid number: %s", matches[1]))
	}
	winningNumbers := parseNumbers(matches[2])
	myNumbers := parseNumbers(matches[3])
	return cardID, winningNumbers, myNumbers
}

func parseNumbers(line string) []int {
	var numbers []int = make([]int, 0)
	for _, num := range strings.Split(line, " ") {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(fmt.Sprintf("invalid number: %s", num))
		}
		numbers = append(numbers, n)
	}
	return numbers
}
