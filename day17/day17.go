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
	north direction = iota
	south
	west
	east
)

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

	p, d := Dijkstra(0, graph)

	fmt.Printf("\n\np:  %+v", p)
	fmt.Printf("\n\nd:  %+v", d)
}

func Dijkstra(startAt int, weights [][]int) (paths, distances []int) {
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

	for !q.isEmpty() {
		v = q.dequeue().key

		for w = 0; w < V; w++ {
			// if Adj[v][w] == 1 {
			if weights[v][w] > 0 /* and steps <4*/ {
				// TODO: can we travers  back here to check direction? or can we track the direction in each step?
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
			}
		}
	}

	paths[startAt] = -1
	return paths, distances
}

func Ex2() {

}
