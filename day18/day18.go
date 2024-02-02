package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type direction int

const (
	right direction = iota
	down
	left
	up
)

var byteToDir = []byte{'R', 'D', 'L', 'U'}

type digSpec struct {
	dir direction
	len int
}

var inputNameSuffix = ""

// var inputNameSuffix= "_t"

func extractData() []digSpec {

	f, err := os.Open("./day18/data" + inputNameSuffix + ".txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	plan := make([]digSpec, 0, 750)
	i := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		i++
		n := int(line[2] - '0')

		if line[3] != ' ' {
			n = n*10 + int(line[3]-'0')
		}

		spec := digSpec{
			dir: direction(slices.Index(byteToDir, line[0])),
			len: n,
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
	fmt.Printf("\n\n outline total: %d", len(outline.vertices))

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
	vertices               []point
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
		case up:
			y--
			nextY := y - length
			for ; y > nextY; y-- {
				path = append(path, point{y, x})
			}
			if y < minY {
				minY = y
			}
			y++
		case down:
			y++
			nextY := y + length
			for ; y < nextY; y++ {
				path = append(path, point{y, x})
			}
			if y > maxY {
				maxY = y
			}
			y--
		case left:
			x--
			nextX := x - length
			for ; x > nextX; x-- {
				path = append(path, point{y, x})
			}
			if x < minX {
				minX = x
			}
			x++
		case right:
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
			panic("unknown direction: " + string(byteToDir[dir]))
		}
	}

	return outlineInfo{path, minY, minX, maxY, maxX}
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

	path := outline.vertices
	for y := outline.minY; y < outline.maxY; y++ {
		for x := outline.minX; x < outline.maxX; x++ {
			if isInside(path, point{y, x}) || slices.Contains(path, point{y, x}) {
				total++
			}
		}
	}

	return total
}

func printLagoon(outline outlineInfo, fill bool) [][]byte {

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

	path := outline.vertices
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

	return lagoon
}

// //////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////////////

func extractData2() []digSpec {

	f, err := os.Open("./day18/data" + inputNameSuffix + ".txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	plan := make([]digSpec, 0, 750)
	i := 0

	for scanner.Scan() {
		line := scanner.Bytes()

		i++
		colorIdx := 6

		if line[3] != ' ' {
			colorIdx += 1
		}

		n, err := strconv.ParseInt(string(line[colorIdx:len(line)-2]), 16, 64)
		if err != nil {
			fmt.Printf("\n line: %3d: %v", i, err)
			panic("\n")
		}

		spec := digSpec{
			dir: direction(slices.Index(byteToDir, line[0])),
			len: int(n),
		}

		plan = append(plan, spec)
	}

	help.IfErr(scanner.Err())

	return plan
}

func Ex2() {

	plan := extractData()
	fmt.Print()

	// for _, s := range plan {
	// 	fmt.Printf("\n %+v", s)
	// }

	total := solvePart2(plan)

	fmt.Printf("\n\n  Total: %d", total)
}

func solvePart2(plan []digSpec) int {
	outline := extractOutlineInfo(plan)
	printLagoon(outline, false)
	// part 2
	outline = extractOutlineInfo2(plan)
	printLagoon(outline, false)
	fmt.Printf("\n\n outline total: %d", len(outline.vertices))

	// printLagoon(outline, true)
	total := calcTrench2(outline)

	return total
}
func isInside2(path []point, p point) bool {

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

func calcTrench2(outline outlineInfo) int {

	// for _, p := range outline.path {
	// 	if !isInside(outline.path, p) {
	// 		ok := slices.Contains(outline.path, p)
	// 		fmt.Printf("\np: %#v - %t", p, ok)
	// 	} else {
	// 		fmt.Printf("\np: %#v +", p)
	// 	}
	// }

	total := 0

	path := outline.vertices
	for y := outline.minY; y < outline.maxY; y++ {
		for x := outline.minX; x < outline.maxX; x++ {
			if isInside2(path, point{y, x}) || slices.Contains(path, point{y, x}) {
				total++
			}
		}
	}

	return total
}

func extractOutlineInfo2(plan []digSpec) outlineInfo {

	specsCnt := len(plan)
	y, x, maxY, maxX, minY, minX := 0, 0, 0, 0, 0, 0

	vertices := make([]point, 0, 750)

	for i := 0; i < specsCnt; i++ {
		dir := plan[i].dir
		length := plan[i].len
		switch dir {
		case up:
			y = y - length
			if y < minY {
				minY = y
			}
		case down:
			y = y + length
			if y > maxY {
				maxY = y
			}
		case left:
			x = x - length
			if x < minX {
				minX = x
			}
		case right:
			x = x + length
			if x > maxX {
				maxX = x
			}
		default:
			panic("unknown direction: " + string(byteToDir[dir]))
		}
		vertices = append(vertices, point{y, x})
	}

	return outlineInfo{vertices, minY, minX, maxY + 1, maxX + 1}
}
