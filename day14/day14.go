package day14

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

const roundedRock byte = 'O'
const cubeRock byte = '#'
const emptySpace byte = '.'

func extractData() [][]byte {

	f, err := os.Open("./day14/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)
	allLines := make([][]byte, 0, 100)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		allLines = append(allLines, line)
	}

	help.IfErr(scanner.Err())

	return allLines
}

func Ex1() {
	dish := extractData()
	print(string(bytes.Join(dish, []byte{'\n'})))
	total := processDishEx1(dish)

	fmt.Printf("\n\n Total: %d", total)
}

func processDishEx1(dish [][]byte) int {

	moveRocksEx1(dish)
	fmt.Printf("\n\n%s", string(bytes.Join(dish, []byte{'\n'})))

	total := calcLoad(dish)

	return total
}

func moveRocksEx1(dish [][]byte) {

	rows := len(dish)
	cols := len(dish[0])
	for colIdx := 0; colIdx < cols; colIdx++ {
		lastObstacle := -1
		for rowIdx := 0; rowIdx < rows; rowIdx++ {
			rock := dish[rowIdx][colIdx]
			switch rock {
			case roundedRock:
				lastObstacle += 1
				if lastObstacle < rowIdx {
					dish[lastObstacle][colIdx] = roundedRock
					dish[rowIdx][colIdx] = emptySpace
				}

			case cubeRock:
				lastObstacle = rowIdx
			case emptySpace:

			}
		}

	}

}

func calcLoad(dish [][]byte) int {
	total := 0

	rows := len(dish)
	cols := len(dish[0])
	for rowIdx := 0; rowIdx < rows; rowIdx++ {
		cnt := 0
		for colIdx := 0; colIdx < cols; colIdx++ {
			if dish[rowIdx][colIdx] == roundedRock {
				cnt++
			}
		}

		total += cnt * (rows - rowIdx)
	}

	return total
}

func Ex2() {
	dish := extractData()
	// print(string(bytes.Join(dish, []byte{'\n'})))
	cycles := 1_000_000_000
	// cycles := 10_000_000
	// cycles := 300
	total := processDish(dish, cycles)
	// fmt.Printf("\n\n%s", string(bytes.Join(dish, []byte{'\n'})))

	fmt.Printf("\n\n Total: %d", total)
}

func processDish(dish [][]byte, cycles int) int {
	sep := []byte{'|'}
	cache := make(map[string]int, 200)

	// const logEveryRows = 10_000_000
	// fmt.Printf("\n expect x:  %9d logs\n", cycles/logEveryRows)

	var key string

	i := 0
	firstHitIdx := -1
	firstHitKey := ""
	skipCache := false
	for ; i < cycles; i++ {
		key = string(bytes.Join(dish, sep))
		for k := 0; k < 4; k++ {
			moveRocks(dish, k)
		}
		if skipCache {
			continue
		}

		_, ok := cache[key]
		if ok {
			if firstHitIdx == -1 {
				firstHitIdx = i
				firstHitKey = key
			} else if key == firstHitKey {
				skipCache = true

				cacheCycleSize := i - firstHitIdx

				// fmt.Println()
				// f := firstHitIdx
				// u := 0
				// for ; f < cycles; f += cacheCycleSize {
				// 	u++
				// }

				uu := ((cycles - firstHitIdx) / cacheCycleSize)

				// fmt.Printf("\n u: %d, uu: %d", u, uu)

				// f -= cacheCycleSize
				ff := firstHitIdx + cacheCycleSize*uu
				// if ff != f {
				// 	fmt.Printf("\n f: %d, ff: %d", f, ff)
				// }
				// i = f
				i = ff
			}

		} else {
			rollLoad := calcLoad(dish)
			cache[key] = rollLoad
		}

	}

	total := calcLoad(dish)

	return total
}

func moveRocks(dish [][]byte, direction int) {

	rows := len(dish)
	cols := len(dish[0])

	switch direction {
	case 0: // to top
		for colIdx := 0; colIdx < cols; colIdx++ {
			lastObstacle := -1
			for rowIdx := 0; rowIdx < rows; rowIdx++ {
				rock := dish[rowIdx][colIdx]
				switch rock {
				case roundedRock:
					lastObstacle += 1
					if lastObstacle < rowIdx {
						dish[lastObstacle][colIdx] = roundedRock
						dish[rowIdx][colIdx] = emptySpace
					}

				case cubeRock:
					lastObstacle = rowIdx
				case emptySpace:

				}
			}
		}
	case 1: // to left
		for rowIdx := 0; rowIdx < rows; rowIdx++ {
			lastObstacle := -1
			for colIdx := 0; colIdx < cols; colIdx++ {
				rock := dish[rowIdx][colIdx]
				switch rock {
				case roundedRock:
					lastObstacle += 1
					if lastObstacle < colIdx {
						dish[rowIdx][lastObstacle] = roundedRock
						dish[rowIdx][colIdx] = emptySpace
					}
				case cubeRock:
					lastObstacle = colIdx
				case emptySpace:
				}
			}
		}
	case 2: // to bottom
		for colIdx := cols - 1; colIdx >= 0; colIdx-- {
			lastObstacle := rows
			for rowIdx := rows - 1; rowIdx >= 0; rowIdx-- {
				rock := dish[rowIdx][colIdx]
				switch rock {
				case roundedRock:
					lastObstacle -= 1
					if lastObstacle > rowIdx {
						dish[lastObstacle][colIdx] = roundedRock
						dish[rowIdx][colIdx] = emptySpace
					}

				case cubeRock:
					lastObstacle = rowIdx
				case emptySpace:

				}
			}
		}
	case 3: // to right
		for rowIdx := rows - 1; rowIdx >= 0; rowIdx-- {
			lastObstacle := cols
			for colIdx := cols - 1; colIdx >= 0; colIdx-- {
				rock := dish[rowIdx][colIdx]
				switch rock {
				case roundedRock:
					lastObstacle -= 1
					if lastObstacle > colIdx {
						dish[rowIdx][lastObstacle] = roundedRock
						dish[rowIdx][colIdx] = emptySpace
					}
				case cubeRock:
					lastObstacle = colIdx
				case emptySpace:
				}
			}
		}
	}
}
