package day15

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]byte {

	splitFn := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		commaidx := bytes.IndexByte(data, ',')
		if commaidx > 0 {
			// we need to return the next position
			buffer := data[:commaidx]
			return commaidx + 1, buffer, nil
		}

		// if we are at the end of the string, just return the entire buffer
		if atEOF {
			// but only do that when there is some data. If not, this might mean
			// that we've reached the end of our input CSV string
			if len(data) > 0 {
				return len(data), data, nil
			}
		}

		// when 0, nil, nil is returned, this is a signal to the interface to read
		// more data in from the input reader. In this case, this input is our
		// string reader and this pretty much will never occur.
		return 0, nil, nil
	}

	f, err := os.Open("./day15/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(splitFn)

	input := make([][]byte, 0, 4001)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		input = append(input, line)
	}

	help.IfErr(scanner.Err())

	return input
}

func Ex1() {
	total := 0
	input := extractData()
	for _, str := range input {
		total += hash([]byte(str))
	}

	fmt.Printf("\n\n Total: %d", total)
}

func hash(str []byte) int {
	currValue := 0
	for _, b := range str {
		currValue += int(b)
		currValue *= 17
		currValue %= 256
	}

	return currValue
}

const remove byte = '-'
const upsert byte = '='

type lens struct {
	label    string
	opt      byte
	focalLen int
}

type theBox struct {
	lenses []lens
}

func Ex2() {

	total := 0
	input := extractData()
	lenses := make([]lens, len(input))

	boxes := make([]*theBox, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = &theBox{
			lenses: make([]lens, 0, 4000/256),
		}
	}

	for i, str := range input {
		l := parseStep(str)
		lenses[i] = l
	}

	for _, l := range lenses {
		boxIdx := hash([]byte(l.label))
		// fmt.Printf("\n [%s %d] => %d", l.label, l.focalLen, boxIdx)
		switch l.opt {
		case remove:
			removeLens(boxes, l, boxIdx)
		case upsert:
			upsertLens(boxes, l, boxIdx)
		}
	}

	for i, b := range boxes {
		if len(b.lenses) == 0 {
			continue
		}
		focalLen := 0
		// fmt.Printf("\nBox %4d:", i)
		for k, l := range b.lenses {
			lpopwer := calcFocusingPower(l.focalLen, i, k)
			// fmt.Printf(" [%s %d -> %2d]", l.label, l.focalLen, lpopwer)
			focalLen += lpopwer
		}
		// fmt.Printf(" -> %3d", focalLen)
		total += focalLen

	}

	fmt.Printf("\n\n Total: %d", total)
}

// removeLens removes lens from a box if present.
// It returns box idx and boolean indicating if anything was removed
func removeLens(boxes []*theBox, l lens, boxIdx int) (int, bool) {

	lenses := boxes[boxIdx].lenses
	for i, bl := range lenses {
		if bl.label == l.label {
			for k := i + 1; k < len(lenses); k++ {
				lenses[k-1] = lenses[k]
			}
			boxes[boxIdx].lenses = lenses[:len(lenses)-1]
			return i, true
		}
	}

	return 0, false
}

// upsertLens replaces lens with the same label if any exist otherwise inserts it at the end.
// It returns idx of replaced or inserted element and boolean indicating if anything was replaced.
func upsertLens(boxes []*theBox, l lens, boxIdx int) (int, bool) {

	lenses := boxes[boxIdx].lenses
	for i, bl := range lenses {
		if bl.label == l.label {
			boxes[boxIdx].lenses[i] = l
			return i, true
		}
	}

	boxes[boxIdx].lenses = append(boxes[boxIdx].lenses, l)

	return len(lenses), false
}

func calcFocusingPower(focalLen int, boxNo int, slotNo int) int {
	power := (1 + boxNo) * (slotNo + 1) * (focalLen)
	return power
}

func parseStep(str []byte) lens {

	last := str[len(str)-1]

	if last == remove || last == upsert {
		return lens{
			opt:   last,
			label: string(str[:len(str)-1]),
		}
	}

	return lens{
		opt:      str[len(str)-2],
		focalLen: int(last - '0'),
		label:    string(str[:len(str)-2]),
	}
}
