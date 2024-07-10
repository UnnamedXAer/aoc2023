package day23

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
)

const inputNameSuffix = ""

// const inputNameSuffix = "_t"
const inputName = "./day23/data" + inputNameSuffix + ".txt"

type point struct {
	y, x int
}

type World struct {
	w           [][]terrainType
	size        int
	entrancePos point
	exitPos     point
	adjacency   map[point][]point
}

type terrainType = byte

const (
	path      terrainType = '.'
	forrest   terrainType = '#'
	slopRight terrainType = '>'
	slopDown  terrainType = 'v'
	slopUp    terrainType = '^'
	slopLeft  terrainType = '<'
)

func (w World) String() string {
	s := strings.Builder{}

	s.WriteByte('\n')

	for _, line := range w.w {
		s.WriteString(fmt.Sprintf("\n%s", string(line)))
	}
	s.WriteString(fmt.Sprintf("\nentrance: %+v, exit: %+v", w.entrancePos, w.exitPos))

	return s.String()
}

func extractData() World {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	world := World{
		w: make([][]terrainType, 0, 141),
	}

	cnt := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		world.w = append(world.w, make([]terrainType, len(line)))
		copy(world.w[cnt], line)
		cnt++
	}

	help.IfErr(scanner.Err())

	world.size = cnt

	for i, b := range world.w[0] {
		if b == path {
			world.entrancePos = point{0, i}
			break
		}
	}

	for i, b := range world.w[cnt-1] {
		if b == path {
			world.exitPos = point{cnt - 1, i}
			break
		}
	}

	return world
}

func ExtractData() World {
	w := extractData()
	// fillAdjacencyMap(&w, false)
	return w
}

func Ex1(world World) {
	// fmt.Print(world)

	fillAdjacencyMap(&world, false)
	// we make "one step" from the entrance Pos to skip checking bounds for the world
	// the world is surrounded by forest so that should protect us from "index out of range"
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hike := make([]point, 1, world.size*5)
	hike[0] = world.entrancePos
	fmt.Print("\ncalcLongestHikeEx1")
	total, ok := calcLongestHikeEx1(false, world, entrance, 1, hike)
	if !ok {
		fmt.Printf("not ok")
	}
	fmt.Printf("\n\n total: %d", total)

}

func Ex1_1(world World) {

	fillAdjacencyMap(&world, false)
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hikeMap := make(map[point]bool, world.size*world.size)
	hikeMap[world.entrancePos] = true
	fmt.Print("\n\ncalcLongestHikeEx1HikeMap")
	total, ok := calcLongestHikeEx1HikeMap(false, world, entrance, 1, hikeMap)
	if !ok {
		fmt.Printf(" - not ok")
	}
	fmt.Printf("\n total hike map: %d", total)
}

func Ex1_2(world World) {

	fillAdjacencyMap(&world, false)
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hikeMap := make(map[point]bool, world.size*world.size)
	hikeMap[world.entrancePos] = true
	fmt.Print("\n\ncalcLongestHikeEx1HikeMapCameFrom")
	total, ok := calcLongestHikeEx1HikeMapCameFrom(false, world, entrance, world.entrancePos, 1, hikeMap)
	if !ok {
		fmt.Printf(" - not ok")
	}
	fmt.Printf("\n total hike map: %d", total)
}

func Ex1_3(world World) {

	fillAdjacencyMap(&world, false)
	entrance := point{y: world.entrancePos.y + 1, x: world.entrancePos.x}
	hikeMap := make(map[point]bool, world.size*world.size)
	hikeMap[world.entrancePos] = true
	fmt.Print("\n\ncalcLongestHikeMapWithAdjacencyMap")
	total, ok := calcLongestHikeMapWithAdjacencyMap(world, entrance, world.entrancePos, 1, hikeMap)
	if !ok {
		fmt.Printf(" - not ok")
	}

	fmt.Printf("\n total hike map: %d", total)
}

func calcLongestHikeMapWithAdjacencyMap(world World, pos point, cameFrom point, step int, hike map[point]bool) (int, bool) {

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
		if p == cameFrom {
			continue
		}

		tmpStep, ok := calcLongestHikeMapWithAdjacencyMap(world, p, pos, step, hike)
		// hike[p] = false
		delete(hike, p)
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	return max, true
	// if max > 0 {
	// 	return max, true
	// }

	// return 0, false
}

func calcLongestHikeEx1HikeMapCameFrom(canClimb bool, world World, pos, cameFrom point, step int, hike map[point]bool) (int, bool) {
	// no need for copy because caller won't see changes that are made outside "its length"
	// hike := _hike

	var tp terrainType = world.w[pos.y][pos.x]

	if tp == forrest {
		return 0, false
	}

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

	if canClimb || tp == path {

		for _, p := range directionsOffsets {
			p.y += pos.y
			p.x += pos.x

			if cameFrom == p {
				continue
			}
			if p.y == pos.y-1 && p.x == world.size-2 {
				continue
			}

			tmpStep, ok := calcLongestHikeEx1HikeMapCameFrom(canClimb, world, p, pos, step, hike)
			if ok && tmpStep > max {
				max = tmpStep
			}
			// hike[p] = false
			delete(hike, p)
		}
	} else {
		p := pos
		switch tp {
		case slopUp:
			p.y += -1
		case slopLeft:
			p.x += -1
		case slopDown:
			p.y += 1
		case slopRight:
			p.x += 1
		default:
			panic("unknown terrain: " + string(tp))
		}

		tmpStep, ok := calcLongestHikeEx1HikeMapCameFrom(canClimb, world, p, pos, step, hike)
		if ok && tmpStep > max {
			max = tmpStep
		}
		// hike[p] = false
		delete(hike, p)
	}

	if max > 0 {
		return max, true
	}

	return 0, false
}

func calcLongestHikeEx1HikeMap(canClimb bool, world World, pos point, step int, hike map[point]bool) (int, bool) {
	// no need for copy because caller won't see changes that are made outside "its length"
	// hike := _hike

	var tp terrainType = world.w[pos.y][pos.x]

	if tp == forrest {
		return 0, false
	}

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

	if canClimb || tp == path {

		for _, p := range directionsOffsets {
			p.y += pos.y
			p.x += pos.x

			tmpStep, ok := calcLongestHikeEx1HikeMap(canClimb, world, p, step, hike)
			if ok && tmpStep > max {
				max = tmpStep
			}
			delete(hike, p)
		}
	} else {
		p := pos
		switch tp {
		case slopUp:
			p.y += -1
		case slopLeft:
			p.x += -1
		case slopDown:
			p.y += 1
		case slopRight:
			p.x += 1
		default:
			panic("unknown terrain: " + string(tp))
		}

		tmpStep, ok := calcLongestHikeEx1HikeMap(canClimb, world, p, step, hike)
		if ok && tmpStep > max {
			max = tmpStep
		}
		delete(hike, p)
	}

	if max > 0 {
		return max, true
	}

	return 0, false
}

func calcLongestHikeEx1(canClimb bool, world World, pos point, step int, hike []point) (int, bool) {
	// no need for copy because caller won't see changes that are made outside "its length"
	// hike := _hike

	var tp terrainType = world.w[pos.y][pos.x]

	if tp == forrest {
		return 0, false
	}

	if pos == world.exitPos {
		return step, true
	}

	if containsPoint(hike, pos) {
		return 0, false
	}

	step++
	hike = append(hike, pos)

	max := 0

	if canClimb || tp == path {

		for _, p := range directionsOffsets {
			p.y += pos.y
			p.x += pos.x

			tmpStep, ok := calcLongestHikeEx1(canClimb, world, p, step, hike)
			if ok && tmpStep > max {
				max = tmpStep
			}
		}
	} else {
		p := pos
		switch tp {
		case slopUp:
			p.y += -1
		case slopLeft:
			p.x += -1
		case slopDown:
			p.y += 1
		case slopRight:
			p.x += 1
		default:
			panic("unknown terrain: " + string(tp))
		}

		tmpStep, ok := calcLongestHikeEx1(canClimb, world, p, step, hike)
		if ok && tmpStep > max {
			max = tmpStep
		}
	}

	if max > 0 {
		return max, true
	}

	return 0, false
}

func containsPoint(hike []point, p point) bool {
	// return slices.Contains(hike, p)

	for i := len(hike) - 1; i >= 0; i-- {
		if hike[i] == p {
			return true
		}
	}
	return false
}

var directionsOffsets = [4]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
