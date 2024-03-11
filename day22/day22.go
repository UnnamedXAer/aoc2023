package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"

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
	idLetter := b.idToLetter()
	return fmt.Sprintf("{%+v, %+v | %4d | %s}", b.p1, b.p2, b.id, string(idLetter))
}

func (b brick) idToLetter() byte {
	base := b.id % (58 - 6)
	if base > 25 {
		base += 6
	}
	idLetter := byte(base) + 'A'
	return idLetter
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

func (b brick) clone() *brick {
	return &b
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

	if b.p1.x > b.p2.x ||
		b.p1.y > b.p2.y ||
		b.p1.z > b.p2.z {
		fmt.Printf("\n p1 is greater then p2: %s", b)
	}
	return &b
}

func Ex1() {
	bricks := extractData()
	// for _, b := range bricks {
	// 	fmt.Printf("\n%v", b)
	// }

	fmt.Println()

	for _, b := range bricks {
		calcVolume(b)
	}

	// fmt.Println()
	// printBricksFromPerspective(bricks, pX, pZ)
	// printBricksFromPerspective(bricks, pY, pZ)
	// printBricksFromPerspective(bricks, pZ, pX)
	// printBricksFromPerspective(bricks, pZ, pY)
	// printBricksFromPerspective(bricks, pX, pY)

	bricks = fallBricks(bricks)

	supports, supportedby := determineWhatSupportWhat(bricks)
	total := calcDisintegrable(bricks, supports, supportedby)

	fmt.Printf("\n\n Total: %d", total)
	// 1330 - too high
}

func determineWhatSupportWhat(bricks []*brick) (map[*brick][]*brick, map[*brick][]*brick) {
	supports := make(map[*brick][]*brick, len(bricks))
	supportedBy := make(map[*brick][]*brick, len(bricks))

	for _, b1 := range bricks {
		for _, b2 := range bricks {
			if isSupporting(b1, b2) {
				supports[b1] = append(supports[b1], b2)
				supportedBy[b2] = append(supportedBy[b2], b1)
			}
		}
	}

	// fmt.Println("is supporting: ")
	// for _, b1 := range bricks {
	// 	list := supports[b1]
	// 	fmt.Printf("\n%s: ", string(b1.idToLetter()))
	// 	for _, b2 := range list {
	// 		fmt.Printf(", %s", string(b2.idToLetter()))
	// 	}
	// }

	// fmt.Println()
	// fmt.Println("is supported by: ")

	// for _, b1 := range bricks {
	// 	list := supportedBy[b1]
	// 	fmt.Printf("\n%s: ", string(b1.idToLetter()))
	// 	for _, b2 := range list {
	// 		fmt.Printf(", %s", string(b2.idToLetter()))
	// 	}
	// }

	return supports, supportedBy
}

func calcDisintegrable(bricks []*brick, supports map[*brick][]*brick, supportedBy map[*brick][]*brick) int {
	disintegrable := map[*brick]bool{}
	fmt.Println()
	for _, b1 := range bricks {
		disintegrable[b1] = false
		list := supports[b1]

		if len(list) == 0 {
			disintegrable[b1] = true
			continue
		}

		for _, b2 := range list {
			if len(supportedBy[b2]) > 1 {
				disintegrable[b1] = true
			}
		}
	}

	cnt := 0
	for _, v := range disintegrable {
		if v {
			cnt++
		}
	}
	return cnt
}

func fallBricks(bricks []*brick) []*brick {

	sortedBricks := make([]*brick, len(bricks))
	for i, b := range bricks {
		sortedBricks[i] = b.clone()
	}

	slices.SortFunc(sortedBricks, func(a, b *brick) int {
		return a.smallerValue(pZ) - b.smallerValue(pZ)
	})

	fmt.Printf("\n")

	for {
		cnt := 0
		for _, b := range sortedBricks {

			for canFall(sortedBricks, b) {
				cnt++
				b.p1.z -= 1
				b.p2.z -= 1
			}
		}
		fmt.Printf("\n falls: %d", cnt)
		if cnt == 0 {
			break
		}
	}

	return sortedBricks
}

func isSupporting(b1, b2 *brick) bool {
	if b1.greaterValue(pZ) != b2.smallerValue(pZ)-1 {
		return false
	}

	b1px1 := b1.smallerValue(pX)
	b1px2 := b1.greaterValue(pX)

	b1py1 := b1.smallerValue(pY)
	b1py2 := b1.greaterValue(pY)

	b2x1 := b2.smallerValue(pX)
	b2x2 := b2.greaterValue(pX)

	b2y1 := b2.smallerValue(pY)
	b2y2 := b2.greaterValue(pY)

	noOverlap := false
	tmp := false

	tmp = b1px1 > b2x2
	noOverlap = noOverlap || tmp
	tmp = b1px2 < b2x1
	noOverlap = noOverlap || tmp
	tmp = b1py2 < b2y1
	noOverlap = noOverlap || tmp
	tmp = b1py1 > b2y2
	noOverlap = noOverlap || tmp

	return !noOverlap
}

func canFall(bricks []*brick, b *brick) bool {
	z := b.smallerValue(pZ)
	if z <= 1 {
		return false
	}

	bpx1 := b.smallerValue(pX)
	bpx2 := b.greaterValue(pX)

	bpy1 := b.smallerValue(pY)
	bpy2 := b.greaterValue(pY)

	bId := b.id

	for _, currBrick := range bricks {
		if currBrick.id == bId {
			continue
		}

		currBrickZ := currBrick.greaterValue(pZ)
		if !(currBrickZ+1 == z) {
			continue
		}

		// does bricks overlap on plane x/y
		cbpx1 := currBrick.smallerValue(pX)
		cbpx2 := currBrick.greaterValue(pX)

		cbpy1 := currBrick.smallerValue(pY)
		cbpy2 := currBrick.greaterValue(pY)

		noOverlap := false
		tmp := false

		tmp = bpx1 > cbpx2
		noOverlap = noOverlap || tmp
		tmp = bpx2 < cbpx1
		noOverlap = noOverlap || tmp
		tmp = bpy2 < cbpy1
		noOverlap = noOverlap || tmp
		tmp = bpy1 > cbpy2
		noOverlap = noOverlap || tmp

		if !noOverlap {
			return false
		}
	}

	return true
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
