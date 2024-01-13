package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type rowDocumentation struct {
	record  []byte
	numbers []int
}

type documentation []*rowDocumentation

func (d *documentation) String() string {
	b := strings.Builder{}
	for _, r := range *d {
		b.WriteString(fmt.Sprintf("\n%s %v", r.record, r.numbers))
	}
	return b.String()
}

func extractData() *documentation {

	f, err := os.Open("./day12/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	doc := make(documentation, 0, 1000)

	for scanner.Scan() {
		var record []byte
		var numbers []int
		var line []byte = []byte(scanner.Text())
		lineSize := len(line)
		i := lineSize - 1
		for ; i >= 0 && line[i] != ' '; i-- {

			if line[i] == ',' {
				continue
			}
			n, l := help.ReadNumValueFromEnd(line, i)
			numbers = append(numbers, n)
			i -= l - 1
		}
		slices.Reverse(numbers)

		for k := 0; k < i; k++ {
			record = append(record, line[k])
		}

		doc = append(doc, &rowDocumentation{record: record, numbers: numbers})
	}

	return &doc
}

func Ex1() {
	doc := extractData()

	// fmt.Printf("\n%s", doc)
	total := 0

	for _, r := range *doc {

		possibilities := generateAllPossibilities(r.record)
		// fmt.Printf("\n\nr: %v | possibilities: %d", string(r.record), len(possibilities))
		docTotal := 0
		for _, p := range possibilities {
			if isPossible(p, r.numbers) {
				docTotal++
			}
		}
		fmt.Printf("\ndoc total: %d", docTotal)
		total += docTotal
	}

	fmt.Printf("\n\n Total: %d", total)
}

func isPossible(pattern []byte, numbers []int) bool {

	numIdx := 0
	groupSize := 0
	patternSize := len(pattern)
	numbersSize := len(numbers)

	for i := 0; i < patternSize; i++ {
		b := pattern[i]
		if b == operational {
			if groupSize > 0 {
				if numIdx >= numbersSize {
					return false
				}
				if groupSize != numbers[numIdx] {
					return false
				}
				numIdx++
			}
			groupSize = 0
			continue
		}
		groupSize++
	}

	if groupSize > 0 {
		if numIdx != numbersSize-1 {
			return false
		}

		if numbers[numIdx] != groupSize {
			return false
		}

		// fmt.Printf("\nP: %v | %v", string(pattern), numbers)
		return true
	}

	if numIdx != numbersSize {
		return false
	}

	// fmt.Printf("\nP: %v | %v", string(pattern), numbers)
	return true

}

func generateAllPossibilities(record []byte) [][]byte {
	recordSize := len(record)
	possibilities := make([][]byte, 0, 10)

	buff := make([]string, 0, 30)

	for k := 0; k < recordSize; k++ {

		if record[k] != unknown {
			continue
		}

		if len(buff) == 0 {
			c := make([]byte, recordSize)
			copy(c, record)
			c[k] = operational
			buff = append(buff, string(c))
			c = make([]byte, recordSize)
			copy(c, record)
			c[k] = damaged
			buff = append(buff, string(c))
			// buff = appendOptions(recordSize, string(record), k, buff)
			continue
		}

		buff = appendOptions(recordSize, k, buff)
	}

	for _, b := range buff {
		possibilities = append(possibilities, []byte(b))
	}

	// fmt.Printf("\n\npossibilities for: %v\n%v", string(record), string(bytes.Join(possibilities, []byte{'\n'})))

	return possibilities
}

// appendOption assumes that buff is not empty
func appendOptions(recordSize int, springPos int, buff []string) []string {

	initialSize := len(buff)

	for k := 0; k < initialSize; k++ {
		curr := buff[k]

		partialRecord := make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = operational
		buff[k] = string(partialRecord)
		// buff = append(buff, string(partialRecord))

		partialRecord = make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = damaged
		buff = append(buff, string(partialRecord))
	}

	return buff
}

const operational = byte('.')
const damaged = byte('#')
const unknown = byte('?')

func count(cfg []byte, nums []int) int {

	return -1
}

func Ex2() {

}
