package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type rowDocumentation struct {
	record         []byte
	numbers        []int
	unknownCnt     int
	operationalCnt int
	damagedCnt     int
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

	linesCnt := -1
	for scanner.Scan() {
		linesCnt++
		var record []byte
		var numbers []int
		var unknownCnt, operationalCnt, damagedCnt int
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

		j := 0
		for k := 0; k < i; k++ {
			j = k
			for ; j < i && line[j] == operational; j++ {
				// skip multiple '.'

			}
			if j != k {
				record = append(record, operational)
				operationalCnt++
			}
			if j >= i {
				break
			}

			k = j
			if line[k] == unknown {
				unknownCnt++
			} else if line[k] == damaged {
				damagedCnt++
			} else {
				panic(fmt.Sprintf("\nunknown character %q in line %d, col %d", line[k], linesCnt, k))
			}

			record = append(record, line[k])
		}

		doc = append(doc, &rowDocumentation{
			record:         record,
			numbers:        numbers,
			unknownCnt:     unknownCnt,
			operationalCnt: operationalCnt,
			damagedCnt:     damagedCnt,
		})
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
		for _, record := range possibilities {
			if isPossible(record, r.numbers) {
				docTotal++
			}
		}
		// fmt.Printf("\ndoc total: %d", docTotal)
		total += docTotal
	}

	fmt.Printf("\npossibilities total cnt: %d", possibilitiesCntTotal)
	fmt.Printf("\n\n Total: %d", total)
}

func isPossible(record []byte, numbers []int) bool {

	numIdx := 0
	groupSize := 0
	recordSize := len(record)
	numbersSize := len(numbers)

	for i := 0; i < recordSize; i++ {
		b := record[i]
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
	try := 3
	if try == 1 {
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

	} else if try == 2 {
		///////////////////////

		doc := extractData()
		total := 0
		singleTotal := 0
		for i, r := range *doc {
			// fmt.Printf("\nrecord: %v", string(r.record))
			docIdx := i + 1
			recordSize := len(r.record)
			numbersSize := len(r.numbers)

			possibilities := generateAllPossibilitiesWithEarlyValidation(r.record, r.numbers)
			noMultiplicationPossibilitiesCnt := len(possibilities)
			singleTotal += noMultiplicationPossibilitiesCnt
			const expectedMultiplication = 5
			const multiplier = 5

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
	} else if try == 3 {
		ex2Try3()
	} else {
		panic("wrong part value")
	}
}

func ex2Try3() {
	doc := extractData()
	totalInt := 0

	cache = resultCache{data: make(map[string]int, 5_000_000)}

	docTotal := make(chan int)
	jobs := make(chan *rowDocumentation, 10)

	go processJobs(jobs, docTotal)

	go func() {
		for _, r := range *doc {
			jobs <- r
		}
		close(jobs)
	}()

	for k := 0; k < len(*doc); k++ {
		totalInt += <-docTotal
		// fmt.Printf("\n%3d. jobs processed - with total: %10d", k+1, totalInt)
	}

	fmt.Printf("\n\n Total: %d, | %d", totalInt, len(cache.data))
}

func processJobs(jobs <-chan *rowDocumentation, results chan<- int) {
	i := 0
	for job := range jobs {
		i++
		go func(i int, r *rowDocumentation) {
			results <- calcDocument(i, r)
		}(i, job)
	}
}

func calcDocument(i int, r *rowDocumentation) int {
	// docIdx := i + 1
	recordSize := len(r.record)
	numbersSize := len(r.numbers)

	const multiplier = 5

	record := make([]byte, recordSize*multiplier+(multiplier-1))
	numbers := make([]int, numbersSize*multiplier)
	unknownCnt := r.unknownCnt * multiplier

	for i := 0; i < multiplier; i++ {
		copy(record[i*recordSize+i:recordSize*(i+1)+i], r.record)
		if i < multiplier-1 {
			record[recordSize*(i+1)+i] = unknown
			unknownCnt++
		}
		copy(numbers[i*numbersSize:numbersSize*(i+1)], r.numbers)
	}

	docTotal := tryNext(record, -1, numbers, unknownCnt)
	// fmt.Printf("\n%3d. %25s / %v => %5d", docIdx, string(r.record), r.numbers, docTotal)

	return docTotal
}

type resultCache struct {
	data map[string]int
	mu   sync.RWMutex
}

// var cache = sync.Map{}

func (c *resultCache) Store(key string, value int) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *resultCache) Load(key string) (int, bool) {
	c.mu.RLock()
	res, ok := c.data[key]
	c.mu.RUnlock()

	return res, ok
}

var cache resultCache

func tryNext(record []byte, pos int, numbers []int, unknownLeft int) int {
	pos += 1
	cnt := 0

	recordSize := len(record)
	for ; pos < recordSize && record[pos] != unknown; pos++ {
	}

	if pos == recordSize {
		return cnt
	}

	unknownLeft -= 1

	key := string(record)
	for _, v := range numbers {
		key += string(v)
	}

	res, ok := cache.Load(key)
	if ok {
		// fmt.Printf("\n hit: %d, cnt: %d -> %t | %s", old, cnt, eq, key)
		// return res.(int)
		return res
	}

	cnt += tryWith(operational, record, pos, numbers, unknownLeft)
	cnt += tryWith(damaged, record, pos, numbers, unknownLeft)

	cache.Store(key, cnt)
	// cache[key] = cnt

	return cnt
}

func tryWith(spring byte, record []byte, pos int, numbers []int, unknownLeft int) int {
	cnt := 0

	record[pos] = spring

	isOk, numIdx, lastPassedGroupEndPos := isPossibleWithMoreInfo(record, numbers)

	// fmt.Printf("\nfor: %4v, len: %3d, ok: %5t, numIdx: %3v, passedGroups: %3v", string(rec), len(rec), isOk, numIdx, lastPassedGroupEndPos)
	if isOk {
		if unknownLeft == 0 {
			cnt++
		}

		if unknownLeft > 0 {
			posNew := pos
			if lastPassedGroupEndPos != 0 {
				posNew = 0
			}

			cnt += tryNext(record[lastPassedGroupEndPos:], posNew, numbers[numIdx:], unknownLeft)
		}

	}
	record[pos] = unknown

	return cnt
}

func isPossibleWithMoreInfo(record []byte, numbers []int) (isOk bool, numIdx int, lastPassedGroupEndPos int) {

	groupSize := 0
	recordSize := len(record)
	numbersSize := len(numbers)

	i := 0
	for ; i < recordSize; i++ {
		b := record[i]
		if b == unknown {
			return true, numIdx, lastPassedGroupEndPos
		}
		if b == operational {
			if groupSize > 0 {
				if numIdx >= numbersSize {
					return false, numIdx, lastPassedGroupEndPos
				}
				if groupSize != numbers[numIdx] {
					return false, numIdx, lastPassedGroupEndPos
				}
				lastPassedGroupEndPos = i
				numIdx++
			}
			groupSize = 0
			continue
		}
		groupSize++
		if numIdx < numbersSize && groupSize > numbers[numIdx] {
			return false, numIdx, lastPassedGroupEndPos
		}
	}

	if groupSize > 0 {
		if numIdx != numbersSize-1 {
			return false, numIdx, lastPassedGroupEndPos
		}

		if numbers[numIdx] != groupSize {
			return false, numIdx, lastPassedGroupEndPos
		}

		return true, numIdx + 1, len(record)
	}

	if numIdx != numbersSize {
		return false, numIdx, lastPassedGroupEndPos
	}

	return true, numIdx, len(record)
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
