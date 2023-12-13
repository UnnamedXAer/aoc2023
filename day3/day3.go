package day3

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func Day3ex1() {

	f, err := os.Open("./day3/day3ex1.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)

	var lineAbove, line, lineBelow []byte

	lastLineScanned := !scanner.Scan()
	// using scanner.Bytes() produce 3 lines that were out of order... more than an hour lost...
	lineBelow = []byte(scanner.Text())
	count := len(lineBelow)

	cnt := 1

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
			// order issue
			lineBelow = []byte(scanner.Text())
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
				num := extractInt(line, i-1, numLen)
				total += num
				continue
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
							num := extractInt(line, i-1, numLen)
							total += num
							break loopAboveBelow
						}
					}
				}
				tmpLine = lineBelow
			}
		}
		cnt++
	}

	fmt.Printf("\n\nTotal: %d", total)
}

func isSymbol(char byte) bool {
	return char != '.' && !(char >= '0' && char <= '9')
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

func Day3ex2() {

	f, err := os.Open("./day3/day3ex1_test.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)

	var line1, line2, line3 []byte

	lastLineScanned := !scanner.Scan()
	// using scanner.Bytes() produce 3 lines that were out of order... more than an hour lost...
	line1 = []byte(scanner.Text())
	lineSize := len(line1)

	scanner.Scan()
	line2 = []byte(scanner.Text())

	cnt := 0
	for {
		if lastLineScanned {
			break
		}

		line1 = line2
		lastLineScanned = !scanner.Scan()
		if lastLineScanned {
			line2 = nil
		} else {
			// order issue
			line2 = []byte(scanner.Text())
		}

		for i := 0; i < lineSize; i++ {

			// go until find number
			if !(line1[i] >= '0' && line1[i] <= '9') {
				continue
			}

			numLen := 0
			for ; i < lineSize && (line1[i] >= '0' && line1[i] <= '9'); i++ {
				numLen++
			}

			numValue := extractInt(line1, i-1, numLen)
			fmt.Printf(", %3d", numValue)

			if i > 0 {
			}

		}

		cnt++
	}
	fmt.Printf("\n\nTotal: %d", total)
}
