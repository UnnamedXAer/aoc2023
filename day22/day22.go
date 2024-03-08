package day22

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

// const inputNameSuffix= ""
const inputNameSuffix = "_t"
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

func (b brick) getPosStartAndDelta(p perspective) (s int, d int) {
	// x := b.p1.x
	// if b.p2.x < x {
	// 	x = b.p2.x
	// }
	// dx := abs(b.p1.x-b.p2.x) + 1

	v1 := b.val(p, 1)
	v2 := b.val(p, 2)

	if v1 < v2 {
		s = v1
		d = v2 - v1
	} else {
		s = v2
		d = v1 - v2
	}

	d += 1
	// d = abs(v1-v2) + 1

	return s, d
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
	printBricksFromPerspective(bricks, pX)
}

type perspective byte

const (
	pX perspective = 'x'
	pY perspective = 'y'
	pZ perspective = 'z'
)

func printBricksFromPerspective(b []*brick, p perspective) {

	maxX := 0
	maxZ := 0
	for _, b := range b {
		if b.p1.x > maxX {
			maxX = b.p1.x
		}
		if b.p2.x > maxX {
			maxX = b.p2.x
		}

		if b.p1.z > maxZ {
			maxZ = b.p1.z
		}
		if b.p2.z > maxZ {
			maxZ = b.p2.z
		}
	}

	maxX += 1

	fmt.Println()

	for i := 0; i < maxX; i++ {
		fmt.Print(i % 10)
	}

	switch p {
	case pX:
		for k := len(b) - 1; k >= 0; k-- {
			b := b[k]
			// x := b.p1.x
			// if b.p2.x < x {
			// 	x = b.p2.x
			// }
			// dx := abs(b.p1.x-b.p2.x) + 1
			x, dx := b.getPosStartAndDelta(p)

			line := make([]byte, maxX)

			for i := 0; i < maxX; i++ {
				if i >= x && i < x+dx {
					line[i] = byte(b.id) + 'A'
				} else {
					line[i] = '.'
				}
			}

			fmt.Printf("\n%2d %s", k, line)
		}

		fmt.Println()

	default:
		panic("not implemented perspective: " + string(p))
	}
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
