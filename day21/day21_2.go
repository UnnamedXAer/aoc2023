package day21

import (
	"bytes"
	"fmt"
	"time"
)

func Ex2() {
	garden, startPos := extractData()

	expansionFactor := 4
	expandedGarden := make([][]byte, len(garden)*expansionFactor)

	for k, line := range garden {
		expandedLine := make([]byte, len(garden)*expansionFactor)
		for i := 0; i < expansionFactor; i++ {
			copy(expandedLine[i*len(garden):], line)
		}
		for i := 0; i < expansionFactor; i++ {
			expandedGarden[k+(i*len(garden))] = expandedLine
		}
	}
	fmt.Printf("\n%v", string(bytes.Join([][]byte(expandedGarden), []byte{'\n'})))

	total := processPart2(expandedGarden, startPos, 1)

	fmt.Printf("\n%v", string(bytes.Join([][]byte(garden), []byte{'\n'})))

	fmt.Printf("\n\n  Total: %d", total)

}

func processPart2(garden [][]puzzleElement, nextPos point, step int) int {
	// const STEPS_TO_MAKE = 26501365
	// var STEPS_TO_MAKE = 1266
	// var STEPS_TO_MAKE = 10
	var STEPS_TO_MAKE = 20
	fmt.Printf("\n STEPS_TO_MAKE: %d", STEPS_TO_MAKE)

	// const STEPS_TO_MAKE = 6

	visits := make(map[point]int, STEPS_TO_MAKE)
	var t = time.Now().UnixMilli()
	var skipped = 0
	var move func(nextPos point, step int)
	size := len(garden)
	possibleEndTilesCnt := 0
	calculated := 0

	runs := -1

	var p point
	move = func(nextPos point, step int) {

		for _, o := range offsets {
			runs++
			if runs%1_000_000_000 == 0 {
				fmt.Printf("\nruns: %d", runs)
			}
			p.y = (nextPos.y + o.y)
			p.x = (nextPos.x + o.x)

			if p.x < 0 {
				p.x = size - 1
			} else if p.x == size {
				p.x = 0
			}

			if p.y < 0 {
				p.y = size - 1
			} else if p.y == size {
				p.y = 0
			}

			v, ok := visits[p]
			if ok {
				if v == 0 {
					continue
				}

				if v <= step {
					skipped++
					if skipped%250_000_000*4 == 0 {
						end := time.Now().UnixMilli()
						fmt.Printf("\nskipped: %d, time: %d, visits size: %d", skipped, end-t, len(visits))
						t = end
					}
					continue
				}
			} else {
				calculated++
				if calculated%100 == 0 {
					fmt.Printf("\ncalculated: %dc", calculated)
				}
				the_rock := garden[p.x][p.y] == ROCK
				// if garden[p.x][p.y] != ROCK {
				if the_rock {
					visits[p] = 0
					continue
				}

				visits[p] = step
			}

			possibleEndTilesCnt++

			// we make all the steps in this "run"
			if step == STEPS_TO_MAKE {
				continue
			}

			move(p, step+1)
		}
	}

	move(nextPos, step)

	total := calcTotal(visits, garden)

	fmt.Printf("\n\npossibilities: %d", possibleEndTilesCnt)

	return total
}

func calcTotal(visits map[point]int, garden [][]byte) int {
	total := 0
	fmt.Println()
	for p, v := range visits {
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			garden[p.y][p.x] = 'O'
			total++
		}
	}
	return total
}
