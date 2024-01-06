package day11

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

type dezzCosmos struct {
	width    int
	height   int
	galaxies []help.Point
	pairsCnt int
}

func (c *dezzCosmos) String() string {
	return fmt.Sprintf("%v\nw: %d, h: %d, g: %d, p: %d", "cosmos", c.width, c.height, len(c.galaxies), c.pairsCnt)
}

func extractData() *dezzCosmos {

	f, err := os.Open("./day11/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	cosmos := dezzCosmos{
		galaxies: make([]help.Point, 0, 280),
	}

	var colHaveGalaxy []bool
	var rowHaveGalaxy []bool = make([]bool, 0, 140)
	cnt := 0
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
		for i := lineSize - 1; i >= 0; i-- {
			if line[i] == '#' {
				colHaveGalaxy[i] = true
				rowHaveGalaxy[cnt] = true
				galaxiesCnt++
				cosmos.galaxies = append(cosmos.galaxies, help.Point{X: i, Y: cnt})
			}
		}

		cnt++
	}
	cosmos.height = cnt
	cosmos.width = lineSize

	offsetX := 0
	galaxiesCnt := len(cosmos.galaxies)
	for i := 0; i < len(colHaveGalaxy); i++ {
		x := i + offsetX
		if colHaveGalaxy[i] {
			continue
		}
		offsetX++
		for k := 0; k < galaxiesCnt; k++ {
			if cosmos.galaxies[k].X > x {
				cosmos.galaxies[k].X++
			}
		}
	}
	cosmos.width += offsetX

	offsetY := 0
	for i := 0; i < len(rowHaveGalaxy); i++ {
		y := i + offsetY
		if rowHaveGalaxy[i] {
			continue
		}
		offsetY++
		for k := 0; k < galaxiesCnt; k++ {
			if cosmos.galaxies[k].Y > y {
				cosmos.galaxies[k].Y++
			}
		}
	}
	cosmos.height += offsetY

	pairsCnt := 0
	for i := 1; i < len(cosmos.galaxies); i++ {
		pairsCnt += i
	}

	cosmos.pairsCnt = pairsCnt

	return &cosmos
}

func Ex1() {
	cosmos := extractData()

	// fmt.Printf("%v", cosmos)

	fmt.Println()

	total := calculateTotalOfClosestPaths(cosmos)

	fmt.Printf("\n\nTotal: %d", total)
}

func calculateTotalOfClosestPaths(cosmos *dezzCosmos) int {

	cnt := 0

	total := 0
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

func calcClosestDistance(g1, g2 help.Point) int {

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

}
