package day11

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

type point struct {
	X, Y int64
}

type dezzCosmos struct {
	// width    int
	// height   int
	galaxies []point
	// pairsCnt int
}

func (c *dezzCosmos) String() string {
	return fmt.Sprintf("%v\ng: %d", "cosmos" /* c.width, c.height,*/, len(c.galaxies))
}

func extractData(expansionRatio point) *dezzCosmos {

	f, err := os.Open("./day11/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	cosmos := dezzCosmos{
		galaxies: make([]point, 0, 280),
	}

	var colHaveGalaxy []bool
	var rowHaveGalaxy []bool = make([]bool, 0, 140)
	cnt := int64(0)
	lineSize := 0
	for scanner.Scan() {

		line := scanner.Text()
		lineSize = len(line)
		if lineSize == 0 { // for tests
			break
		}

		rowHaveGalaxy = append(rowHaveGalaxy, false)
		if colHaveGalaxy == nil {
			colHaveGalaxy = make([]bool, lineSize)
		}
		galaxiesCnt := 0
		for i := int64(lineSize - 1); i >= 0; i-- {
			if line[i] == '#' {
				colHaveGalaxy[i] = true
				rowHaveGalaxy[cnt] = true
				galaxiesCnt++
				cosmos.galaxies = append(cosmos.galaxies, point{X: i, Y: cnt})
			}
		}

		cnt++
	}
	// cosmos.height = cnt
	// cosmos.width = lineSize

	expandCosmos(expansionRatio, &cosmos, colHaveGalaxy, rowHaveGalaxy)

	return &cosmos
}

func expandCosmos(expansionRation point, cosmos *dezzCosmos, horizontalLineHaveGalaxy []bool, verticalLineHaveGalaxy []bool) {
	if expansionRation.X > 1 {
		expansionRation.X -= 1
	}
	if expansionRation.Y > 1 {
		expansionRation.Y -= 1
	}

	offsetX := int64(0)
	galaxiesCnt := len(cosmos.galaxies)
	horizontalLineHaveGalaxySize := int64(len(horizontalLineHaveGalaxy))
	for i := int64(0); i < horizontalLineHaveGalaxySize; i++ {
		x := i + offsetX
		if horizontalLineHaveGalaxy[i] {
			continue
		}
		offsetX += expansionRation.X
		for k := 0; k < galaxiesCnt; k++ {
			if cosmos.galaxies[k].X > x {
				cosmos.galaxies[k].X += expansionRation.X
			}
		}
	}
	// cosmos.width += expansionRation.X * offsetX

	offsetY := int64(0)
	verticalLineHaveGalaxySize := int64(len(verticalLineHaveGalaxy))
	for i := int64(0); i < verticalLineHaveGalaxySize; i++ {
		y := i + offsetY
		if verticalLineHaveGalaxy[i] {
			continue
		}
		offsetY += expansionRation.Y
		for k := 0; k < galaxiesCnt; k++ {
			if cosmos.galaxies[k].Y > y {
				cosmos.galaxies[k].Y += expansionRation.Y
			}
			offsetY += 0
		}
	}
	// cosmos.height += expansionRation.Y * offsetY
}

func Ex1() {

	exercise(point{X: 1, Y: 1})
}

// just use Point type to pass expansion values
func exercise(expansionRatio point) {
	cosmos := extractData(expansionRatio)
	// fmt.Printf("%v", cosmos)
	total := calculateTotalOfClosestPaths(cosmos)
	fmt.Printf("\n\nTotal: %d", total)
}

func calculateTotalOfClosestPaths(cosmos *dezzCosmos) int64 {

	cnt := 0

	total := int64(0)
	for i := len(cosmos.galaxies) - 1; i >= 0; i-- {
		g1 := cosmos.galaxies[i]
		for k := i - 1; k >= 0; k-- {
			cnt++
			g2 := cosmos.galaxies[k]
			dist := calcClosestDistance(g1, g2)
			total += dist
			// fmt.Printf(" - dist %v - %v => %v", g1, g2, dist)
		}
	}

	return total
}

func calcClosestDistance(g1, g2 point) int64 {

	yDif := g1.Y - g2.Y
	xDif := g1.X - g2.X
	if yDif < 0 {
		yDif = -yDif
	}
	if xDif < 0 {
		xDif = -xDif
	}

	// if yDif < xDif {
	// return	2*yDif + (xDif - yDif)
	// }
	return 2*xDif + (yDif - xDif)
	// fmt.Printf("\ny: %d, x: %d", yDif, xDif)

}

func Ex2() {
	exercise(point{X: 1, Y: 1})
	exercise(point{X: 10, Y: 10})
	exercise(point{X: 100, Y: 100})
	exercise(point{X: 1000000, Y: 1000000})
}
