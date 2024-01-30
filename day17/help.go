package day17

import (
	"golang.org/x/exp/slices"
)

func outsizeRange(size, row, col int) bool {

	if row < 0 || col < 0 || row == size || col == size {
		return true
	}

	return false
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
	key      int
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

func (q *priorityQueue) updatePriority(key int, priority int) {

	for _, x := range q.arr {
		if x.key < key {
			x.priority = priority
			return
		}
	}

	q.arr = append(q.arr, &pqElement{key, priority})
}

type queueE struct {
	x, y int
}
type queue []*queueE

func (q queue) isEmpty() bool {
	return len(q) == 0
}

func (q queue) len() int {
	return len(q)
}

func (q queue) push(element *queueE) {

	q = append(q, element)
}

func (q queue) pop() *queueE {
	if len(q) == 0 {
		return nil
	}

	first := q[0]
	q[0] = nil
	q = slices.Delete(q, 0, 0+1)

	return first
}
