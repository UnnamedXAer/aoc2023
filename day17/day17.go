package day17

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]int {

	f, err := os.Open("./day17/data_t.txt")
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
var offsets = [...][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var results = make([]int, 0, 100)

func Ex1() {
	blocks := extractData()
	size := len(blocks)

	q := priorityQueue{}

	for i, v := range blocks {
		for j, v := range v {
			q.enqueue(i*size+j, v)
		}
	}

	for !q.isEmpty() {
		fmt.Printf(", %v", q.dequeue())
	}

	graph := buildGraph(blocks)

	p, d := Dijkstra(0, graph, size)

	pos := size*size - 1
	logs := ""
	fmt.Println()
	for pos > -1 {
		direction := getDirection(pos, p[pos], size)
		logs = fmt.Sprintf("\n I'm at: %3d, came from: %3d, moving: %3s", pos, p[pos], directionsTranslation[direction]) + logs
		pos = p[pos]
	}

	fmt.Println(logs)

	fmt.Printf("\n\np:  %+v", p)
	fmt.Printf("\n\nd:  %+v", d)

	fmt.Printf("\n\n  Total: %d", d[size*size-1])
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

func Dijkstra(startAt int, weights [][]int, size int) (paths, distances []int) {
	V := len(weights)
	var v, w, d int
	q := priorityQueue{}
	paths = make([]int, V)
	distances = make([]int, V)

	for v = 0; v < V; v++ {
		distances[v] = -1
	}
	distances[startAt] = 0
	q.enqueue(startAt, weights[startAt][startAt])

	var prevDirection direction
	var straightSteps int

	for !q.isEmpty() {
		v = q.dequeue().key

		for w = 0; w < V; w++ {

			// ///////////////////
			pos := w
			fmt.Printf("\npaths: ")
			for pos < len(paths) && pos > 0 {
				fmt.Printf(", %d", pos)
				dir := getDirection(pos, paths[pos], size)
				if dir == prevDirection {
					straightSteps++
					if straightSteps > 3 {
						break
					}
				} else {
					prevDirection = dir
					break
				}
				pos = paths[pos]
			}

			fmt.Printf(" - direction: %s", directionsTranslation[prevDirection])

			// fmt.Printf(", %d", pos)
			/////////////////////

			if weights[v][w] > 0 {
				// if straightSteps < 4 {
				// TODO: can we travers back to check direction? or can we track the direction in each step?
				d = distances[v] + weights[v][w]

				if distances[w] == -1 {
					distances[w] = d
					q.enqueue(w, d)
					paths[w] = v
				}

				if distances[w] > d {
					distances[w] = d
					q.updatePriority(w, d)
					paths[w] = v
				}
				// }
			}
		}
	}

	paths[startAt] = -1
	return paths, distances
}

func Ex2() {

}
