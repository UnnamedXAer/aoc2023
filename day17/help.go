package day17

import (
	"golang.org/x/exp/slices"
)

func outsizeRange(ySize, xSize, row, col int) bool {

	if row < 0 || col < 0 || row == ySize || col == xSize {
		return true
	}

	return false
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

func buildGraph(blocks [][]int) [][]int {

	// var offsets = [...][2]int{
	offsets := map[string][2]int{
		"north": {-1, 0},
		"south": {1, 0},
		"west":  {0, -1},
		"east":  {0, 1},
	}
	size := len(blocks)
	bigSize := size * size
	graph := make([][]int, bigSize)

	for graphRow := 0; graphRow < bigSize; graphRow++ {
		graph[graphRow] = make([]int, size*size)
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {

			for _, o := range offsets {
				nr, nc := row+o[0], col+o[1]
				if nr > -1 && nr < size && nc > -1 && nc < size {
					graph[row*size+col][nr*size+nc] = blocks[nr][nc]
				}
			}

		}
	}
	return graph
}

// keys with lowest priority value first
type pqElement struct {
	key      queueE
	priority int
}

type priorityQueue struct {
	arr []*pqElement
}

func (q *priorityQueue) isEmpty() bool {
	return len(q.arr) == 0
}

func (q *priorityQueue) len() int {
	return len(q.arr)
}

// lowest priority first
func (q *priorityQueue) push(element *pqElement) {

	q.arr = append(q.arr, element)
}

func (q *priorityQueue) pop() *pqElement {
	if len(q.arr) == 0 {
		return nil
	}

	min := q.arr[0]
	minsIdx := 0
	for i, x := range q.arr {
		if x.priority < min.priority {
			minsIdx = i
			min = x
		}
	}

	q.arr[minsIdx] = nil
	q.arr = slices.Delete(q.arr, minsIdx, minsIdx+1)

	return min
}

func (q *priorityQueue) updatePriority(incomingKey queueE, priority int) {

	for _, x := range q.arr {
		if x.key.x == incomingKey.x && x.key.y == incomingKey.y {
			x.priority = priority
			return
		}
	}

	q.arr = append(q.arr, &pqElement{incomingKey, priority})
}

type queueE struct {
	x, y     int
	distance int
	dir      direction
}

type queue []queueE

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func (q queue) len() int {
	return len(q)
}

func (q *queue) push(element queueE) {

	*q = append(*q, element)
}

func (q *queue) pop() queueE {
	if len(*q) == 0 {
		panic("popping from empty queue")
	}

	first := (*q)[0]
	// q[0] = nil
	*q = slices.Delete(*q, 0, 0+1)

	return first
}

type qElementAny = any
type queueAny []qElementAny

func (q queueAny) isEmpty() bool {
	return len(q) == 0
}

func (q queueAny) len() int {
	return len(q)
}

func (q *queueAny) push(element qElementAny) {

	*q = append(*q, element)
}

func (q *queueAny) pop() qElementAny {
	if len(*q) == 0 {
		panic("popping from empty queue")
	}

	first := (*q)[0]
	// q[0] = nil
	*q = slices.Delete(*q, 0, 0+1)

	return first
}
