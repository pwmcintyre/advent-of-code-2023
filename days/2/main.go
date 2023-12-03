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
	maxBlue  = 13
	maxGreen = 14
)

var max = maxRed + maxBlue + maxGreen

var source io.Reader = os.Stdin

func main() {
	var (
		line  string
		sumP1 int
		// sumP2  int
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
	}
	fmt.Fprintln(os.Stdout, sumP1)
}

type game struct {
	ID      int
	Reveals []Reveal
}

type Reveal struct {
	RGB RGB
	Sum int
}

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
				r.RGB.R = count
			case "green":
				r.RGB.G = count
			case "blue":
				r.RGB.B = count
			}
			r.Sum += count
		}

		g.Reveals = append(g.Reveals, r)
	}
	return g
}

func part1(line string) (val int) {

	game := parse(line)
	// defer func() { fmt.Printf("game %v %d\n", game, val) }()
	fmt.Printf("%v\n", game)

	// check if game is valid
	for _, reveal := range game.Reveals {
		if reveal.RGB.R > maxRed {
			fmt.Println("🔥 red too high")
		}
		if reveal.RGB.G > maxGreen {
			fmt.Println("🔥 green too high")
		}
		if reveal.RGB.B > maxBlue {
			fmt.Println("🔥 blue too high")
		}
		if reveal.RGB.R > maxRed || reveal.RGB.B > maxBlue || reveal.RGB.G > maxGreen {
			return 0
		}
	}

	fmt.Println("✅ valid")
	return game.ID

}
