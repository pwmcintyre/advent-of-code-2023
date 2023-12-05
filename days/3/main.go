package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

var source io.Reader = os.Stdin

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	fmt.Fprintf(os.Stdout, "%d\n", part1(lines))
	fmt.Fprintf(os.Stdout, "%d\n", part2(lines))
}

var (
	white = color.New(color.FgWhite, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	red   = color.New(color.FgRed, color.Bold)
)

// setSymbolAt sets the sign at the given coordinates
// (checks for bounds)
func setSymbolAt(board [][]bool, x, y int) {
	if y < 0 || y >= len(board) {
		return
	}
	if x < 0 || x >= len(board[y]) {
		return
	}
	board[y][x] = true
}

// setSymbolAtNeighbours sets the sign at the given coordinates, and also all the spaces around it
func setSymbolAtNeighbours(board [][]bool, x, y int) {
	for y2 := -1; y2 < 2; y2++ {
		for x2 := -1; x2 < 2; x2++ {
			setSymbolAt(board, x+x2, y+y2)
		}
	}
}

func getSymbolNear(board [][]bool, x, y int) string {
	for y2 := -1; y2 < 2; y2++ {
		for x2 := -1; x2 < 2; x2++ {
			ly := y + y2
			lx := x + x2
			if ly >= 0 && ly < len(board) && lx >= 0 && lx < len(board[ly]) && board[ly][lx] {
				return fmt.Sprintf("%d,%d", x+x2, y+y2)
			}
		}
	}
	return ""
}

func part1(input []string) int {

	// allocate board
	// assumes width of input is constant
	var board [][]bool = make([][]bool, len(input))
	for y := range board {

		// NOTE: i'm appending a "." to the end of each line,
		// it makes things easier when digits abut the end of the line
		input[y] = input[y] + "."
		board[y] = make([]bool, len(input[y]))
	}

	// process symbols to create a 2D board of where the symbols are
	// ALSO: set true for spaces around symbols
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			// is digit
			if input[y][x] >= '0' && input[y][x] <= '9' {
				continue
			}

			// not a sign (".")
			if input[y][x] == '.' {
				continue
			}

			// assume sign
			setSymbolAt(board, x, y)
		}

	}

	// process numbers
	var sumP1 = 0
	for y := 0; y < len(input); y++ {
		var digits []int    // represents the digits of the part number as we're reading them
		var nearSymbol bool // represents if there is a sign near the number
		for x := 0; x < len(input[y]); x++ {

			// if it's a number; parse the digit; record if it's near a sign
			if input[y][x] >= '0' && input[y][x] <= '9' {
				digits = append(digits, int(input[y][x]-'0'))
				nearSymbol = nearSymbol || getSymbolNear(board, x, y) != ""
				continue
			}

			// else if it's not a number, and we have digits,
			// then we need to process the digits
			if len(digits) > 0 {

				// convert the digits to a number
				var number = 0
				for j := 0; j < len(digits); j++ {
					number = number*10 + digits[j]
				}

				// reset digits
				digits = []int{}

				// it needs to have been near a sign
				if nearSymbol {
					green.Fprintf(os.Stdout, "%d", number)
					sumP1 += number
				} else {
					red.Fprintf(os.Stdout, "%d", number)
				}

				// reset
				nearSymbol = false

			}

			white.Fprintf(os.Stdout, "%c", input[y][x])
		}
		fmt.Fprintf(os.Stdout, "\n")
	}

	return sumP1
}

func part2(input []string) int {

	// allocate board
	// assumes width of input is constant
	var board [][]bool = make([][]bool, len(input))
	for y := range board {

		// NOTE: i'm appending a "." to the end of each line,
		// it makes things easier when digits abut the end of the line
		input[y] = input[y] + "."
		board[y] = make([]bool, len(input[y]))
	}

	// find gears
	var gears map[string][]int = make(map[string][]int)
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if input[y][x] == '*' {

				// set board location
				setSymbolAt(board, x, y)

				// set gear array to store numbers
				gears[fmt.Sprintf("%d,%d", x, y)] = []int{}
			}
		}
	}

	// process numbers
	for y := 0; y < len(input); y++ {
		var digits []int    // represents the digits of the part number as we're reading them
		var nearGear string // represents if there is a sign near the number
		for x := 0; x < len(input[y]); x++ {

			// if it's a number; parse the digit; record if it's near a sign
			if input[y][x] >= '0' && input[y][x] <= '9' {
				digits = append(digits, int(input[y][x]-'0'))
				if nearGear == "" {
					nearGear = getSymbolNear(board, x, y)
				}
				continue
			}

			// else if it's not a number, and we have digits,
			// then we need to process the digits
			if len(digits) > 0 {

				// convert the digits to a number
				var number = 0
				for j := 0; j < len(digits); j++ {
					number = number*10 + digits[j]
				}

				// reset digits
				digits = []int{}

				// it needs to have been near a gear
				if nearGear != "" {
					green.Fprintf(os.Stdout, "%d", number)
					gears[nearGear] = append(gears[nearGear], number)
				} else {
					red.Fprintf(os.Stdout, "%d", number)
				}

				// reset
				nearGear = ""

			}

			white.Fprintf(os.Stdout, "%c", input[y][x])
		}
		fmt.Fprintf(os.Stdout, "\n")
	}

	// find the gears with exactly 2 numbers
	var sumP1 = 0
	for _, numbers := range gears {
		if len(numbers) == 2 {
			sumP1 += numbers[0] * numbers[1]
		}
	}

	return sumP1
}
