package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

const inputNameSuffix = ""

// const inputNameSuffix = "_t"
const inputName = "./day22/data" + inputNameSuffix + ".txt"

type vector struct {
	x, y, z int
}
type brick struct {
	id     int
	p1, p2 vector
}

func (b brick) String() string {
	return fmt.Sprintf("{%+v, %+v | %3d}", b.p1, b.p2, b.id)
}

func (b brick) val(p perspective, pn int) (v int) {
	switch p {
	case pX:
		if pn == 1 {
			return b.p1.x
		}
		return b.p2.x
	case pY:
		if pn == 1 {
			return b.p1.y
		}
		return b.p2.y
	case pZ:
		if pn == 1 {
			return b.p1.z
		}
		return b.p2.z
	}

	return v
}

func (b brick) smallerValue(p perspective) int {
	return min(b.val(p, 1), b.val(p, 2))
}

func (b brick) greaterValue(p perspective) int {
	return max(b.val(p, 1), b.val(p, 2))
}

func extractData() []*brick {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	bricks := make([]*brick, 0, 1400)

	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		b := extractBrick(line, i)
		bricks = append(bricks, b)
		i++
	}

	help.IfErr(scanner.Err())

	return bricks
}

func extractBrick(line []byte, id int) *brick {

	vectors := bytes.Split(line, []byte{'~'})

	b := brick{id: id}

	for vIdx, vector := range vectors {
		coords := bytes.Split(vector, []byte{','})

		for cIdx, coord := range coords {
			multiplier := 1
			value := 0
			for i := len(coord) - 1; i >= 0; i-- {

				value += int(coord[i]-'0') * multiplier
				multiplier++
			}

			if vIdx == 0 {
				switch cIdx {
				case 0:
					b.p1.x = value
				case 1:
					b.p1.y = value
				case 2:
					b.p1.z = value
				}
			} else {
				switch cIdx {
				case 0:
					b.p2.x = value
				case 1:
					b.p2.y = value
				case 2:
					b.p2.z = value
				}
			}
		}
	}

	return &b
}

func Ex1() {
	bricks := extractData()
	for _, b := range bricks {
		fmt.Printf("\n%v", b)
	}

	fmt.Println()

	for _, b := range bricks {
		calcVolume(b)
	}

	fmt.Println()
	printBricksFromPerspective(bricks, pX, pZ)
	printBricksFromPerspective(bricks, pY, pZ)
	printBricksFromPerspective(bricks, pZ, pX)
	printBricksFromPerspective(bricks, pZ, pY)
	printBricksFromPerspective(bricks, pX, pY)
}

type perspective byte

const (
	pX perspective = 'X'
	pY perspective = 'Y'
	pZ perspective = 'Z'
)

func printBricksFromPerspective(bricks []*brick, ph perspective, pv perspective) {
	if ph == pv {
		panic("\nperspectives cannot be the same")
	}

	maxPh, maxPv := getMaxForPerspectives(bricks, ph, pv)

	var currentClosestBrick *brick
	picture := make([][]byte, 0, maxPv)

	// fill 2D grid with blocks
	endIdx := 0
	if pv == pZ {
		endIdx = 1
	}
	for v := maxPv - 1; v >= endIdx; v-- {
		line := make([]byte, maxPh)

		for h := 0; h < maxPh; h++ {
			// most likely it would be more efficient the other way around to loop through
			// bricks and put them on the grid, and skip where the block is behind already
			// placed block in a given point
			currentClosestBrick, _ = findClosestBrick(bricks, ph, pv, h, v)
			if currentClosestBrick == nil {
				line[h] = '.'
			} else {
				// display letters A-Za-z - and yes, there will be duplicates if we have more bricks
				// then available letters
				base := currentClosestBrick.id % (58 - 6)
				if base > 25 {
					base += 6
				}
				line[h] = byte(base) + 'A'
			}
		}

		picture = append(picture, line)
	}

	printBricksPicture(picture, ph, pv, maxPh, maxPv)
}

func printBricksPicture(picture [][]byte, ph, pv perspective, maxPh, maxPv int) {
	fmt.Println()

	for i := 0; i < (maxPh/2)+1; i++ {
		fmt.Print(" ")
	}
	fmt.Print(string(ph))
	fmt.Print("\n ")
	for i := 0; i < maxPh; i++ {
		fmt.Print(i % 10)
	}

	for i, v := range picture {
		fmt.Printf("\n %s %d", string(v), (maxPv - (i + 1)))
		if i == maxPv/2 {
			fmt.Printf(" %s", string(pv))
		}
	}
	fmt.Print("\n ")
	for i := 0; i < maxPh; i++ {
		fmt.Print("-")
	}
	if pv == pZ {
		fmt.Println(" 0")
	} else {
		fmt.Println(" -")

	}
}

func findClosestBrick(bricks []*brick, pH, pV perspective, idxH, idxV int) (*brick, int) {
	bricksCnt := len(bricks)
	if bricksCnt == 0 {
		return nil, math.MaxInt
	}

	var p perspective

	switch pH {
	case pX:
		if pV == pZ {
			p = pY
		} else if pV == pY {
			p = pZ
		}
	case pY:
		if pV == pZ {
			p = pX
		} else if pV == pX {
			p = pZ
		}
	case pZ:
		if pV == pY {
			p = pX
		} else if pV == pX {
			p = pY
		}
	}

	var closestBrick *brick
	closest := math.MaxInt
	for _, b := range bricks {
		value := b.smallerValue(p)
		if value >= closest {
			continue
		}

		// check if projection of current brick includes point (x,z)
		value1 := b.smallerValue(pH)
		value2 := b.greaterValue(pH)
		if !(idxH >= value1 && idxH <= value2) {
			continue
		}
		value1 = b.smallerValue(pV)
		value2 = b.greaterValue(pV)
		if !(idxV >= value1 && idxV <= value2) {
			continue
		}

		closest = value
		closestBrick = b
	}

	return closestBrick, closest
}

func getMaxForPerspectives(bricks []*brick, ph, pv perspective) (int, int) {
	maxX := 0
	maxY := 0
	maxZ := 0
	maxPh := 0
	maxPv := 0
	for _, b := range bricks {
		if b.p1.x > maxX {
			maxX = b.p1.x
		}
		if b.p2.x > maxX {
			maxX = b.p2.x
		}

		if b.p1.y > maxY {
			maxY = b.p1.y
		}
		if b.p2.y > maxY {
			maxY = b.p2.y
		}

		if b.p1.z > maxZ {
			maxZ = b.p1.z
		}
		if b.p2.z > maxZ {
			maxZ = b.p2.z
		}
	}

	maxX += 1
	maxZ += 1
	maxY += 1

	switch pv {
	case pX:
		maxPv = maxX
	case pY:
		maxPv = maxY
	case pZ:
		maxPv = maxZ
	default:
		panic("unknown perspective: " + string(pv))
	}

	switch ph {
	case pX:
		maxPh = maxX
	case pY:
		maxPh = maxY
	case pZ:
		maxPh = maxZ
	}

	return maxPh, maxPv
}

func calcVolume(b *brick) {

	x := abs(b.p2.x-b.p1.x) + 1
	y := abs(b.p2.y-b.p1.y) + 1
	z := abs(b.p2.z-b.p1.z) + 1

	v := x * y * z

	fmt.Printf("\nb: %20v - %3d", b, v)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
