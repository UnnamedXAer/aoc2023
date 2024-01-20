package day13

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

const ash = byte('.')
const rocks = byte('#')

type pattern struct {
	lines    [][]byte
	lineSize int
	linesCnt int
}

func extractData() []pattern {

	f, err := os.Open("./day13/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	patterns := make([]pattern, 0, 100)

	var lines [][]byte = make([][]byte, 0, 20)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) == 0 {
			patterns = append(patterns, pattern{
				lines:    lines,
				lineSize: len(lines[0]),
				linesCnt: len(lines),
			})

			lines = make([][]byte, 0, 20)
			continue
		}
		lines = append(lines, line)
	}

	patterns = append(patterns, pattern{
		lines:    lines,
		lineSize: len(lines[0]),
		linesCnt: len(lines),
	})

	return patterns
}

func Ex1() {
	patterns := extractData()
	total := 0

	for _, p := range patterns {
		// fmt.Printf("\n\n%25s", string(bytes.Join(p.lines, []byte{'\n'})))
		patternTotal := processPattern(p)
		total += patternTotal
		// fmt.Printf(" - total: %v", patternTotal)
	}

	fmt.Printf("\n\n  Total: %d", total)
	// 28302 - too low
}

func processPattern(p pattern) int {

	total := 0
	lines := p.lines

colsLoop:
	for colIdx := 1; colIdx < p.lineSize; colIdx++ {
		for rowIdx := 0; rowIdx < p.linesCnt; rowIdx++ {
			if lines[rowIdx][colIdx-1] != lines[rowIdx][colIdx] {
				continue colsLoop
			}
		}

		colsToCompare := min(colIdx-1, p.lineSize-(colIdx+1))

		l1 := make([]byte, colsToCompare)
		var l2 []byte
		toRightIdxStart := colIdx + 1
		toLeftIdxEnd := colIdx - 2
		for rowIdx := 0; rowIdx < p.linesCnt; rowIdx++ {

			tmpL1 := lines[rowIdx][toLeftIdxEnd-colsToCompare+1 : toLeftIdxEnd+1]
			l2 = lines[rowIdx][toRightIdxStart : toRightIdxStart+colsToCompare]

			copy(l1, tmpL1)
			slices.Reverse(l1)
			if !bytes.Equal(l1, l2) {
				continue colsLoop
			}
		}

		total += colIdx
		break
	}

rowsLoop:
	for rowIdx := 1; rowIdx < p.linesCnt; rowIdx++ {

		if !linesEq(lines[rowIdx-1], lines[rowIdx]) {
			continue
		}

		rowsToCompare := min(rowIdx-1, p.linesCnt-(rowIdx+1))
		for k := 0; k < rowsToCompare; k++ {
			if !linesEq(lines[rowIdx-2-k], lines[rowIdx+1+k]) {
				continue rowsLoop
			}
		}

		total += 100 * rowIdx
		break
	}

	return total
}

func linesEq(l1, l2 []byte) bool {

	size := len(l1)
	for i := 0; i < size; i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func Ex2() {

}
