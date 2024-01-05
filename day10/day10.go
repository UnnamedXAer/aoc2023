package day10

import (
	"bufio"
	"bytes"
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
	diagramSize := len(diagram)
	if diagramSize != len(diagram[0]) {
		panic("diagram not a square")
	}
	farthest, _ := traverse(diagram, startPos)
	fmt.Printf("\n\nfarthest: %d", farthest)
}

func traverse(diagram [][]byte, startPos point) (int, []point) {
	diagramSize := len(diagram)

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
			return (pathCnt / 2) - 1, pathPoints
		}
		if next.x != -1 {
			// fmt.Printf("\n%2d. moving to: %v", pathCnt-1, next)
			// fmt.Printf(" => %s", string(diagram[next.y][next.x]))
		} else {
			fmt.Printf("\ncouldn't find move from: %v", next)
			return -1, nil
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

	diagram, startPos := extractData()

	fmt.Printf("\n s is At: %+v -> %s", startPos, string(diagram[startPos.y][startPos.x]))
	fmt.Printf("\n%s", bytes.Join(diagram, []byte{'\n'}))
	_, path := traverse(diagram, startPos)
	tilesCount := countTiles(diagram, path)
	fmt.Printf("\n result: %d", tilesCount)
	// 736 -- too high
}

func countTiles(diagram [][]byte, path []point) int {

	pathMark := byte(':')

	for _, p := range path {
		diagram[p.y][p.x] = pathMark
	}
	fmt.Printf("\n\nDiagram with marked path:\n%s", bytes.Join(diagram, []byte{'\n'}))

	tilesCnt := 0
	rows := len(diagram)
	cols := len(diagram[0])

	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			p := point{y, x}
			if !isOnPath(path, y, x) {
				if isInside(path, p) {
					diagram[p.y][p.x] = 'I'
					tilesCnt++
				}
			}
		}
	}

	fmt.Printf("\n\nDiagram with marked inside:\n%s", bytes.Join(diagram, []byte{'\n'}))

	return tilesCnt
}

// https://www.codingninjas.com/studio/library/check-if-a-point-lies-in-the-interior-of-a-polygon
func isInside(path []point, p point) bool {

	y, x := p.y, p.x
	p1 := path[0]

	inside := false

	for _, p2 := range path {
		if y > min(p1.y, p2.y) {
			if y <= max(p1.y, p2.y) {
				if x <= max(p1.x, p2.x) {
					intersection_x := (y-p1.y)*(p2.x-p1.x)/(p2.y-p1.y) + p1.x

					if p1.x == p2.x || x <= intersection_x {
						inside = !inside
					}
				}
			}
		}
		p1 = p2
	}

	return inside
}

func countTiles_olc(diagram [][]byte, path []point) int {

	pathMark := byte(' ')
	outsideMark := byte('`')

	for _, p := range path {
		diagram[p.y][p.x] = pathMark
	}
	fmt.Printf("\n\nDiagram with marked path:\n%s", bytes.Join(diagram, []byte{'\n'}))

	rows := len(diagram)
	cols := len(diagram[0])

	q := make([]point, 0, 100)

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if diagram[y][x] == pathMark {
				break
			}
			diagram[y][x] = outsideMark
			q = append(q, point{y, x})
		}
	}

	for y := 0; y < rows; y++ {
		for x := cols - 1; x >= 0; x-- {
			if diagram[y][x] == pathMark {
				break
			}

			// diagram[y][x] = outsideMark
			q = append(q, point{y, x})
		}
	}

	fmt.Printf("\n\nDiagram after initial marking Outside:\n%s\n", bytes.Join(diagram, []byte{'\n'}))

	isFree := func(diagram [][]byte, p point) bool {
		return diagram[p.y][p.x] != pathMark && diagram[p.y][p.x] != outsideMark
	}

	top := len(q)
	startingPos := q[top-1]
	fmt.Printf("\n\nstarting at: %v => %s", startingPos, string(diagram[startingPos.y][startingPos.x]))

	for top > 0 {
		top--
		p := q[0]
		q = q[1:]
		diagram[p.y][p.x] = outsideMark

		for y := p.y - 1; y < p.y+3; y++ {
			if y < 0 || y >= rows {
				continue
			}
			for x := p.x - 1; x < p.x+3; x++ {
				if x < 0 || x >= cols {
					continue
				}

				pp := point{y, x}
				if isFree(diagram, pp) {
					q = append(q, pp)
					top++
				}
			}
		}
	}

	tailsCnt := 0
	outsideCnt := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if diagram[y][x] == outsideMark {
				outsideCnt++
			}

			if isFree(diagram, point{y, x}) {
				tailsCnt++
			}
		}
	}

	fmt.Printf("\n\nDiagram finale:\n%s\n", bytes.Join(diagram, []byte{'\n'}))

	pathCnt := len(path) - 2 // we have twice added 'S' at the end
	calculatedTailsCnt := rows*cols - outsideCnt - pathCnt
	fmt.Printf("\n   totalCnt: %d", rows*cols)
	fmt.Printf("\n    pathCnt: %d", pathCnt)
	fmt.Printf("\n outsideCnt: %d", outsideCnt)
	fmt.Printf("\ncalc tailsCnt: %d", calculatedTailsCnt)
	fmt.Printf("\n   tailsCnt: %d", tailsCnt)

	return tailsCnt
}

func isOnPath(path []point, y, x int) bool {
	return slices.Contains(path, point{y, x})
}

func addTile(diagram [][]byte, rows, cols int, path []point, tiles *[]point, tilePoint point) {
	// if tilePoint.x < cols && tilePoint.y < rows {
	if !slices.Contains(path, tilePoint) {
		if !slices.Contains(*tiles, tilePoint) {
			*tiles = append(*tiles, tilePoint)
		} else {
			fmt.Printf("\n2. tile exists: %v", tilePoint)
		}
	}

}
