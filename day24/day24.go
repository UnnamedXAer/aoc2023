package day24

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

// const inputNameSuffix= ""
const inputNameSuffix = "_t"
const inputName = "./day24/data" + inputNameSuffix + ".txt"

type point struct{ x, y, z int64 }
type velocity struct{ x, y, z int64 }

type hailstone struct {
	p point
	v velocity
}

func extractData() any {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	hailstones := make([]hailstone, 0, 302)

	for scanner.Scan() {
		line := scanner.Bytes()
		size := len(line) - 1
		z, n := help.ReadNumValueFromEnd64(line, size)
		size = skipToNumber(line, size-n)
		y, n := help.ReadNumValueFromEnd64(line, size)
		size = skipToNumber(line, size-n)
		x, n := help.ReadNumValueFromEnd64(line, size)

		vel := velocity{
			x: x, y: y, z: z,
		}

		size = skipToNumber(line, size-n)

		z, n = help.ReadNumValueFromEnd64(line, size)
		size = skipToNumber(line, size-n)
		y, n = help.ReadNumValueFromEnd64(line, size)
		size = skipToNumber(line, size-n)
		x, _ = help.ReadNumValueFromEnd64(line, size)

		p := point{
			x: x, y: y, z: z,
		}

		hailstones = append(hailstones, hailstone{p, vel})
	}

	help.IfErr(scanner.Err())

	return hailstones
}

func skipToNumber(line []byte, size int) int {
	for !help.IsNumber(line[size]) {
		size--
	}
	return size
}

func Ex1() {
	x := extractData().([]hailstone)
	fmt.Printf("\n%v", x)

	fmt.Println()

	// no idea how to solve that :)

}

func trajectory(t int64, h hailstone) (int64, int64) {
	return h.p.x + h.v.x*t, h.p.y + h.v.y*t
}

const min = int64(7)
const max = int64(27)

// const min = int64(200000000000000)
// const max = int64(400000000000000)
