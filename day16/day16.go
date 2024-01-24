package day16

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]byte {

	f, err := os.Open("./day16/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	tiles := make([][]byte, 0, 110)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		tiles = append(tiles, line)
	}

	help.IfErr(scanner.Err())

	return tiles
}

func Ex1() {
	total := 0
	tiles := extractData()
	size := len(tiles)
	visits := make([]direction, size*size)

	traverse(tiles, size, visits, 0, 0, east)

	var zeroDirection direction
	for i := 0; i < size; i++ {
		// fmt.Printf("\n")
		for k := 0; k < size; k++ {
			v := visits[i*size+k]
			// fmt.Print(v)
			if v != zeroDirection {
				total++
			}
		}
	}

	fmt.Printf("\n\n Total: %d", total)
}

type direction byte

const north direction = 'N'
const south direction = 'S'
const west direction = 'W'
const east direction = 'E'
const endOfPath direction = 'X'

func traverse(tiles [][]byte, size int, visits []direction, row, col int, direct direction) {
	var col1, row1 int
	var direct1 direction

	lastDirection := endOfPath
	for row != -1 {
		lastDirection = visits[row*size+col]
		if lastDirection == direct {
			return
		}

		visits[row*size+col] = direct

		// fmt.Printf("\n [%2d %2d]", row, col)
		// for i := 0; i < size; i++ {
		// fmt.Printf("\n")
		// for k := 0; k < size; k++ {
		// v := visits[i*size+k]
		// fmt.Print(v)
		// }
		// }

		tile := tiles[row][col]
		switch tile {
		case '.':
			row, col = handleEmptyTile(size, row, col, direct)

		case '\\':
			row, col, direct = handleMirrorBackslash(size, row, col, direct)
		case '/':
			row, col, direct = handleMirrorSlash(size, row, col, direct)
		case '|':
			row, col, direct, row1, col1, direct1 = handleSplitPipe(size, row, col, direct)
			traverse(tiles, size, visits, row1, col1, direct1)
		case '-':
			row, col, direct, row1, col1, direct1 = handleSplitDash(size, row, col, direct)
			traverse(tiles, size, visits, row1, col1, direct1)
		default:
			panic(fmt.Sprintf("\n unknown tile: %v", string(tile)))
		}
	}
}

// handles '|'
func handleSplitPipe(size, row, col int, direct direction) (int, int, direction, int, int, direction) {
	direct1 := south
	col1, row1 := col, row

	switch direct {
	case north:
		row--
		row1, col1 = -1, -1
	case south:
		row++
		row1, col1 = -1, -1
	case west:
		row--
		row1++
		direct = north
	case east:
		row--
		row1++
		direct = north
	}

	if outsizeRange(size, row, col) {
		col, row = -1, -1
		direct = endOfPath
	}

	if outsizeRange(size, row1, col1) {
		col1, row1 = -1, -1
		direct1 = endOfPath
	}

	return row, col, direct, row1, col1, direct1
}

// handles '-'
func handleSplitDash(size, row, col int, direct direction) (int, int, direction, int, int, direction) {

	direct1 := east
	col1, row1 := col, row

	switch direct {
	case north:
		col--
		col1++
		direct = west
	case south:
		col--
		col1++
		direct = west
	case west:
		col--
		row1, col1 = -1, -1
	case east:
		col++
		row1, col1 = -1, -1
	}

	if outsizeRange(size, row, col) {
		col, row = -1, -1
		direct = endOfPath
	}

	if outsizeRange(size, row1, col1) {
		col1, row1 = -1, -1
		direct1 = endOfPath
	}

	return row, col, direct, row1, col1, direct1
}

// handles '\'
func handleMirrorBackslash(size, row, col int, direct direction) (int, int, direction) {

	switch direct {
	case north:
		col--
		direct = west
	case south:
		col++
		direct = east
	case west:
		row--
		direct = north
	case east:
		row++
		direct = south
	}

	if outsizeRange(size, row, col) {
		return -1, -1, endOfPath
	}

	return row, col, direct
}

// handles '/'
func handleMirrorSlash(size, row, col int, direct direction) (int, int, direction) {

	switch direct {
	case north:
		col++
		direct = east
	case south:
		col--
		direct = west
	case west:
		row++
		direct = south
	case east:
		row--
		direct = north
	}

	if outsizeRange(size, row, col) {
		return -1, -1, endOfPath
	}

	return row, col, direct
}

func handleEmptyTile(size, row, col int, direct direction) (int, int) {

	switch direct {
	case north:
		row--
	case south:
		row++
	case west:
		col--
	case east:
		col++
	}

	if outsizeRange(size, row, col) {
		return -1, -1
	}

	return row, col
}

func outsizeRange(size, row, col int) bool {

	if row < 0 || col < 0 || row == size || col == size {
		return true
	}

	return false
}

func Ex2() {
	total := 0
	tiles := extractData()
	size := len(tiles)
	var zeroDirection direction

	direct := west
	col := 0
	row := 0
	for i := 0; i < 2; i++ {
		for row = 0; row < size; row++ {
			visits := make([]direction, size*size)
			traverse(tiles, size, visits, row, col, direct)
			tmpTotal := 0
			for i := 0; i < size; i++ {
				// fmt.Printf("\n")
				for k := 0; k < size; k++ {
					v := visits[i*size+k]
					// fmt.Print(v)
					if v != zeroDirection {
						tmpTotal++
					}
				}
			}

			if tmpTotal > total {
				total = tmpTotal
			}
		}
		col = size - 1
		direct = west
	}

	direct = south
	row = 0
	for i := 0; i < 2; i++ {
		for col = 0; col < size; col++ {
			visits := make([]direction, size*size)
			traverse(tiles, size, visits, row, col, direct)
			tmpTotal := 0
			for i := 0; i < size; i++ {
				// fmt.Printf("\n")
				for k := 0; k < size; k++ {
					v := visits[i*size+k]
					// fmt.Print(v)
					if v != zeroDirection {
						tmpTotal++
					}
				}
			}

			if tmpTotal > total {
				total = tmpTotal
			}
		}
		row = size - 1
		direct = north
	}

	fmt.Printf("\n\n Total: %d", total)

}
