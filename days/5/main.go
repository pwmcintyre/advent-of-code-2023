package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
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

type Mapping struct {
	Source string
	Dest   string
	Ranges []Range
}

type Range struct {
	SourceStart int
	DestStart   int
	Length      int
}

func part1(input []string) int {
	var err error

	// assume first line is seeds
	seeds, err := parseSeeds(input[0])
	if err != nil {
		panic(fmt.Errorf("error at line %s: %w", input[0], err))
	}
	fmt.Println(seeds)

	// parse mappings
	var source, dest string
	var inventory map[string]*Mapping = make(map[string]*Mapping)
	var mapping *Mapping
	for _, line := range input[1:] {
		if line == "" {
			continue
		}

		// parse map header
		if strings.HasSuffix(line, "map:") {
			source, dest, err = parseMapHeader(line)
			if err != nil {
				panic(fmt.Errorf("error at line %s: %w", line, err))
			}
			mapping = &Mapping{
				Source: source,
				Dest:   dest,
				Ranges: make([]Range, 0),
			}
			inventory[source] = mapping

		} else { // parse map section
			numbers, err := parseNumbers(line)
			if err != nil {
				panic(fmt.Errorf("error at line %s: %w", line, err))
			}
			if len(numbers) != 3 {
				panic(fmt.Sprintf("invalid line, need 3 numbers: %s", line))
			}
			inventory[source].Ranges = append(inventory[source].Ranges, Range{
				DestStart:   numbers[0],
				SourceStart: numbers[1],
				Length:      numbers[2],
			})
		}
	}

	// sort mappings
	for _, m := range inventory {
		sort.Slice(m.Ranges, func(i, j int) bool {
			return m.Ranges[i].SourceStart < m.Ranges[j].SourceStart
		})
	}

	// traverse graph; find smallest
	var smallest int = -1
	var ok bool
	for _, seed := range seeds {
		val := seed
		mapping = inventory["seed"]
		for {

			// find the range which meets the critera
			for _, r := range mapping.Ranges {
				if val >= r.SourceStart && val <= r.SourceStart+r.Length {
					val = r.DestStart + (val - r.SourceStart)
					break
				}
			}

			// move to the next category; until no more to move
			mapping, ok = inventory[mapping.Dest]
			if !ok {
				break
			}
		}

		// check if smallest
		if smallest == -1 || val < smallest {
			smallest = val
		}

	}

	return smallest
}

func parseSeeds(line string) ([]int, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid seed input: %s", line)
	}
	return parseNumbers(parts[1])
}

func parseMapHeader(line string) (string, string, error) {
	re := regexp.MustCompile(`^(\w+)-to-(\w+) map:$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid category line: %s", line)
	}
	return matches[1], matches[2], nil
}

func parseNumbers(line string) ([]int, error) {
	var numbers []int = make([]int, 0)
	for _, num := range strings.Split(line, " ") {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", num)
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
