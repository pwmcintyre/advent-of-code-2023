package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var source io.Reader = os.Stdin

func main() {

	// read whole of input
	var line []byte
	var input [][]byte = make([][]byte, 0)
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		line = scanner.Bytes()
		if len(line) == 0 {
			break
		}
		input = append(input, line)
	}

	fmt.Fprintf(os.Stdout, "%d\n", part1(input))
}

// returns true if there is a sign at the given coordinates
func SignAt(board [][]bool, x, y int) bool {
	if y < 0 || y >= len(board) {
		return false
	}
	if x < 0 || x >= len(board[y]) {
		return false
	}
	return board[y][x]
}

// returns true if there is a sign above, below or at the given coordinates
func SignVertical(board [][]bool, x, y int) bool {
	return SignAt(board, x, y-1) || SignAt(board, x, y) || SignAt(board, x, y+1)
}

var (
	white = color.New(color.FgWhite, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	blue  = color.New(color.FgBlue, color.Bold)
	red   = color.New(color.FgRed, color.Bold)
)

func part1(input [][]byte) int {

	// process signs to create a 2D board of where the signs are
	var board [][]bool = make([][]bool, 0)
	for y := 0; y < len(input); y++ {
		board = append(board, make([]bool, len(input[y])))
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
			board[y][x] = true
		}

	}

	// display
	for y := 0; true && y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			// is digit
			if input[y][x] >= '0' && input[y][x] <= '9' {
				switch {
				case SignVertical(board, x-1, y):
					blue.Printf("%c", input[y][x])
				case SignVertical(board, x, y):
					green.Printf("%c", input[y][x])
				case SignVertical(board, x+1, y):
					blue.Printf("%c", input[y][x])
				default:
					white.Printf("%c", input[y][x])
				}
				continue
			}

			// not a sign (".")
			if input[y][x] == '.' {
				fmt.Printf(" ")
				continue
			}

			// assume sign
			fmt.Printf("×")
		}
		fmt.Printf("\n")

	}

	// process numbers
	var sumP1 = 0
	for y := 0; y < len(input); y++ {
		var digits []int  // represents the digits of the part number as we're reading them
		var nearSign bool // represents if there is a sign near the number
		for x := 0; x < len(input[y]); x++ {

			// if it's a number; parse the digit; record if it's near a sign
			if input[y][x] >= '0' && input[y][x] <= '9' {
				digits = append(digits, int(input[y][x]-'0'))
				nearSign = nearSign ||
					SignVertical(board, x+1, y) ||
					SignVertical(board, x, y) ||
					SignVertical(board, x-1, y)

				if x != len(input[y])-1 {
					continue
				}
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
				if nearSign {
					green.Printf("%d", number)
					sumP1 += number
				} else {
					red.Printf("%d", number)
				}
				nearSign = false

			}

			// not a sign (".")
			if input[y][x] == '.' {
				fmt.Printf(" ")
			} else if !(input[y][x] >= '0' && input[y][x] <= '9') {
				fmt.Printf("×")
			}

		}
		fmt.Printf("\n")
	}

	return sumP1
}

func part2(line string) int {
	return 0
}
