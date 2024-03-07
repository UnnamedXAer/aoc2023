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
	p1, p2 vector
}

func (b brick) String() string {
	return fmt.Sprintf("{%+v, %+v}", b.p1, b.p2)
}

func extractData() []brick {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	bricks := make([]brick, 0, 1400)

	for scanner.Scan() {
		line := scanner.Bytes()
		b := extractBrick(line)
		bricks = append(bricks, b)
	}

	help.IfErr(scanner.Err())

	return bricks
}

func extractBrick(line []byte) brick {

	vectors := bytes.Split(line, []byte{'~'})

	b := brick{}

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

	return b
}

func Ex1() {
	bricks := extractData()
	for _, b := range bricks {
		fmt.Printf("\n%v", b)
	}
}
