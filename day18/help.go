package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

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
		fmt.Printf("\n%3d. %+v", i, spec)

		plan = append(plan, spec)
	}

	help.IfErr(scanner.Err())

	return plan
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

func printLagoon(outline outlineInfo, fill bool) [][]byte {

	ySize := outline.sizeY //-outline.minY + outline.maxY
	xSize := outline.sizeX //-outline.minX + outline.maxX

	lagoon := make([][]byte, 0, ySize)

	templateLine := make([]byte, xSize)
	templateX := make([]byte, xSize)
	for i := 0; i < xSize; i++ {
		templateLine[i] = '.'
		templateX[i] = byte(i%10) + '0'
	}

	for i := 0; i < ySize; i++ {
		lagoon = append(lagoon, make([]byte, xSize))
		copy(lagoon[len(lagoon)-1], templateLine)
	}

	path := outline.vertices
	for _, p := range path {
		lagoon[p.y][p.x] = '#'
		// lagoon[p.y+-outline.minY][p.x+-outline.minX] = '#'
	}

	if fill {
		for y := 0; y < ySize; y++ {
			for x := 0; x < xSize; x++ {
				if isInside(path, point{y, x}) {
					lagoon[y][x] = '#'
				}
			}
		}
	}

	lagoon = append(lagoon, templateX)

	fmt.Println()
	size := len(lagoon)
	for i, v := range lagoon {
		fmt.Printf("\n%3d %s", size-i-2, string(v))
	}

	return lagoon
}

func printLagoon2(outline outlineInfo) [][]byte {

	ySize := outline.sizeY
	xSize := outline.sizeX

	lagoon := make([][]byte, 0, ySize)

	templateLine := make([]byte, xSize)
	templateX := make([]byte, xSize)
	for i := 0; i < xSize; i++ {
		templateLine[i] = '.'
		templateX[i] = byte(i%10) + '0'
	}

	for i := 0; i < ySize; i++ {
		lagoon = append(lagoon, make([]byte, xSize))
		copy(lagoon[len(lagoon)-1], templateLine)
	}

	path := outline.vertices
	for _, p := range path {
		lagoon[p.y][p.x] = '#'
	}

	lagoon = append(lagoon, templateX)

	fmt.Println()
	size := len(lagoon)
	for i, v := range lagoon {
		fmt.Printf("\n%3d %s", size-i-2, string(v))
	}

	return lagoon
}
