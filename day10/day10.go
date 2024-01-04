package day10

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type point struct {
	y, x int
}

func extractData() ([][]byte, point) {
	f, err := os.Open("./day10/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	diagram := make([][]byte, 0, 140)
	sCol := -1
	sRow := -1
	linesCnt := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) == 0 {
			break
		}
		diagram = append(diagram, line)
		if sCol == -1 {
			sCol = slices.Index(diagram[linesCnt], 'S')
			if sCol != -1 {
				sRow = linesCnt
			}
		}
		linesCnt++
	}

	if sCol == -1 || sRow == -1 {
		panic("missing s point")
	}

	return diagram, point{y: sRow, x: sCol}
}

func Ex1() {

	diagram, startPos := extractData()

	fmt.Printf("\n s is At: %+v -> %s", startPos, string(diagram[startPos.y][startPos.x]))
	// fmt.Printf("\n%s", bytes.Join(diagram, []byte{'\n'}))

	diagramSize := len(diagram)
	if diagramSize != len(diagram[0]) {
		panic("diagram not a square")
	}

	farthest := traverse(diagram, startPos)
	fmt.Printf("\n\nfarthest: %d", farthest)

}

func traverse(diagram [][]byte, startPos point) int {
	diagramSize := len(diagram)

	// fmt.Printf("\n\n 0.     first: %v => %s", startPos, string(diagram[startPos.y][startPos.x]))

	finish := false
	next := startPos
	pathPoints := make([]point, 0, 10)
	pathPoints = append(pathPoints, startPos)

	pathCnt := len(pathPoints)
	for {
		next, finish = getNext(diagram, diagramSize, pathPoints[max(pathCnt-2, 0)], next)
		pathPoints = append(pathPoints, next)
		pathCnt++
		if finish {
			fmt.Printf("\npathCnt: %d", pathCnt)
			return (pathCnt / 2) - 1
		}
		if next.x != -1 {
			// fmt.Printf("\n%2d. moving to: %v", pathCnt-1, next)
			// fmt.Printf(" => %s", string(diagram[next.y][next.x]))
		} else {
			fmt.Printf("\ncouldn't find move from: %v", next)
			return -1
		}
	}
}

// getNext is looking for a next available pipe to be traversed
func getNext(diagram [][]byte, diagramSize int, prevPos, currentPos point) (point, bool) {
	prev := diagram[prevPos.y][prevPos.x]
	curr := diagram[currentPos.y][currentPos.x]

	switch curr {
	case '|':
		if prevPos.y < currentPos.y {
			return point{currentPos.y + 1, currentPos.x}, false
		}
		if prevPos.y > currentPos.y {
			return point{currentPos.y - 1, currentPos.x}, false
		}
		return point{-1, -1}, false
	case '-':

		if prevPos.x < currentPos.x {
			return point{currentPos.y, currentPos.x + 1}, false
		}
		if prevPos.x > currentPos.x {
			return point{currentPos.y, currentPos.x - 1}, false
		}
		return point{-1, -1}, false

	case 'L': // |_
		if prevPos.x > currentPos.x {
			return point{currentPos.y - 1, currentPos.x}, false
		}
		if prevPos.y < currentPos.y {
			return point{currentPos.y, currentPos.x + 1}, false
		}
		return point{-1, -1}, false

	case 'J': // _|

		if prevPos.x < currentPos.x {
			return point{currentPos.y - 1, currentPos.x}, false
		}
		if prevPos.y < currentPos.y {
			return point{currentPos.y, currentPos.x - 1}, false
		}
		return point{-1, -1}, false

	case '7': // -|

		if prevPos.x < currentPos.x {
			return point{currentPos.y + 1, currentPos.x}, false
		}
		if prevPos.y > currentPos.y {
			return point{currentPos.y, currentPos.x - 1}, false
		}
		return point{-1, -1}, false

	case 'F': // |-

		if prevPos.x > currentPos.x {
			return point{currentPos.y + 1, currentPos.x}, false
		}
		if prevPos.y > currentPos.y {
			return point{currentPos.y, currentPos.x + 1}, false
		}
		return point{-1, -1}, false

	case '.':
	case 'S':
		if prev == 'S' {
			// we are starting
			if currentPos.y > 0 {
				p := point{currentPos.y - 1, currentPos.x}
				next, _ := getNext(diagram, diagramSize, currentPos, p)
				if next.y != -1 {
					return p, false
				}
			}
			if currentPos.y < diagramSize-1 {
				p := point{currentPos.y + 1, currentPos.x}
				next, _ := getNext(diagram, diagramSize, currentPos, p)
				if next.y != -1 {
					return p, false
				}
			}

			if currentPos.x > 0 {
				p := point{currentPos.y, currentPos.x - 1}
				next, _ := getNext(diagram, diagramSize, currentPos, p)
				if next.y != -1 {
					return p, false
				}
			}
			if currentPos.x < diagramSize-1 {
				p := point{currentPos.y, currentPos.x + 1}
				next, _ := getNext(diagram, diagramSize, currentPos, p)
				if next.y != -1 {
					return p, false
				}
			}

			panic("cannot move from the start")
		}

		// we finished
		return currentPos, true

	}

	return point{-1, -1}, false
}

func Ex2() {

}
