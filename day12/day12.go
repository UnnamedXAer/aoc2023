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

	f, err := os.Open("./day12/data_t.txt")
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

	possibilitiesCntTotal := 0
	for _, r := range *doc {

		possibilities := generateAllPossibilities(r.record)
		possibilitiesCntTotal += len(possibilities)
		docTotal := 0
		for _, p := range possibilities {
			if isPossible(p, r.numbers) {
				docTotal++
			}
		}
		// fmt.Printf("\ndoc total: %d", docTotal)
		total += docTotal
	}

	fmt.Printf("\npossibilities total cnt: %d", possibilitiesCntTotal)
	fmt.Printf("\n\n Total: %d", total)
}

func isPossible(pattern []byte, numbers []int) bool {

	numIdx := 0
	groupSize := 0
	patternSize := len(pattern)
	numbersSize := len(numbers)

	for i := 0; i < patternSize; i++ {
		b := pattern[i]
		if b == unknown {
			return true
		}
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
		if numIdx < numbersSize && groupSize > numbers[numIdx] {
			return false
		}
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

	buff := make([][]byte, 0, 30)

	for k := 0; k < recordSize; k++ {

		if record[k] != unknown {
			continue
		}

		if len(buff) == 0 {
			c := make([]byte, recordSize)
			copy(c, record)
			c[k] = operational
			buff = append(buff, c)
			c = make([]byte, recordSize)
			copy(c, record)
			c[k] = damaged
			buff = append(buff, c)
			continue
		}

		buff = appendOptions(recordSize, k, buff)
	}

	// fmt.Printf("\n\npossibilities for: %v\n%v", string(record), string(bytes.Join(buff, []byte{'\n'})))

	return buff
}

// appendOption assumes that buff is not empty
func appendOptions(recordSize int, springPos int, buff [][]byte) [][]byte {

	initialSize := len(buff)

	for k := 0; k < initialSize; k++ {
		curr := buff[k]

		partialRecord := make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = operational
		buff[k] = partialRecord
		// buff = append(buff, string(partialRecord))

		partialRecord = make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = damaged
		buff = append(buff, partialRecord)
	}

	return buff
}

const operational = byte('.')
const damaged = byte('#')
const unknown = byte('?')

func Ex2() {
	p1 := !true
	if p1 {
		doc := extractData()
		total := 0
		possibilitiesCntTotal := 0
		for _, r := range *doc {

			possibilities := generateAllPossibilitiesWithEarlyValidation(r.record, r.numbers)
			possibilitiesCntTotal += len(possibilities)
			docTotal := 0
			for _, p := range possibilities {
				if isPossible(p, r.numbers) {
					docTotal++
				}
			}
			// fmt.Printf("\ndoc total: %d", docTotal)
			total += docTotal
		}

		fmt.Printf("\npossibilities total cnt: %d", possibilitiesCntTotal)

		fmt.Printf("\n\n Total: %d", total)

	} else {
		///////////////////////

		doc := extractData()
		total := 0
		singleTotal := 0
		for i, r := range *doc {
			docIdx := i + 1
			recordSize := len(r.record)
			numbersSize := len(r.numbers)

			possibilities := generateAllPossibilitiesWithEarlyValidation(r.record, r.numbers)
			noMultiplicationPossibilitiesCnt := len(possibilities)
			singleTotal += noMultiplicationPossibilitiesCnt
			const expectedMultiplication = 5
			const multiplier = 2

			record := make([]byte, recordSize*multiplier+(multiplier-1))
			numbers := make([]int, numbersSize*multiplier)

			for i := 0; i < multiplier; i++ {
				copy(record[i*recordSize+i:recordSize*(i+1)+i], r.record)
				if i < multiplier-1 {
					record[recordSize*(i+1)+i] = unknown
				}
				copy(numbers[i*numbersSize:numbersSize*(i+1)], r.numbers)
			}

			possibilities = generateAllPossibilitiesWithEarlyValidation(record, numbers)
			docTotal := len(possibilities)
			divider := float64(docTotal) / float64(noMultiplicationPossibilitiesCnt)
			intD := int(divider)

			for i := 0; i < expectedMultiplication-multiplier; i++ {
				docTotal *= intD
			}
			fmt.Printf("\n%3d. %d | %d | %f | %d", docIdx, noMultiplicationPossibilitiesCnt, len(possibilities), divider, docTotal)

			// fmt.Printf("\ndoc total: %d", docTotal)
			total += docTotal
			// break // TODO:
		}

		fmt.Printf("\n\n singleTotal: %d", singleTotal)
		fmt.Printf("\n\n Total: %d", total)
	}
}

// appendOption assumes that buff is not empty
func appendPossibleOptions(recordSize int, springPos int, buff [][]byte, numbers []int) [][]byte {

	initialSize := len(buff)

	for k := 0; k < initialSize; k++ {
		curr := buff[k]

		partialRecord := make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = operational

		if isPossible(partialRecord, numbers) {
			buff[k] = partialRecord
		} else {
			buff[k] = nil
		}

		partialRecord = make([]byte, recordSize)
		copy(partialRecord, curr)
		partialRecord[springPos] = damaged
		if isPossible(partialRecord, numbers) {
			if buff[k] == nil {
				buff[k] = partialRecord
			} else {
				buff = append(buff, partialRecord)
			}
		}
	}

	buff = removeEmpty(buff)

	return buff
}

func removeEmpty(buff [][]byte) [][]byte {
	buffS := make([]string, len(buff))
	for i, v := range buff {
		buffS[i] = string(v)
	}
	buffSize := len(buff)

	emptyRowIdx := 0
	for emptyRowIdx != -1 { // TODO: can we remove this one?
		k := emptyRowIdx
		emptyRowIdx = -1
		for ; k < buffSize; k++ {
			if buff[k] == nil {
				emptyRowIdx = k - 1
				buffSize--
				for j := buffSize; j > k; j-- {
					if buff[j] == nil {
						buffSize--
						continue
					}
					break
				}

				buff[k] = buff[buffSize]
			}
		}
	}

	buff = buff[:buffSize]
	buffS = make([]string, len(buff))
	for i, v := range buff {
		buffS[i] = string(v)
	}

	return buff
}

func generateAllPossibilitiesWithEarlyValidation(record []byte, numbers []int) [][]byte {
	recordSize := len(record)
	buff := make([][]byte, 0, 30)

	for k := 0; k < recordSize; k++ {

		if record[k] != unknown {
			continue
		}

		if len(buff) == 0 {
			c := make([]byte, recordSize)
			copy(c, record)
			c[k] = operational
			if isPossible(c, numbers) {
				buff = append(buff, c)
			}

			c = make([]byte, recordSize)
			copy(c, record)
			c[k] = damaged
			// buff = append(buff, c)
			if isPossible(c, numbers) {
				buff = append(buff, c)
			}
			continue
		}

		buff = appendPossibleOptions(recordSize, k, buff, numbers)
	}

	// fmt.Printf("\n\npossibilities for: %v\n%v", string(record), string(bytes.Join(buff, []byte{'\n'})))

	return buff
}
