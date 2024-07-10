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

type point struct{ x, y, z int }
type velocity struct{ x, y, z int }

type world struct {
	points     []point
	velocities []velocity
	size       int
}

func extractData() any {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	world := world{
		points:     make([]point, 0, 301),
		velocities: make([]velocity, 0, 301),
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		size := len(line) - 1
		z, n := help.ReadNumValueFromEnd(line, size)
		size = skipToNumber(line, size-n)
		y, n := help.ReadNumValueFromEnd(line, size)
		size = skipToNumber(line, size-n)
		x, n := help.ReadNumValueFromEnd(line, size)

		vel := velocity{
			x: x, y: y, z: z,
		}

		size = skipToNumber(line, size-n)

		z, n = help.ReadNumValueFromEnd(line, size)
		size = skipToNumber(line, size-n)
		y, n = help.ReadNumValueFromEnd(line, size)
		size = skipToNumber(line, size-n)
		x, _ = help.ReadNumValueFromEnd(line, size)

		p := point{
			x: x, y: y, z: z,
		}

		world.points = append(world.points, p)
		world.velocities = append(world.velocities, vel)
	}

	help.IfErr(scanner.Err())

	world.size = len(world.points)
	return world
}

func skipToNumber(line []byte, size int) int {
	for !help.IsNumber(line[size]) {
		size--
	}
	return size
}

func Ex1() {
	x := extractData()
	fmt.Printf("\n%v", x)
}
