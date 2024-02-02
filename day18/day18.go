package day18

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type direction byte

type digSpec struct {
	dir   direction
	len   int
	color string
}

func extractData() []digSpec {

	f, err := os.Open("./day18/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	plan := make([]digSpec, 0, 750)
	i := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		i++
		n := int(line[2] - '0')
		colorIdx := 5

		if line[3] != ' ' {
			colorIdx += 1
			n = n*10 + int(line[3]-'0')
		}

		spec := digSpec{
			dir:   direction(line[0]),
			len:   n,
			color: string(line[colorIdx : len(line)-1]),
		}

		plan = append(plan, spec)
	}

	help.IfErr(scanner.Err())

	return plan
}

func Ex1() {
	plan := extractData()
	total := solvePart1(plan)

	fmt.Printf("\n\n  Total: %d", total)
	// 31729 - too high
}

func solvePart1(plan []digSpec) int {
	outline := extractOutlineInfo(plan)
	printLagoon(outline, false)
	fmt.Printf("\n\n outline total: %d", len(outline.path))

	printLagoon(outline, true)
	total := calcTrench(outline)

	return total
}

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

type outlineInfo struct {
	path                   []point
	minY, minX, maxY, maxX int
}
type point struct {
	y, x int
}

func extractOutlineInfo(plan []digSpec) outlineInfo {

	specsCnt := len(plan)
	y, x, maxY, maxX, minY, minX := 0, 0, 0, 0, 0, 0

	path := make([]point, 0, 1000)
	// path = append(path, point{0, 0})

	for i := 0; i < specsCnt; i++ {
		dir := plan[i].dir
		length := plan[i].len
		switch dir {
		case 'U':
			y--
			nextY := y - length
			for ; y > nextY; y-- {
				path = append(path, point{y, x})
			}
			if y < minY {
				minY = y
			}
			y++
		case 'D':
			y++
			nextY := y + length
			for ; y < nextY; y++ {
				path = append(path, point{y, x})
			}
			if y > maxY {
				maxY = y
			}
			y--
		case 'L':
			x--
			nextX := x - length
			for ; x > nextX; x-- {
				path = append(path, point{y, x})
			}
			if x < minX {
				minX = x
			}
			x++
		case 'R':
			x++
			nextX := x + length
			for ; x < nextX; x++ {
				path = append(path, point{y, x})
			}
			if x > maxX {
				maxX = x
			}
			x--
		default:
			panic("unknown direction: " + string(dir))
		}
	}

	return outlineInfo{path, minY, minX, maxY, maxX}

	// for i := 0; i < specsCnt; i++ {
	// 	dir := plan[i].dir
	// 	length := plan[i].len + 1
	// 	switch dir {
	// 	case 'U':

	// 		y -= length
	// 		if y < minY {
	// 			minY = y
	// 		}
	// 	case 'D':
	// 		y += length
	// 		if y > maxY {
	// 			maxY = y
	// 		}
	// 	case 'L':
	// 		x -= length
	// 		if x < minX {
	// 			minX = x
	// 		}
	// 	case 'R':
	// 		x += length
	// 		if x > maxX {
	// 			maxX = x
	// 		}
	// 	default:
	// 		panic("unknown direction: " + string(dir))
	// 	}
	// }

	// fmt.Printf("\n minY: %4d, minX: %4d, maxY: %4d, maxX: %4d", minY, minX, maxY, maxX)

	// outline := make([][]byte, 0, maxY)

	// maxY = maxY + -minY
	// maxX = maxX + -minX
	// y, x = -minY, -minX

	// templateLine := make([]byte, maxX)
	// for i := 0; i < maxX; i++ {
	// 	templateLine[i] = '.'
	// }

	// for i := 0; i < maxY; i++ {
	// 	outline = append(outline, make([]byte, maxX))
	// 	copy(outline[len(outline)-1], templateLine)
	// }
	// outline[y][x] = '#'

	// for i := 0; i < specsCnt; i++ {
	// 	dir := plan[i].dir
	// 	length := plan[i].len
	// 	switch dir {
	// 	case 'U':
	// 		y--
	// 		nextY := y - length
	// 		for ; y > nextY; y-- {
	// 			outline[y][x] = '#'
	// 		}
	// 		y++

	// 	case 'D':
	// 		y++
	// 		nextY := y + length
	// 		for ; y < nextY; y++ {
	// 			outline[y][x] = '#'
	// 		}
	// 		y--

	// 	case 'L':
	// 		x--
	// 		nextX := x - length
	// 		for ; x > nextX; x-- {
	// 			outline[y][x] = '#'
	// 		}
	// 		x++

	// 	case 'R':
	// 		x++
	// 		nextX := x + length
	// 		for ; x < nextX; x++ {
	// 			outline[y][x] = '#'
	// 		}
	// 		x--

	// 	default:
	// 		panic("unknown direction: " + string(dir))
	// 	}
	// }

	// return outline
}

func calcTrench(outline outlineInfo) int {

	// for _, p := range outline.path {
	// 	if !isInside(outline.path, p) {
	// 		ok := slices.Contains(outline.path, p)
	// 		fmt.Printf("\np: %#v - %t", p, ok)
	// 	} else {
	// 		fmt.Printf("\np: %#v +", p)
	// 	}
	// }

	total := 0

	path := outline.path
	for y := outline.minY; y < outline.maxY; y++ {
		for x := outline.minX; x < outline.maxX; x++ {
			if isInside(path, point{y, x}) || slices.Contains(path, point{y, x}) {
				total++
			}
		}
	}

	return total
}

func printLagoon(outline outlineInfo, fill bool) {

	ySize := -outline.minY + outline.maxY
	xSize := -outline.minX + outline.maxX

	lagoon := make([][]byte, 0, ySize)

	templateLine := make([]byte, xSize)
	for i := 0; i < xSize; i++ {
		templateLine[i] = '.'
	}

	for i := 0; i < ySize; i++ {
		lagoon = append(lagoon, make([]byte, xSize))
		copy(lagoon[len(lagoon)-1], templateLine)
	}

	path := outline.path
	for _, p := range path {
		lagoon[p.y+-outline.minY][p.x+-outline.minX] = '#'
	}

	if fill {
		for y := outline.minY; y < ySize; y++ {
			for x := outline.minX; x < xSize; x++ {
				if isInside(path, point{y, x}) {
					lagoon[y+-outline.minY][x+-outline.minX] = '#'
				}
			}
		}
	}

	fmt.Println()
	for _, v := range lagoon {
		fmt.Printf("\n%s", string(v))
	}
}

func Ex2() {

}
