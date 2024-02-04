package day18

import (
	"fmt"

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

// var inputNameSuffix = ""

var inputNameSuffix = "_t"

func Ex1() {
	plan := extractData()
	total := solvePart1(plan)

	fmt.Printf("\n\n  Total: %d", total)
	// 31729 - too high
}

func solvePart1(plan []digSpec) int {
	outline := extractOutlineInfo(plan)
	// printLagoon(outline, false)
	// fmt.Printf("\n\n outline total: %d", len(outline.vertices))

	// printLagoon(outline, true)
	total := calcTrench(outline)

	return total
}

type outlineInfo struct {
	vertices               []point
	minY, minX, maxY, maxX int
	sizeY, sizeX           int
}
type point struct {
	y, x int
}

func extractOutlineInfo(plan []digSpec) outlineInfo {

	specsCnt := len(plan)
	y, x, maxY, maxX, minY, minX := 0, 0, 0, 0, 0, 0

	path := make([]point, 0, 1000)

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

	for i := 0; i < len(path); i++ {
		path[i].x += -minX
		path[i].y += -minY
	}

	sizeX := maxX + -minX
	sizeY := maxY + -minY

	return outlineInfo{path, minY, minX, maxY, maxX, sizeY, sizeX}
}

func calcTrench(outline outlineInfo) int {
	total := 0

	vertices := outline.vertices
	for y := 0; y < outline.sizeY; y++ {
		for x := 0; x < outline.sizeX; x++ {
			v := point{y, x}

			inside := isInside(vertices, v)
			inVertices := slices.Contains(vertices, v)

			if inside || inVertices {
				total++
			}
		}
	}

	return total
}

// //////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////////////////////////////////

func Ex2() {

	plan := extractData2()
	fmt.Print()

	// for _, s := range plan {
	// 	fmt.Printf("\n %+v", s)
	// }

	total, totalPolygonArea := solvePart2(plan)
	// total+=len(plan)

	fmt.Printf("\n\n  Total: %d, polygon area: %d", total, totalPolygonArea)
}

func solvePart2(plan []digSpec) (int, int) {
	fmt.Println()
	fmt.Println()
	fmt.Println("PART 2")
	outline := extractOutlineInfo2(plan)
	// lagoon := printLagoon2(outline)
	// for _, s := range outline.vertices {
	// 	fmt.Printf("\n %+v", s)
	// }
	// for _, v := range outline.vertices {
	// 	lagoon[v.y][v.x] = 'O'
	// }

	// fmt.Printf("\n    marking vertices over lagoon from part 1")
	// for _, v := range lagoon {
	// 	fmt.Printf("\n%s", string(v))
	// }

	// outline = extractOutlineInfo2(plan)
	// printLagoon2(outline)
	// total := 0
	fmt.Printf("\n\n outline total: %d", len(outline.vertices))

	total := polygonArea(outline)
	// printLagoon(outline, true)
	// total := calcTrench2(outline)

	// tmpVertices := []point{
	// 	{0, 7},
	// 	{6, 7},
	// 	{6, 5},
	// 	{7, 5},
	// 	{7, 7},
	// 	{10, 7},
	// 	{10, 1},
	// 	{8, 1},
	// 	{8, 0},
	// 	{5, 0},
	// 	{5, 2},
	// 	{3, 2},
	// 	{3, 0},
	// 	{0, 0},
	// }

	// for i, v := range outline.vertices {
	// 	if i >= len(tmpVertices) {
	// 		break
	// 	}
	// 	fmt.Printf("\n%3d. %+v | %+v", i, tmpVertices[i], v)
	// }

	// totalPol := polygonArea(outlineInfo{vertices: tmpVertices})
	// t2 := polygonArea(outlineInfo{vertices: []point{
	// 	{0, 3},
	// 	{2, 3},
	// 	{2, 6},
	// 	{6, 6},
	// 	{6, 1},
	// 	{5, 1},
	// 	{5, 0},
	// 	{0, 0},
	// }})
	// fmt.Printf("\n\nt2: %d", t2)
	// total = polygonArea(outline)

	return total, -1
}

func polygonArea(outline outlineInfo) int {
	// #A function to apply the Shoelace algorithm
	vertices := outline.vertices
	numberOfVertices := len(vertices)
	sum1 := 0
	sum2 := 0

	// for i :range(0,numberOfVertices-1):
	for i := 0; i < numberOfVertices-1; i++ {
		sum1 = sum1 + vertices[i].x*vertices[i+1].y
		sum2 = sum2 + vertices[i].y*vertices[i+1].x
	}
	// #Add xn.y1
	sum1 = sum1 + vertices[numberOfVertices-1].x*vertices[0].y
	// #Add x1.yn
	sum2 = sum2 + vertices[0].x*vertices[numberOfVertices-1].y

	// area := math.Abs(sum1-sum2) / 2
	area := (sum1 - sum2) / 2
	if area < 0 {
		area = -area
	}

	return area
}

func extractOutlineInfo2(plan []digSpec) outlineInfo {

	specsCnt := len(plan)
	y, x, maxY, maxX, minY, minX := 0, 0, 0, 0, 0, 0

	vertices := make([]point, 0, 750)

	prevDir := plan[1].dir
	nextDir := direction(-1)

	for i := 0; i < specsCnt; i++ {
		dir := plan[i].dir
		length := plan[i].len
		if i == 0 {
			length += 1
		}
		if i < specsCnt-1 {
			nextDir = plan[i+1].dir
		} else {
			nextDir = plan[0].dir
		}
		switch dir {
		case up:
			if prevDir != nextDir {
				if prevDir == right && nextDir == left {
					length--
				} else {
					length++
				}
			}

			y = y + length
		case down:
			if prevDir != nextDir {
				if prevDir == left && nextDir == right {
					length--
				} else {
					length++
				}
			}

			y = y - length
		case right:
			if prevDir != nextDir {
				if prevDir == down && nextDir == up {
					length--
				} else {
					length++
				}
			}

			x = x + length
		case left:
			if prevDir != nextDir {
				if prevDir == up && nextDir == down {
					length--
				} else {
					length++
				}
			}

			x = x - length
		default:
			panic("unknown direction: " + string(byteToDir[dir]))
		}

		minY, maxY, minX, maxX = updateMinMax(y, minY, maxY, x, minX, maxX)
		prevDir = dir
		vertices = append(vertices, point{y, x})
	}

	sizeY := -minY + maxY + 1
	sizeX := -minX + maxX + 1

	return outlineInfo{vertices, minY, minX, maxY, maxX, sizeY, sizeX}
}

func updateMinMax(y int, minY int, maxY int, x int, minX int, maxX int) (int, int, int, int) {
	if y < minY {
		// fmt.Printf("\nchanging min y from %3d to %3d", minY, y)
		minY = y
	}
	if y > maxY {
		// fmt.Printf("\nchanging max y from %3d to %3d", maxY, y)
		maxY = y
	}
	if x < minX {
		// fmt.Printf("\nchanging min x from %3d to %3d", minX, x)
		minX = x
	}
	if x > maxX {
		// fmt.Printf("\nchanging max x from %3d to %3d", maxX, x)
		maxX = x
	}
	return minY, maxY, minX, maxX
}
