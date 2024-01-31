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

	if len(blocks) != len(blocks[0]) {
		panic("blocks not a square")
	}

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

const MAX_STEPS int = 3

func Ex1() {
	blocks := extractData()
	size := len(blocks)

	total := solve(blocks, size)

	// position := d[size-1][size-1]
	// total := math.MaxInt
	// for _, dirs := range position.dirs {
	// 	for _, heatLoss := range dirs {
	// 		if heatLoss < total {
	// 			total = heatLoss
	// 		}
	// 	}
	// }

	fmt.Printf("\n\n  Total: %d", total)
}

type PositionDirInfo []int // [3]int

type Position struct {
	y, x     int
	heatLoss int
	dirs     map[direction]PositionDirInfo
}

func solve(graph [][]int, size int) int {

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

				dirs: map[direction]PositionDirInfo{
					north: {v, v, v},
					south: {v, v, v},
					west:  {v, v, v},
					east:  {v, v, v},
				},
			}

		}
	}

	positions[0][0].heatLoss = 0

	q := priorityQueue{}
	q.push(&pqElement{key: point{0, 0, 0, noMove}, priority: 0})

	for !q.isEmpty() {
		p := q.pop()

		nextPos := move(p.key, positions, west, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = move(p.key, positions, east, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = move(p.key, positions, north, size)
		if nextPos != nil {
			q.push(&pqElement{key: *nextPos, priority: positions[nextPos.y][nextPos.x].heatLoss})
		}

		nextPos = move(p.key, positions, south, size)
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
func move(p point, positions [][]Position, heading direction, size int) *point {

	nextP := addOffsetToPoint(p, heading)
	if outsizeRange(size, nextP.y, nextP.x) {
		return nil
	}

	positionFrom := positions[p.y][p.x]

	// if we were heading west in previous step(s)
	if p.dir == heading {

		// newDistance := positionFrom.dirs[north][p.distance] + 1
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
	currHeatLoss := 0
	// point at 0,0 has dir noMove
	if p.dir >= north {
		currHeatLoss = positionFrom.dirs[p.dir][p.distance]
	}
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

func getRowAndCol(n int, size int) (row, col int) {
	row = n / size
	col = n - (row * size)
	return row, col
}

func getDirection(idx int, prevIdx int, size int) direction {
	if prevIdx < 0 {
		return noMove
	}
	r, c := getRowAndCol(idx, size)
	pr, pc := getRowAndCol(prevIdx, size)

	if r < pr {
		return north
	}
	if r > pr {
		return south
	}
	if c < pc {
		return west
	}
	if c > pc {
		return east
	}

	return unknown
}

func getIdx(r, c, size int) int {
	return r*size + c
}
