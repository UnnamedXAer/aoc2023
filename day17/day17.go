package day17

import (
	"bufio"
	"fmt"
	"math"
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
var directionsTranslationChar = [...]byte{
	'+',
	'O',
	'^',
	'v',
	'<',
	'>',
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

	// q := &priorityQueue{}

	// for i, v := range blocks {
	// 	for j, v := range v {
	// 		// q.Push(i*size+j, v)
	// 		q.Push(&queElement{i*size + j, v})
	// 	}
	// }

	// for !q.isEmpty() {
	// 	fmt.Printf(", %v", q.Pop())
	// }

	graph := buildGraph(blocks)
	// p, d := DijkstraMatrix(0, graph, size)

	p, d := Dijkstra3(graph, size*size, 0)

	// return
	// blocks[0][0] = 0
	// adj := buildGraphsAdj(blocks)
	// p, d := DijkstraAdjLL(0, adj, size)

	pos := size*size - 1
	logs := ""
	fmt.Println()
	for pos > -1 {
		direction := getDirection(pos, p[pos], size)
		logs = fmt.Sprintf("\n I'm at: %3d, came from: %3d, moving: %3s", pos, p[pos], directionsTranslation[direction]) + logs
		pos = p[pos]
	}

	fmt.Println(logs)

	fmt.Println()

	graphpath := make([][]byte, size)

	for i := 0; i < size; i++ {
		graphpath[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			graphpath[i][j] = '~'
		}
	}

	pos = size*size - 1
	for pos > -1 {
		r, c := getRowAndCol(pos, size)
		d := getDirection(pos, p[pos], size)
		graphpath[r][c] = directionsTranslationChar[d]
		pos = p[pos]
	}

	fmt.Println()
	for _, v := range graphpath {
		fmt.Printf("\n%s", string(v))
	}

	fmt.Printf("\n\np:  %+v", p)
	fmt.Printf("\n\nd:  %+v", d)

	fmt.Printf("\n\n  Total: %d", d[size*size-1])
	// 858 - too high
}

func solve(graph [][]int) {

	size := len(graph)

	currentPrices := make([][]int, size)

	for i := 0; i < size; i++ {
		currentPrices[i] = make([]int, size)
		for k := 0; k < size; k++ {
			currentPrices[i][k] = math.MaxInt
		}
	}

	currentPrices[0][0] = 0
	q := queue{}

	q.push(&queueE{0, 0})

	for !q.isEmpty() {

	}

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

func addOffset(r, c int, offset [2]int) (int, int) {
	return r + offset[0], c + offset[1]
}

func calcIdx(r, c, size int) int {
	return r*size + c
}

type gNode struct {
	vertexIdx int
	next      *gNode
	dir       direction
	weight    int
}

func buildGraphsAdj(blocks [][]int) []*gNode {
	size := len(blocks)
	adj := make([]*gNode, size*size)
	var prev *gNode

	for i := 0; i < size*size; i++ {
		r, c := getRowAndCol(i, size)
		adj[i] = &gNode{vertexIdx: -1} // create dummy node as to not be required to make special logic for checking if prev is nil ect.
		prev = adj[i]
		for k := len(offsets) - 1; k >= 0; k-- {
			nr, nc := addOffset(r, c, offsets[k])
			if outsizeRange(size, nr, nc) {
				continue
			}
			vertexIdx := calcIdx(nr, nc, size)
			n := &gNode{
				vertexIdx: vertexIdx,
				next:      nil,
				dir:       getDirection(vertexIdx, i, size),
				weight:    blocks[nr][nc],
			}

			prev.next = n
			prev = n
		}

		// remove dummy node
		adj[i] = adj[i].next
	}

	return adj
}

func DijkstraAdjLL(startAt int, adj []*gNode, size int) (paths, distances []int) {

	const maxStraightSteps = 3
	V := size * size
	var el *pqElement
	var from, to, d int
	q := priorityQueue{}
	paths = make([]int, V)
	distances = make([]int, V)

	for from = 0; from < V; from++ {
		distances[from] = -1
	}

	distances[startAt] = 0
	q.push(&pqElement{startAt, 0})

	var prevDirection direction
	var straightSteps int
	var node *gNode
	var dummyNode *gNode = &gNode{}

	for !q.isEmpty() {
		el = q.pop()
		from = el.key

		// if from == size*size-1 {
		// break
		// }

		dummyNode.next = adj[from]
		node = dummyNode
		for {
			node = node.next
			if node == nil {
				break
			}

			to = node.vertexIdx

			prevDirection, straightSteps = getStraightSteps(from, paths, size, prevDirection, maxStraightSteps)

			dir := getDirection(to, from, size)
			if dir != node.dir {
				fmt.Printf("\n calc dir: %#v, node dir: %#v", dir, node.dir)
			}

			if prevDirection == south && dir == north || prevDirection == north && dir == south || prevDirection == west && dir == east || prevDirection == east && dir == west {
				continue
			}

			if prevDirection == dir {
				straightSteps++
			}

			fmt.Printf("\n [%3d, %3d] = %5s -> %3d", from, to, directionsTranslation[prevDirection], straightSteps)
			fmt.Printf(" - prev D: %s, dir: %s", directionsTranslation[prevDirection], directionsTranslation[dir])

			if straightSteps > maxStraightSteps {
				fmt.Printf(" - skipping ")
				continue
			}

			d = distances[from] + node.weight

			if distances[to] == -1 {
				distances[to] = d
				q.push(&pqElement{to, d})
				paths[to] = from
			}

			if distances[to] > d {
				distances[to] = d
				q.updatePriority(to, d)
				paths[to] = from
			}

		}
		// }
	}

	paths[startAt] = -1
	return paths, distances
}

func getStraightSteps(pos int, paths []int, size int, prevDirection direction, maxStraightSteps int) (direction, int) {
	straightSteps := 0
	// fmt.Println()
	for pos > 0 {
		// fmt.Printf(" -> %3d -> %3d", pos, paths[pos])
		dir := getDirection(pos, paths[pos], size)
		if dir == prevDirection {
			straightSteps++
			if straightSteps > maxStraightSteps {
				break
			}
		} else {
			straightSteps = 0
			prevDirection = dir
			break
		}
		pos = paths[pos]
	}

	return prevDirection, straightSteps
}

func Ex2() {

}

// type pQueue []queElement

// // add x as element Len()
// func (q pQueue) Push(x any) {

// }

// // remove and return element Len() - 1.
// func (q pQueue) Pop() any {

// }

func DijkstraMatrix(startAt int, weights [][]int, size int) (paths, distances []int) {
	const maxStraightSteps = 3
	V := len(weights)
	var el *pqElement
	var v, w, d int
	q := priorityQueue{}
	paths = make([]int, V)
	distances = make([]int, V)

	for v = 0; v < V; v++ {
		distances[v] = -1
	}
	distances[startAt] = 0
	q.push(&pqElement{startAt, weights[startAt][startAt]})

	var prevDirection direction
	var straightSteps int

	for !q.isEmpty() {
		el = q.pop()
		v = el.key

		if v == size*size-1 {
			break
		}

		for w = 0; w < V; w++ {
			if weights[v][w] > 0 {

				prevDirection, straightSteps = getStraightSteps(v, paths, size, prevDirection, maxStraightSteps)

				dir := getDirection(w, v, size)
				if prevDirection == south && dir == north || prevDirection == north && dir == south || prevDirection == west && dir == east || prevDirection == east && dir == west {
					continue
				}

				if prevDirection == dir {
					straightSteps++
				}

				fmt.Printf("\n [%3d, %3d] = %5s -> %3d", v, w, directionsTranslation[prevDirection], straightSteps)
				fmt.Printf(" - prev D: %s, dir: %s", directionsTranslation[prevDirection], directionsTranslation[dir])

				if straightSteps > maxStraightSteps {
					fmt.Printf(" - skipping ")
					// q.enqueue(w, distances[w])
					continue
				}

				// ///////////////////

				d = distances[v] + weights[v][w]

				if distances[w] == -1 {
					distances[w] = d
					q.push(&pqElement{w, d})
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

func Dijkstra3(graph [][]int, n int, start int) ([]int, []int) {
	const maxStraightSteps = 3
	INFINITY := 99999999
	MAX := len(graph)
	cost := make([][]int, MAX)
	distances := make([]int, MAX)
	paths := make([]int, MAX)
	visited := make([]int, MAX)

	var count, mindistance, nextnode, i, j int
	// var straightSteps int
	// var prevDirection direction
	// size := MAX

	// Creating cost matrix
	for i = 0; i < n; i++ {
		cost[i] = make([]int, MAX)
		for j = 0; j < n; j++ {
			if graph[i][j] == 0 {
				cost[i][j] = INFINITY
			} else {
				cost[i][j] = graph[i][j]
			}
		}
	}

	for i = 0; i < n; i++ {
		distances[i] = cost[start][i]
		paths[i] = start
		visited[i] = 0
	}

	distances[start] = 0
	visited[start] = 1
	count = 1

	for count < n-1 {
		mindistance = INFINITY

		for i = 0; i < n; i++ {
			if distances[i] < mindistance && visited[i] == 0 {

				mindistance = distances[i]
				nextnode = i
			}
		}

		from := nextnode
		visited[from] = 1
		for to := 0; to < n; to++ {
			if visited[to] == 0 {

				// ///////////////////
				// prevDirection, straightSteps = getStraightSteps(from, paths, size, prevDirection, maxStraightSteps)

				// dir := getDirection(to, from, size)
				// if prevDirection == south && dir == north || prevDirection == north && dir == south || prevDirection == west && dir == east || prevDirection == east && dir == west {
				// 	// visited[from] = 0
				// 	continue
				// }

				// if prevDirection == dir {
				// 	straightSteps++
				// }

				// fmt.Printf("\n [%3d, %3d] = %5s -> %3d", from, to, directionsTranslation[prevDirection], straightSteps)
				// fmt.Printf(" - prev D: %s, dir: %s", directionsTranslation[prevDirection], directionsTranslation[dir])

				// if straightSteps > maxStraightSteps {
				// 	fmt.Printf(" - skipping ")
				// 	// visited[from] = 0
				// 	continue
				// }

				// ///////////////////

				if mindistance+cost[from][to] < distances[to] {
					distances[to] = mindistance + cost[from][to]
					paths[to] = from
				}
			}
		}
		count++
	}

	// Printing the distance
	for i = 0; i < n; i++ {
		if i != start {
			fmt.Printf("\nDistance from source to %d: %d", i, distances[i])
		}
	}

	paths[0] = -1

	return paths, distances
}
