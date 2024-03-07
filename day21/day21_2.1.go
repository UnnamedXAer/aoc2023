package day21

import (
	"fmt"
	"math"

	"github.com/unnamedxaer/aoc2023/help"
)

// rewrite from:
// https://github.com/Nebula83/aoc2023/blob/master/day-21/sol.py

type a_point = [2]int

// possiblePoints returns `next` function.
// The `next` function returns next point for the given point (n,w,s,e) unless it is rock. The `next` function returns 3 values, i.e. next point, bool if point is reachable, and bool indicating if points all points were already generated.
func possiblePoints(p a_point, garden [][]byte) func() (a_point, bool, bool) {
	directions := []a_point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	i := -1
	return func() (a_point, bool, bool) {
		i++
		if i >= len(directions) {
			return a_point{}, false, false // all directions were exhausted.
		}
		d := directions[i]
		y, x := (p[0] + d[0]), (p[1] + d[1])
		newPoint := a_point{y, x}
		y = y % 131
		x = x % 131
		if y < 0 {
			y = 131 + y
		}
		if x < 0 {
			x = 131 + x
		}

		if garden[x][y] == ROCK {
			return a_point{}, false, true // we cannot step here but you should check next direction
		}
		return newPoint, true, true // we can step on this point
	}
}

type qItem struct {
	p    a_point
	dist int
}

func bfs(p a_point, garden [][]byte, maxDist int) map[int]int {

	tiles := map[int]int{}
	visited := map[a_point]bool{}

	q := help.NewQAny[qItem](10)

	q.Push(qItem{p: p, dist: 0})

	for !q.IsEmpty() {
		item := q.Pop()

		if item.dist == (maxDist + 1) {
			continue
		}
		if _, ok := visited[item.p]; ok {
			continue
		}

		p := item.p

		tiles[item.dist] += 1
		visited[p] = true

		next := possiblePoints(p, garden)
		for {
			newPoint, steppable, ok := next()
			if !ok {
				break
			}
			if !steppable {
				continue
			}

			q.Push(qItem{newPoint, item.dist + 1})
		}
	}

	return tiles
}

func calculatePossibleSpots(start a_point, garden [][]byte, maxStep int) int {
	tiles := bfs(start, garden, maxStep)

	amount := 0
	isEven := maxStep % 2
	for dist, item := range tiles {
		if dist%2 == isEven {
			amount += item
		}
	}

	return amount
}

func quad(y []int, n int) int {
	a := (y[2] - (2 * y[1]) + y[0]) / 2
	b := y[1] - y[0] - a
	c := y[0]

	out := (a * int(math.Pow(float64(n), 2))) + (b * n) + c

	return out
}

func Part1() {

	garden, startPos := extractData()

	start := a_point{startPos.y, startPos.x}

	total := calculatePossibleSpots(start, garden, 64)
	fmt.Printf("\n\nTotal: %d", total)
}

func Part2() {

	garden, startPos := extractData()
	start := a_point{startPos.y, startPos.x}
	size := len(garden)
	edge := size / 2
	goal := 26501365

	y := make([]int, 3)
	for i := 0; i < 3; i++ {
		y[i] = calculatePossibleSpots(start, garden, edge+i*size)
	}

	total := quad(y, ((goal - edge) / size))
	fmt.Printf("\n\nTotal: %d", total)
}
