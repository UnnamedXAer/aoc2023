package day3

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func Day3ex1() {

	f, err := os.Open("./day3/day3ex1_test.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)

	var lineAbove, line, lineBelow []byte

	lastLineScanned := !scanner.Scan()
	lineBelow = scanner.Bytes()
	count := len(lineBelow)

	for {
		if lastLineScanned {
			break
		}

		lineAbove = line
		line = lineBelow

		lastLineScanned = !scanner.Scan()
		if lastLineScanned {
			lineBelow = nil
		} else {
			lineBelow = scanner.Bytes()
		}

		numLen := 0
		for i := 0; i < count; i++ {
			if !(line[i] >= '0' && line[i] <= '9') {
				// skip until next number
				continue
			}

			// move after the number
			numLen = 0
			for i < count && line[i] >= '0' && line[i] <= '9' {
				i++
				numLen++
				// continue
			}

			///////////////////////////////////////////
			// number ended
			// check char before and after the number if exists
			if (i < count && isSymbol(line[i])) || (i-numLen-1 >= 0 && isSymbol(line[i-numLen-1])) {
				total += extractInt(line, i-1, numLen)
				numLen = 0
				break
			}

			// check chars line above and below
			tmpLine := lineAbove
		loopAboveBelow:
			for lIdx := 0; lIdx < 2; lIdx++ {
				if tmpLine != nil {
					for j := 0; j < numLen+2; j++ {
						idx := i - j
						if idx < 0 || idx >= count {
							continue
						}
						// if char is a symbol get the number and break from both loops
						if isSymbol(tmpLine[idx]) {
							total += extractInt(line, i-1, numLen)
							numLen = 0
							break loopAboveBelow
						}
					}
				}
				tmpLine = lineBelow
			}

			// check if haveNum

		}
	}
	fmt.Printf("\nTotal: %d", total)
}

func isSymbol(char byte) bool {

	if char != '.' && !(char >= '0' && char <= '9') {
		return true
	}
	return false
}

func extractInt(line []byte, endPos int, count int) int {
	var number int
	multiplier := 1
	for i := 0; i < count; i++ {
		number += multiplier * int(line[endPos-i]-'0')
		multiplier *= 10
	}
	return number
}
