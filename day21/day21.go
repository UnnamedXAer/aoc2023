package day21

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
)

// const inputNameSuffix= ""
const inputNameSuffix = "_t"
const inputName = "./day21/data" + inputNameSuffix + ".txt"

type puzzleElement = byte

const (
	START   puzzleElement = 'S'
	PLOT    puzzleElement = '.'
	ROCK    puzzleElement = '#'
	IM_HERE puzzleElement = 'O'
)

type point struct {
	y, x int
}

func extractData() ([][]puzzleElement, point) {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	garden := make([][]puzzleElement, 0, 135)
	startPos := point{-1, -1}
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, []puzzleElement(line))
		if startPos.y == -1 {
			idx := strings.IndexByte(line, byte(START))
			if idx > -1 {
				startPos.y = len(garden) - 1
				startPos.x = idx
			}
		}
	}

	help.IfErr(scanner.Err())

	return garden, startPos
}

// const STEPS_TO_MAKE = 64
const STEPS_TO_MAKE = 6

func Ex1() {
	garden, startPos := extractData()

	total := 0
	// y := startPos.y + STEPS_TO_MAKE
	// xsize := -1
	// for ; y >= 0; y-- {
	// 	if y < startPos.y {
	// 		xsize -= 1
	// 	} else {
	// 		xsize += 1
	// 	}

	// 	if y < 0 {
	// 		break
	// 	}

	// 	if y >= len(garden) {
	// 		continue
	// 	}

	// 	x := max(0, startPos.x-xsize)
	// 	xend := min(len(garden[y])-1, startPos.x+xsize) + 1

	// 	for ; x < xend; x++ {
	// 		if isAllowed(garden, y, x) {
	// 			total++
	// 		} else {
	// 			// fmt.Printf("\n y %d, x %d, xend %d, startpos: %+v", y, x, xend, startPos)
	// 		}
	// 	}

	// }

	fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	visits := make(map[point]int, STEPS_TO_MAKE*STEPS_TO_MAKE*1.5)
	endPositions := move(garden, startPos, 1, &visits)
	total = len(endPositions)

	// total = 0
	for p, _ := range checked {
		// if v%2 == 0 {
		garden[p.y][p.x] = 'O'
		// total++
		// }
	}
	fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	fmt.Printf("\n\n  Total: %d", total)
}

var offsets = [4]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

// type direction int

// const (
// 	north direction = iota
// 	south
// 	west
// 	east
// )

func (p point) move(offset point) point {
	p.y += offset.y
	p.x += offset.x
	return p
}

var checked = make(map[point]int)

func move(garden [][]puzzleElement, nextPos point, step int, visits *map[point]int) []point {
	out := make([]point, 0, 4)

	for _, o := range offsets {
		p := nextPos.move(o)
		v := (*visits)[p]
		if v > 0 {
			continue
		}

		if !isAllowedPoint(garden, p) {
			continue
		}

		checked[p] = step
		// fmt.Printf("checking: %v", p)
		if step%2 == 0 {
			(*visits)[p] = step
		}

		if step == STEPS_TO_MAKE {
			out = appendPoint(out, p)
			continue
		}

		for _, p := range move(garden, p, step+1, visits) {
			out = appendPoint(out, p)
		}
	}

	return out
}

func appendPoint(points []point, p point) []point {
	if len(points) > 0 {
		for _, p2 := range points {
			if p2.y == p.y && p2.x == p.x {
				return points
			}
		}
	}

	return append(points, p)
}

func isAllowedPoint(garden [][]puzzleElement, position point) bool {
	return position.y > -1 && position.x > -1 && position.y < len(garden) && position.x < len(garden[0]) && garden[position.y][position.x] != ROCK
}

func isAllowed(garden [][]puzzleElement, y, x int) bool {
	if !(y > -1 && x > -1) {
		return false
	}

	if !(y < len(garden) && x < len(garden[y])) {
		return false
	}

	return garden[y][x] != ROCK
}
