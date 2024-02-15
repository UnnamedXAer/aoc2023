package day21

import (
	"bytes"
	"fmt"
	"time"
)

func Ex2() {
	garden, startPos := extractData()

	// fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	total := processPart2(garden, startPos, 1)

	fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	fmt.Printf("\n\n  Total: %d", total)

}

func processPart2(garden [][]puzzleElement, nextPos point, step int) int {
	// const STEPS_TO_MAKE = 26501365
	var STEPS_TO_MAKE = min(len(garden), len(garden[0])) / 2
	// var STEPS_TO_MAKE = 1266
	fmt.Printf("\n STEPS_TO_MAKE: %d", STEPS_TO_MAKE)

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

			if !isAllowedPointPart2(garden, p) {
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

	total := calcTotal(visits, garden)

	return total
}

func calcTotal(visits map[point]int, garden [][]byte) int {
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

func isAllowedPointPart2(garden [][]puzzleElement, position point) bool {
	ysize := len(garden)
	xsize := len(garden[0])

	y := position.y % ysize
	x := position.x % xsize

	if y < 0 {
		y = -y
	}
	if x < 0 {
		x = -x
	}

	return garden[y][x] != ROCK
}
