package day23

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
)

func Ex2() {
	world := extractData()
	// fmt.Print(world)
	fillAdjacencyMap(&world)

	// for p, neighbours := range world.adjacency {
	// 	fmt.Printf("\n%+v:", p)
	// 	for _, neighbour := range neighbours {
	// 		fmt.Printf("\n\t%+v:", neighbour)
	// 	}
	// }

	// we make "one step" from the entrance Pos to skip checking bounds for the world
	// the world is surrounded by forest so that should protect us from "index out of range"
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hike := make(map[point]bool, world.size*world.size)
	hike[world.entrancePos] = true
	// hike := make([]point, 1, world.size*5)
	// hike[0] = world.entrancePos
	// total, ok := calcLongestHike(true, world, entrance, 1, hike)
	// if !ok {
	// 	fmt.Printf("not ok")
	// }

	// fmt.Printf("\n\n total: %d", total)

	total2, ok2 := calcLongestHikeWithConnections(world, entrance, 1, hike)
	if !ok2 {
		fmt.Printf("\n\n----not ok2")
	}

	fmt.Printf("\n\n total: %d", total2)
	// 5030 - too low
}

func calcLongestHikeWithConnections(world World, pos point, step int, hike map[point]bool) (int, bool) {
	// no need for copy because caller won't see changes that are made outside "its length"
	// hike := _hike

	// var tp terrainType = world.w[pos.y][pos.x]

	// if tp == forrest {
	// 	return 0, false
	// }

	if pos == world.exitPos {
		return step, true
	}

	// if contains(hike, pos) {
	if hike[pos] {
		return 0, false
	}

	step++
	// hike = append(hike, pos)
	hike[pos] = true

	max := 0

	for _, p := range world.adjacency[pos] {
		tmpStep, ok := calcLongestHikeWithConnections(world, p, step, hike)
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	if max > 0 {
		return max, true
	}

	return 0, false
}

func fillAdjacencyMap(world *World) {
	adjacency := make(map[point][]point, (world.size*world.size)/2)
	startPos := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	// close exits so we do not have to check for index out of range
	world.w[world.entrancePos.y][world.entrancePos.x] = forrest
	world.w[world.exitPos.y][world.exitPos.x] = forrest
	addConnections(adjacency, world.w, startPos)
	// add connection to exit point
	adjacency[point{y: world.exitPos.y - 1, x: world.exitPos.x}] = append(adjacency[point{y: world.exitPos.y - 1, x: world.exitPos.x}], world.exitPos)
	world.w[world.entrancePos.y][world.entrancePos.x] = path
	world.w[world.exitPos.y][world.exitPos.x] = path
	world.adjacency = adjacency
}
func addConnections(adjacency map[point][]point, w [][]terrainType, p point) {

	q := help.NewQAny[point]()
	q.Push(p)

	for !q.IsEmpty() {
		p = q.Pop()
		if _, ok := adjacency[p]; ok {
			continue
		}

		for _, next := range directionsOffsets {
			next.y += p.y
			next.x += p.x

			if w[next.y][next.x] != forrest {
				adjacency[p] = append(adjacency[p], next)
				q.Push(next)
			}

		}
	}
}
