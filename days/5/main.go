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

type Node struct {
	Name    string
	Next    *Node
	Mapping map[int]int
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
	var sourceStartRange, destinationStartRange, rangeLength int
	var inventory map[string]*Node = make(map[string]*Node)
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
			if _, ok := inventory[dest]; !ok {
				inventory[dest] = &Node{
					Name:    dest,
					Next:    nil,
					Mapping: make(map[int]int),
				}
			}
			if _, ok := inventory[source]; !ok {
				inventory[source] = &Node{
					Name:    source,
					Next:    inventory[dest],
					Mapping: make(map[int]int),
				}
			}
			inventory[source].Next = inventory[dest]
		} else { // parse map section
			numbers, err := parseNumbers(line)
			if err != nil {
				panic(fmt.Errorf("error at line %s: %w", line, err))
			}
			if len(numbers) != 3 {
				panic(fmt.Sprintf("invalid line, need 3 numbers: %s", line))
			}
			destinationStartRange = numbers[0]
			sourceStartRange = numbers[1]
			rangeLength = numbers[2]
			for i := 0; i < rangeLength; i++ {
				inventory[source].Mapping[sourceStartRange+i] = destinationStartRange + i
			}
		}
	}

	// traverse graph; find smallest
	var smallest int = -1
	var sourceID, destID int
	var ok bool
	for _, sID := range seeds {

		this := inventory["seed"]
		next := this.Next
		destID = sID
		for this.Next != nil {

			//  n = next
			next = this.Next
			sourceID = destID
			destID, ok = next.Mapping[sourceID]
			if !ok {
				destID = sourceID
			}
			this = this.Next

		}

		// end of graph; compare ID
		if smallest == -1 || destID < smallest {
			smallest = destID
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
