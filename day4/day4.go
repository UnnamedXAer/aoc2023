package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func Day4ex1() {

	f, err := os.Open("./day4/data.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = []byte(scanner.Text())

	lineSize := len(line)
	startPos := bytes.IndexByte(line, ':')
	pipePos := bytes.IndexByte(line, '|')

	cnt := 0
	for line != nil {
		lineValue := 0
	loopWinning:
		for i := startPos + 2; i < pipePos; i += 3 {
			for k := pipePos + 2; k < lineSize; k += 3 {
				if line[i] == line[k] && line[i+1] == line[k+1] {
					if lineValue == 0 {
						lineValue = 1
					} else {
						lineValue *= 2
					}
					continue loopWinning
				}

			}

		}

		total += lineValue
		// scan next line
		cnt++
		if scanner.Scan() {
			// scan bytes and copy()???
			line = []byte(scanner.Text())
		} else {
			line = nil
		}
	}

	fmt.Printf("\n\nTotal: %d", total)
}

func Day4ex2() {

	f, err := os.Open("./day4/data.txt")
	help.IfErr(err)

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = []byte(scanner.Text())

	lineSize := len(line)
	startPos := bytes.IndexByte(line, ':')
	pipePos := bytes.IndexByte(line, '|')

	linesCnt := make([]int, 0, 20)
	linesValues := make([]int, 0, 20)

	cnt := 0
	for line != nil {
		lineValue := 0
	loopWinning:
		for i := startPos + 2; i < pipePos; i += 3 {
			for k := pipePos + 2; k < lineSize; k += 3 {
				if line[i] == line[k] && line[i+1] == line[k+1] {
					lineValue += 1
					continue loopWinning
				}
			}
		}

		if len(linesCnt) <= cnt {
			linesCnt = append(linesCnt, 0)
		}

		total += linesCnt[cnt] + 1

		for i := 0; i < linesCnt[cnt]+1; i++ {
			for k := 1; k < lineValue+1; k++ {
				if len(linesCnt) <= cnt+k {
					linesCnt = append(linesCnt, 0)
				}
				linesCnt[cnt+k]++
			}
		}

		linesValues = append(linesValues, lineValue)

		// scan next line
		cnt++
		if scanner.Scan() {
			// scan bytes and copy()???
			line = []byte(scanner.Text())
		} else {
			line = nil
		}
	}

	fmt.Printf("\n\n  cnt: %d, %d", cnt, len(linesValues))
	fmt.Printf("\nvalues: %v", linesValues)
	fmt.Printf("\ncounts: %v", linesCnt)

	fmt.Printf("\n\nTotal: %d", total)
}
