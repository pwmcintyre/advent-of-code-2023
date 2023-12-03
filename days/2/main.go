package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

var source io.Reader = os.Stdin

func main() {
	var (
		line  string
		sumP1 int
		sumP2 int
	)

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {

		// read
		line = scanner.Text()

		// check
		if line == "" {
			break
		}

		// process
		sumP1 += part1(line)
		sumP2 += part2(line)
	}
	fmt.Fprintln(os.Stdout, sumP1)
	fmt.Fprintln(os.Stdout, sumP2)
}

type game struct {
	ID      int
	Reveals []Reveal
	Max     RGB
}

type Reveal = RGB

type RGB struct {
	R int
	G int
	B int
}

func parse(line string) game {
	var (
		g game = game{
			Reveals: []Reveal{},
		}
		reveal string
		count  int
		colour string
	)

	// parse id
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		panic(fmt.Errorf("invalid input: %s", line))
	}
	fmt.Sscanf(parts[0], "Game %d", &g.ID)

	// parse reveals
	reveals := strings.Split(parts[1], ";")
	for _, reveal = range reveals {

		// parse reveal into its colour parts
		r := Reveal{}
		colourGroup := strings.Split(reveal, ",")
		for _, cg := range colourGroup {
			fmt.Sscanf(strings.Trim(cg, " "), "%d %s", &count, &colour)
			switch colour {
			case "red":
				r.R = count
				if count > g.Max.R {
					g.Max.R = count
				}
			case "green":
				r.G = count
				if count > g.Max.G {
					g.Max.G = count
				}
			case "blue":
				r.B = count
				if count > g.Max.B {
					g.Max.B = count
				}
			}
		}

		g.Reveals = append(g.Reveals, r)
	}
	return g
}

func part1(line string) (val int) {
	game := parse(line)
	for _, reveal := range game.Reveals {
		if reveal.R > maxRed || reveal.B > maxBlue || reveal.G > maxGreen {
			return 0
		}
	}
	return game.ID
}

func part2(line string) (val int) {
	game := parse(line)
	return game.Max.R * game.Max.G * game.Max.B
}
