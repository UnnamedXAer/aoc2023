package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]int {

	f, err := os.Open("./day17/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	blocks := make([][]int, 0, 141)

	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Bytes()
		blocks = append(blocks, make([]int, len(line)))

		for k, v := range line {
			blocks[i][k] = int(v - '0')
		}
	}

	help.IfErr(scanner.Err())

	return blocks
}

type direction int

const (
	unknown direction = iota
	noMove
	north
	south
	west
	east
	_directionsCnt
)

var directionsTranslation = [...]string{
	"unknown",
	"no move",
	"north",
	"south",
	"west",
	"east",
}
var directionsTranslationChar = [...]byte{
	'+',
	'O',
	'^',
	'v',
	'<',
	'>',
}

var offsets = [...][2]int{
	{math.MaxInt, math.MaxInt},
	{math.MaxInt, math.MinInt},
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var results = make([]int, 0, 100)

func Ex1() {
	blocks := extractData()
	size := len(blocks)

	total := solvePart1(blocks, size)

	fmt.Printf("\n\n  Total: %d", total)
}

type PositionDirInfo []int // [3]int

type Position struct {
	y, x     int
	heatLoss int
	dirs     [_directionsCnt]PositionDirInfo
	// dirs     map[direction]PositionDirInfo
}

func solvePart1(graph [][]int, size int) int {

	positions := make([][]Position, size)
	for i := 0; i < size; i++ {
		positions[i] = make([]Position, size)
		for j := 0; j < size; j++ {
			v := math.MaxInt
			if i == 0 && j == 0 {
				v = 0
			}

			positions[i][j] = Position{
				y:        i,
				x:        j,
				heatLoss: graph[i][j],

				dirs: [_directionsCnt]PositionDirInfo{
					{v, v, v},
					{v, v, v},
					{v, v, v},
					{v, v, v},
					{v, v, v},
					{v, v, v},
				},
			}

		}
	}

	positions[0][0].heatLoss = 0

	q := priorityQueue{}
	q.push(&pqElement{key: point{0, 0, 0, noMove}, priority: 0})

	for !q.isEmpty() {
		p := q.pop()

		nextPos := moveForPart1(p.key, positions, west, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart1(p.key, positions, east, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart1(p.key, positions, north, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart1(p.key, positions, south, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}
	}

	currentMinHeatLoss := math.MaxInt

	position := positions[size-1][size-1]
	for _, dirs := range position.dirs {
		for _, heatLoss := range dirs {
			if heatLoss < currentMinHeatLoss {
				currentMinHeatLoss = heatLoss
			}
		}
	}

	return currentMinHeatLoss
}

// generate next step from given step that is heading west
func moveForPart1(p point, positions [][]Position, heading direction, size int) *point {
	const MAX_STEPS int = 3

	nextP := addOffsetToPoint(p, heading)
	if outsizeRange(size, size, nextP.y, nextP.x) {
		return nil
	}

	positionFrom := positions[p.y][p.x]

	// if we were heading west in previous step(s)
	if p.dir == heading {

		newDistance := p.distance + 1
		if newDistance >= MAX_STEPS {
			return nil
		}

		positionTo := &positions[nextP.y][nextP.x]
		currHeatLoss := positionFrom.dirs[p.dir][p.distance]
		nextHeatLoss := currHeatLoss + positionTo.heatLoss

		if nextHeatLoss >= positionTo.dirs[heading][newDistance] {
			return nil
		}

		positionTo.dirs[heading][newDistance] = nextHeatLoss

		nextP.distance = newDistance
		return &nextP
	}

	// if we were heading e.g. east then, we cannot move back to west
	if heading == west {
		if p.dir == east {
			return nil
		}
	} else if heading == east {
		if p.dir == west {
			return nil
		}
	} else if heading == north {
		if p.dir == south {
			return nil
		}
	} else if heading == south {
		if p.dir == north {
			return nil
		}
	}

	// we were moving north or south in the given step
	// next step will be with distance 0

	newDistance := 0
	positionTo := &positions[nextP.y][nextP.x]
	currHeatLoss := positionFrom.dirs[p.dir][p.distance]
	nextHeatLoss := currHeatLoss + positionTo.heatLoss

	if nextHeatLoss >= positionTo.dirs[heading][newDistance] {
		return nil
	}

	positionTo.dirs[heading][newDistance] = nextHeatLoss

	nextP.distance = newDistance
	return &nextP
}

type point = queueE

func addOffsetToPoint(p point, dir direction) point {
	o := offsets[dir]
	p.y += o[0]
	p.x += o[1]
	p.dir = dir
	return p
}

// //////////////////////////////////////////////////////////////////////////////////
func Ex2() {
	blocks := extractData()
	ySize := len(blocks)
	xSize := len(blocks[0])

	total := solvePart2(blocks, ySize, xSize)

	fmt.Printf("\n\n Ex2:  Total: %d", total)
}

func solvePart2(graph [][]int, ySize, xSize int) int {

	positions := make([][]Position, ySize)
	for i := 0; i < ySize; i++ {
		positions[i] = make([]Position, xSize)
		for j := 0; j < xSize; j++ {
			v := math.MaxInt
			if i == 0 && j == 0 {
				v = 0
			}

			positions[i][j] = Position{
				y:        i,
				x:        j,
				heatLoss: graph[i][j],

				dirs: [_directionsCnt]PositionDirInfo{
					{v, v, v, v, v, v, v, v, v, v},
					{v, v, v, v, v, v, v, v, v, v},
					{v, v, v, v, v, v, v, v, v, v},
					{v, v, v, v, v, v, v, v, v, v},
					{v, v, v, v, v, v, v, v, v, v},
					{v, v, v, v, v, v, v, v, v, v},
				},
			}

		}
	}

	positions[0][0].heatLoss = 0

	q := priorityQueue{}
	q.push(&pqElement{key: point{0, 0, 0, noMove}, priority: 0})

	for !q.isEmpty() {
		p := q.pop()

		nextPos := moveForPart2(p.key, positions, west, ySize, xSize)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart2(p.key, positions, east, ySize, xSize)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart2(p.key, positions, north, ySize, xSize)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = moveForPart2(p.key, positions, south, ySize, xSize)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}
	}

	currentMinHeatLoss := math.MaxInt

	position := positions[ySize-1][xSize-1]
	for _, dirs := range position.dirs {
		for _, heatLoss := range dirs {
			if heatLoss < currentMinHeatLoss {
				currentMinHeatLoss = heatLoss
			}
		}
	}

	return currentMinHeatLoss
}

// generate next step from given step that is heading west
func moveForPart2(p point, positions [][]Position, heading direction, ySize, xSize int) *point {
	const MIN_STEPS int = 4 - 1
	const MAX_STEPS int = 10 - 1

	nextP := addOffsetToPoint(p, heading)
	if outsizeRange(ySize, xSize, nextP.y, nextP.x) {
		return nil
	}

	positionFrom := positions[p.y][p.x]

	// if we were heading west in previous step(s)
	if p.dir == heading {

		newDistance := p.distance + 1
		if newDistance > MAX_STEPS {
			return nil
		}

		positionTo := &positions[nextP.y][nextP.x]
		currHeatLoss := positionFrom.dirs[p.dir][p.distance]
		nextHeatLoss := currHeatLoss + positionTo.heatLoss

		if nextHeatLoss >= positionTo.dirs[heading][newDistance] {
			return nil
		}

		if nextP.y == ySize-1 && nextP.x == xSize-1 {
			if newDistance < MIN_STEPS {
				// fmt.Printf("\n0. cannot accept result with stepsCnt %d, heat loss: %4d | prev steps: %4d, turning %5s, from: %5s", newDistance, nextHeatLoss, p.distance, directionsTranslation[heading], directionsTranslation[p.dir])
				return nil
			}
			// fmt.Printf("\n0.     accepting result with stepsCnt %d, heat loss: %4d | prev steps: %4d, turning %5s, from: %5s", newDistance, nextHeatLoss, p.distance, directionsTranslation[heading], directionsTranslation[p.dir])
		}

		positionTo.dirs[heading][newDistance] = nextHeatLoss

		nextP.distance = newDistance
		return &nextP
	}

	if p.dir != noMove && p.distance < MIN_STEPS {
		return nil
	}

	// if we were heading e.g. east then, we cannot move back to west
	if heading == west {
		if p.dir == east {
			return nil
		}
	} else if heading == east {
		if p.dir == west {
			return nil
		}
	} else if heading == north {
		if p.dir == south {
			return nil
		}
	} else if heading == south {
		if p.dir == north {
			return nil
		}
	}

	// we were moving north or south in the given step
	// next step will be with distance 0

	newDistance := 0
	positionTo := &positions[nextP.y][nextP.x]
	currHeatLoss := positionFrom.dirs[p.dir][p.distance]
	nextHeatLoss := currHeatLoss + positionTo.heatLoss

	if nextHeatLoss >= positionTo.dirs[heading][newDistance] {
		return nil
	}

	if nextP.y == ySize-1 && nextP.x == xSize-1 {
		if newDistance < MIN_STEPS {
			// fmt.Printf("\n1. cannot accept result with stepsCnt %d, heat loss: %4d | prev steps: %4d, turning %5s, from: %5s", newDistance, nextHeatLoss, p.distance, directionsTranslation[heading], directionsTranslation[p.dir])
			return nil
		}
		// fmt.Printf("\n1.    accepting result with stepsCnt %d, heat loss: %4d | prev steps: %4d, turning %5s, from: %5s", newDistance, nextHeatLoss, p.distance, directionsTranslation[heading], directionsTranslation[p.dir])
	}

	positionTo.dirs[heading][newDistance] = nextHeatLoss

	nextP.distance = newDistance
	return &nextP
}
