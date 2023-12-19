package day3

import (
	"bufio"
	"fmt"
	"io"
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

/////////////////////////////////////////////////////////////////////////////////////

func Day3ex2() {

	f, err := os.Open("./day3/day3ex1_test.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)

	var line1, line2, line3 []byte
	fmt.Fprint(io.Discard, line3, line2, line1)

	lastLineScanned := !scanner.Scan()
	// using scanner.Bytes() produce 3 lines that were out of order... more than an hour lost...
	// line1 = []byte(scanner.Text())

	// scanner.Scan()
	line2 = []byte(scanner.Text())

	lineSize := len(line2)
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

			if i < lineSize && line1[i] == '*' { // inline after
				if i+1 < lineSize && isNumber(line1[i+1]) {
					numValue2 := readNumValue(line1, lineSize, i+1)
					tmp := numValue * numValue2
					total += tmp
					break
				} else if line2 != nil {
					for k := -1; k < 2; k++ {
						numValue2 := getMbNumber(line2, lineSize, i+k)
						if numValue2 != 0 {
							tmp := numValue * numValue2
							total += tmp
							break
						}
					}
				}
			} else if i-numLen-1 >= 0 && line1[i-numLen-1] == '*' { // inline before
				// we do not check if there is number before * in this case because if there was we already
				// included this pair because we travel from left to right
				if line2 != nil {
					tmpI := i - numLen - 1
					for k := -1; k < 2; k++ {
						if tmpI+k < 0 {
							continue
						}
						numValue2 := getMbNumber(line2, lineSize, tmpI+k)
						if numValue2 != 0 {
							tmp := numValue * numValue2
							total += tmp
							break
						}
					}
					// if numValue2:= getMbNumber(line2, lineSize, i)
					// if isNumber(line2[i]) {
					// 	numValue2 := readNumValue(line2, lineSize, i)
					// 	tmp := numValue * numValue2
					// 	total += tmp
					// 	break
					// } else if isNumber(line2[i-1]) {
					// 	numValue2 := readNumValue(line2, lineSize, i)
					// 	tmp := numValue * numValue2
					// 	total += tmp
					// 	break
					// } else if isNumber(line1[i+1]) {
					// 	numValue2 := readNumValue(line2, lineSize, i)
					// 	tmp := numValue * numValue2
					// 	total += tmp
					// 	break
					// }
				}
			}

		}

		cnt++
	}
	fmt.Printf("\n\nTotal: %d", total)
}

func getMbNumber(line []byte, lineSize, pos int) int {
	if !isNumber(line[pos]) {
		return 0
	}

	return readNumValue(line, lineSize, pos)
}

func readNumValue(line []byte, lineSize, pos int) int {
	for ; pos < lineSize; pos++ {
		if !(line[pos] >= '0' && line[pos] <= '9') {
			pos--
			break
		}
	}

	numLen := 0
	for ; pos >= 0; pos-- {
		if !(line[pos] >= '0' && line[pos] <= '9') {
			pos++
			break
		}

		numLen++
		continue
	}
	if pos == -1 {
		// pos eq -1 means that the start was at pos eq 0
		pos++
	}

	val := 0
	multiplier := 1
	for i := 0; i < numLen; i++ {
		tmp := int(line[pos+numLen-1-i] - '0')
		val += tmp * multiplier
		multiplier *= 10
	}

	return val
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
