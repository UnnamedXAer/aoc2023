package day21

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/unnamedxaer/aoc2023/help"
)

const inputNameSuffix = ""

// const inputNameSuffix = "_t"
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

func Ex1() {
	garden, startPos := extractData()

	// fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	total := processPart1(garden, startPos, 1)

	fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	fmt.Printf("\n\n  Total: %d", total)
}

const offsetsSize = 4

var offsets = [offsetsSize]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func processPart1(garden [][]puzzleElement, nextPos point, step int) int {
	const STEPS_TO_MAKE = 65
	// const STEPS_TO_MAKE = 6

	visits := make(map[point]int, STEPS_TO_MAKE)
	var t = time.Now().UnixMilli()
	var skipped = 0
	var move func(nextPos point, step int)

	var p point
	move = func(nextPos point, step int) {

		for _, o := range offsets {
			p.y = nextPos.y + o.y
			p.x = nextPos.x + o.x

			v := visits[p]
			if v > 0 && v <= step {
				skipped++
				if skipped%250_000_000*4 == 0 {
					end := time.Now().UnixMilli()
					fmt.Printf("\nskipped: %d, time: %d, visits size: %d", skipped, end-t, len(visits))
					t = end
				}
				continue
			}

			if !isAllowedPoint(garden, p) {
				continue
			}

			visits[p] = step

			if step == STEPS_TO_MAKE {
				continue
			}

			move(p, step+1)
		}
	}

	move(nextPos, step)

	total := 0
	fmt.Println()
	for p, v := range visits {
		if v%2 == 0 {
			garden[p.y][p.x] = 'O'
			total++
		}
	}

	return total
}

func isAllowedPoint(garden [][]puzzleElement, position point) bool {
	return position.y > -1 && position.x > -1 && position.y < len(garden) && position.x < len(garden[0]) && garden[position.y][position.x] != ROCK
}

// func isAllowed(garden [][]puzzleElement, y, x int) bool {
// 	if !(y > -1 && x > -1) {
// 		return false
// 	}

// 	if !(y < len(garden) && x < len(garden[y])) {
// 		return false
// 	}

// 	return garden[y][x] != ROCK
// }
