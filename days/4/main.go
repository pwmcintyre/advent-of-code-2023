package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var source io.Reader = os.Stdin

func main() {
	bytes, _ := io.ReadAll(source)
	lines := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	fmt.Fprintf(os.Stdout, "%d\n", part1(lines))
	fmt.Fprintf(os.Stdout, "%d\n", part2(lines))
}

func part1(input []string) int {
	var sum int = 0
	for _, line := range input {

		// parse
		_, winningNumbers, myNumbers := parseLine(line)

		// check
		var prize int = 0
		for _, num := range myNumbers {
			if _, ok := winningNumbers[num]; ok {
				if prize == 0 {
					prize = 1
				} else {
					prize *= 2
				}
			}
		}

		sum += prize

	}
	return sum
}

func part2(input []string) int {

	// allocate cards; each card exists at least once
	var cards map[int]int = make(map[int]int)
	for k := 1; k <= len(input); k++ {
		cards[k] = 1
	}

	// parse card stack
	for _, line := range input {

		// parse
		cardID, winningNumbers, myNumbers := parseLine(line)

		// check for winning numbers
		var wins int = 0
		for _, num := range myNumbers {
			if _, ok := winningNumbers[num]; ok {
				wins++
			}
		}

		// add a copy of the next cards FOR EACH copy of THIS card
		for i := cardID + 1; i <= cardID+wins; i++ {
			if _, ok := cards[i]; ok {
				cards[i] += cards[cardID]
			}
		}

	}

	// count up the cards
	var sum int = 0
	for _, count := range cards {
		sum += count
	}

	return sum
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

func parseNumbers(line string) map[int]int {
	var numbers map[int]int = make(map[int]int)
	for _, num := range strings.Split(line, " ") {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(fmt.Sprintf("invalid number: %s", num))
		}
		numbers[n] = n
	}

	return numbers
}
