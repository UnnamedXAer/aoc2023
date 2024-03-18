package day23

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/constraints"
)

func Ex2(world World) {
	// world := extractData()
	fillAdjacencyMap(&world, true)

	cnt := 0

	for _, v := range world.w {
		for _, v := range v {
			if v != forrest {
				cnt++
			}
		}
	}
	world.steppableCnt = cnt

	// we make "one step" from the entrance Pos to skip checking bounds for the world
	// the world is surrounded by forest so that should protect us from "index out of range"
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hike := make(map[point]bool, world.size*world.size)
	hike[world.entrancePos] = true

	paths := make(map[point][]point)
	paths[entrance] = []point{world.entrancePos}

	total2, ok2 := calcLongestHikeWithAdjacencyMapEx2(world, entrance, world.entrancePos, 1, hike)
	if !ok2 {
		fmt.Printf("\n\n----not ok2")
	}

	fmt.Printf("\n\n total: %d", total2)
	// 5030 - too low
}

func Ex2_1(world World) {
	// world := extractData()
	fillAdjacencyMap(&world, true)

	cnt := 0

	for _, v := range world.w {
		for _, v := range v {
			if v != forrest {
				cnt++
			}
		}
	}
	world.steppableCnt = cnt

	// we make "one step" from the entrance Pos to skip checking bounds for the world
	// the world is surrounded by forest so that should protect us from "index out of range"
	// entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hike := make([]uint16, 0, cnt/3)

	total2, ok2 := calcLongestHikeWithAdjacencyMapEx2_1(world, world.entrancePos, world.entrancePos, 0, hike)
	if !ok2 {
		fmt.Printf("\n\n----not ok2")
	}

	fmt.Printf("\n\n total: %d", total2)
	// 5030 - too low
	// 5246 - too low
}

func calcLongestHikeWithAdjacencyMapEx2_1(world World, pos point, cameFrom point, step int, hike []uint16) (int, bool) {

	if pos == world.exitPos {
		return step, true
	}

	if step == world.steppableCnt {
		return 0, false
	}

	idx := uint16(pos.y*world.size + pos.x)

	if contains(hike, idx) {
		return 0, false
	}

	step++
	hike = append(hike, idx)

	max := 0

	for _, p := range world.adjacency[pos] {
		if p == cameFrom {
			continue
		}

		if p.y == pos.y-1 && p.x == world.size-2 {
			continue
		}

		if p.y == pos.y-1 && p.x == 1 {
			continue
		}

		tmpStep, ok := calcLongestHikeWithAdjacencyMapEx2_1(world, p, pos, step, hike)
		// hike[p] = false
		// delete(hike, p)
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	if max > 0 {
		return max, true
	}

	return 0, false
}

func calcLongestHikeWithAdjacencyMapEx2(world World, pos point, cameFrom point, step int, hike map[point]bool) (int, bool) {

	if pos == world.exitPos {
		return step, true
	}

	if step == world.steppableCnt {
		return 0, false
	}

	if hike[pos] {
		return 0, false
	}

	step++
	hike[pos] = true

	max := 0

	for _, p := range world.adjacency[pos] {
		if p == cameFrom {
			continue
		}

		if p.y == pos.y-1 && p.x == world.size-2 {
			continue
		}

		if p.y == pos.y-1 && p.x == 1 {
			continue
		}

		tmpStep, ok := calcLongestHikeWithAdjacencyMapEx2(world, p, pos, step, hike)
		// hike[p] = false
		delete(hike, p)
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	if max > 0 {
		return max, true
	}

	return 0, false
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

func contains[T constraints.Ordered](s []T, x T) bool {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == x {
			return true
		}
	}
	return false
}
