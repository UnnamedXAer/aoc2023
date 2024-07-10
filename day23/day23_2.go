package day23

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
)

func Ex2(world World) {

	fillAdjacencyMap(&world, true)

	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hikeMap := make(map[point]bool, world.size*world.size)
	hikeMap[world.entrancePos] = true
	fmt.Print("\n\ncalcLongestHikeMapWithAdjacencyMapEx2")
	total, ok := calcLongestHikeMapWithAdjacencyMapEx2(world.exitPos, world.adjacency, entrance, world.entrancePos, 1, hikeMap)
	if !ok {
		fmt.Printf("not ok")
	}

	fmt.Printf("\n total hike map: %d", total)
}

func calcLongestHikeMapWithAdjacencyMapEx2(
	exitPos point,
	adjacency map[point][]point,
	pos point,
	cameFrom point,
	step int,
	hike map[point]bool,
) (int, bool) {

	if pos == exitPos {
		return step, true
	}

	if hike[pos] {
		return 0, false
	}

	step++
	hike[pos] = true

	max := 0

	var p point
	var ok bool
	var tmpStep int

	for _, p = range adjacency[pos] {
		if p == cameFrom {
			continue
		}

		tmpStep, ok = calcLongestHikeMapWithAdjacencyMapEx2(exitPos, adjacency, p, pos, step, hike)

		// delete(hike, p)
		hike[p] = false
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	return max, true
}

func fillAdjacencyMap(world *World, canClimb bool) {
	adjacency := make(map[point][]point, (world.size*world.size)/2)
	startPos := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	// close exits so we do not have to check for index out of range
	world.w[world.entrancePos.y][world.entrancePos.x] = forrest
	world.w[world.exitPos.y][world.exitPos.x] = forrest
	adjacency[world.entrancePos] = []point{startPos}
	addConnections(adjacency, world.w, startPos, canClimb)
	// add connection to exit point
	adjacency[point{y: world.exitPos.y - 1, x: world.exitPos.x}] = append(adjacency[point{y: world.exitPos.y - 1, x: world.exitPos.x}], world.exitPos)
	world.w[world.entrancePos.y][world.entrancePos.x] = path
	world.w[world.exitPos.y][world.exitPos.x] = path
	world.adjacency = adjacency
}

func addConnections(adjacency map[point][]point, w [][]terrainType, p point, canClimb bool) {

	q := help.NewQAny[point]()
	// q := help.NewStack[point]()
	q.Push(p)

	for !q.IsEmpty() {
		p = q.Pop()
		if _, ok := adjacency[p]; ok {
			continue
		}

		var tp terrainType = w[p.y][p.x]
		if canClimb || tp == path {
			for _, next := range directionsOffsets {
				next.y += p.y
				next.x += p.x

				if w[next.y][next.x] != forrest {
					adjacency[p] = append(adjacency[p], next)
					q.Push(next)
				}

			}
		} else {
			next := p
			switch tp {
			case slopUp:
				next.y += -1
			case slopLeft:
				next.x += -1
			case slopDown:
				next.y += 1
			case slopRight:
				next.x += 1
			default:
				panic("unknown terrain: " + string(tp))
			}

			if w[next.y][next.x] != forrest {
				adjacency[p] = append(adjacency[p], next)
				q.Push(next)
			}
		}
	}
}
