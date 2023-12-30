package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

type theezzMaps struct {
	keys []string
	maps map[string][][3]int
}

func newTheezzMaps() theezzMaps {
	m := theezzMaps{
		keys: []string{
			// "seed-soil",
			// "soil-fertilizer",
			// "fertilizer-water",
			// "water-light",
			// "light-temperature",
			// "temperature-humidity",
			// "humidity-location",
		},
		maps: make(map[string][][3]int, 0),
	}

	// for _, key := range m.keys {
	// 	m.maps[key] = make([][3]int, 0, 5)
	// }

	return m
}

func (m theezzMaps) String() string {
	builder := strings.Builder{}

	builder.WriteString("map[")
	for _, k := range m.keys {
		builder.WriteString("\n  " + k + ": [")
		for _, v := range m.maps[k] {
			builder.WriteString(fmt.Sprintf("\n    %v", v))
		}
		builder.WriteString(" ]")
	}

	builder.WriteString(" ]")

	return builder.String()
}

func Day5ex1() {

	f, err := os.Open("./day5/data.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = []byte(scanner.Text())

	lineSize := len(line)

	seeds := make([]int, 0, 20)
	seedsCnt := 0
	// read seeds:
	for i := lineSize - 1; line[i] != ':'; i-- {
		seeds = append(seeds, 0)
		multiplier := 1

		for ; line[i] != ' '; i-- {
			seeds[seedsCnt] += int(line[i]-'0') * multiplier
			multiplier *= 10
		}

		seedsCnt++
	}

	fmt.Printf("\nseeds: %d, %v", seedsCnt, seeds)

	scanner.Scan() // skip empty line

	mapsInfo := newTheezzMaps()
	var key string
	for scanner.Scan() {
		// get the next map name
		line = scanner.Bytes()
		lineSize := len(line)

		key = string(line[:lineSize-5])

		readMap(scanner, &mapsInfo, key)
	}

	var lowestLocation int = math.MaxInt

	fmt.Println()
	for _, seed := range seeds {
		prevNumber := seed
		for _, k := range mapsInfo.keys {
			fmt.Printf("\n%23s: %3d -> ", k, prevNumber)
			prevNumber = findCorrespondingNumber(prevNumber, mapsInfo.maps[k])
			fmt.Printf("%3d", prevNumber)

		}
		fmt.Printf("\n -- for seed: %d, we got location: %d", seed, prevNumber)
		if prevNumber < lowestLocation {
			lowestLocation = prevNumber
		}
	}

	fmt.Printf("\n\nlowest location: %d", lowestLocation)
}

func findCorrespondingNumber(src int, m [][3]int) int {

	size := len(m)
	var rangeStart, rangeEnd int
	for i := 0; i < size; i++ {
		rangeStart = m[i][1]
		rangeEnd = rangeStart + m[i][2] - 1
		if src >= rangeStart && src <= rangeEnd {
			offset := src - rangeStart
			dstRangeStart := m[i][0]
			dst := dstRangeStart + offset
			return dst
		}
	}

	return src // not found so it maps to the same number
}

func readMap(scanner *bufio.Scanner, mapsInfo *theezzMaps, key string) {
	mapsInfo.keys = append(mapsInfo.keys, key)

	maps := mapsInfo.maps
	cnt := 0
	maps[key] = make([][3]int, 0, 40)
	for scanner.Scan() {
		var line []byte = scanner.Bytes()
		lineSize := len(line)
		if lineSize == 0 {
			return
		}

		maps[key] = append(maps[key], [3]int{})

		multiplier := 1
		value := 0
		pos := 2
		for i := lineSize - 1; i >= 0; i-- {
			if line[i] == ' ' {
				maps[key][cnt][pos] = value
				multiplier = 1
				value = 0
				pos--
				continue
			}

			value += int(line[i]-'0') * multiplier
			multiplier *= 10
		}

		maps[key][cnt][pos] = value
		cnt++
	}
}

func Day5ex2() {

	startTime := time.Now().UnixNano()
	f, err := os.Open("./day5/data_t.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = scanner.Bytes()

	seeds := readSeeds(line)
	mapsInfo := readMapsInfo(scanner, line)
	out := removeOverlapsInSeeds(seeds)

	lowestLocation, startMappingTime := findLowestRange(mapsInfo, out)

	fmt.Printf("\n\nlowest location: %d", lowestLocation)
	if lowestLocation != 37384986 {
		fmt.Printf("\nincorrect answer: %d, want: %d", lowestLocation, 37384986)
		//  342851781700 - base total time
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("\ntotal time: %v, mapping time: %v", endTime-startTime, endTime-startMappingTime)
}

func readSeeds(line []byte) []int {
	lineSize := len(line)

	seeds := make([]int, 0, 20)
	seedsCnt := 0

	// read seeds:
	for i := lineSize - 1; line[i] != ':'; i-- {
		seeds = append(seeds, 0)
		multiplier := 1

		for ; line[i] != ' '; i-- {
			seeds[seedsCnt] += int(line[i]-'0') * multiplier
			multiplier *= 10
		}

		seedsCnt++
	}
	return seeds
}

func readMapsInfo(scanner *bufio.Scanner, line []byte) theezzMaps {
	// skip empty line
	scanner.Scan()

	mapsInfo := newTheezzMaps()
	var key string
	for scanner.Scan() {

		line = scanner.Bytes()
		lineSize := len(line)

		// get the next map name
		key = string(line[:lineSize-5])

		readMap(scanner, &mapsInfo, key)
	}
	return mapsInfo
}

func findLowestRange(mapsInfo theezzMaps, out []int) (int, int64) {
	keys := mapsInfo.keys
	var lowestLocation int = math.MaxInt

	startMappingTime := time.Now().UnixNano()
	for i := len(out) - 1; i > 0; i -= 2 {

		for seed := out[i]; seed < out[i]+out[i-1]; seed++ {
			prevNumber := seed
			for _, k := range keys {

				prevNumber = findCorrespondingNumber(prevNumber, mapsInfo.maps[k])

			}

			if prevNumber < lowestLocation {
				lowestLocation = prevNumber
			}
		}
	}
	return lowestLocation, startMappingTime
}

func removeOverlapsInSeeds(seeds []int) []int {

	anyMerged := false

	ranges := SortSeeds(seeds)
	// fmt.Printf("\nsorted seeds: %v", ranges)

	rangesCnt := len(ranges)
	out := make([]int, 0, rangesCnt)

	out = append(out, ranges[0][1], ranges[0][0])

	for i := 0; i < rangesCnt-1; i++ {
		rangeStart := ranges[i][0]
		rangeSize := ranges[i][1]
		rangeEnd := rangeStart + rangeSize - 1

		rangeStart2 := ranges[i+1][0]
		rangeSize2 := ranges[i+1][1]
		rangeEnd2 := rangeStart2 + rangeSize2 - 1

		if rangeStart2 > rangeEnd {
			// range2 is entirely after range 1
			out = append(out, rangeSize2, rangeStart2)
			// anyMerged = true
			continue
		}

		if rangeStart2 >= rangeStart && rangeEnd2 <= rangeEnd {
			// range entirely inside range 1
			continue
		}

		// we are left with ranges that starts inside range 1 and end outside range 1

		rs := rangeEnd + 1
		re := rangeEnd2
		rSize := re - rs + 1

		anyMerged = true
		out = append(out, rs, rSize)
	}

	// fmt.Printf("\nout: %v", out)
	if anyMerged {
		return removeOverlapsInSeeds(out)
	}

	return out
}

func SortSeeds(seeds []int) [][2]int {
	seedsCnt := len(seeds)
	ranges := make([][2]int, seedsCnt/2)

	insertedCnt := 0
	for i := 0; i < seedsCnt; i += 2 {
		r := [2]int{seeds[i+1], seeds[i]}
		for j := 0; j < len(ranges); j++ {
			if ranges[j][0] == 0 {
				ranges[j] = r
				insertedCnt++
				break
			}

			if r[0] > ranges[j][0] {
				continue
			}

			// more next elements by one position
			for k := i / 2; k >= j && k > 0; k-- {
				ranges[k] = ranges[k-1]
			}

			ranges[j] = r
			break
		}

	}

	return ranges
}

func extractSeedsAndMaps() ([]int, theezzMaps) {
	f, err := os.Open("./day5/data_t.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = scanner.Bytes()

	seeds := readSeeds(line)
	mapsInfo := readMapsInfo(scanner, line)
	return seeds, mapsInfo
}

func Day5ex2_v2() {
	startTime := time.Now().UnixNano()
	seeds, mapsInfo := extractSeedsAndMaps()
	slices.Reverse(seeds) // I don't want to copy/modify previous functions so we just quick fix by reverse

	// seedsIntervals are [start,end] inclusive
	seedsIntervals := seedsToIntervals(seeds)
	fmt.Printf("\nseeds intervals:\n%v", seedsIntervals)

	updateMapsInfo(&mapsInfo)

	resultLocationIntervals := processSeedsThroughMaps(seedsIntervals, mapsInfo)

	lowestLocation := findLowestLocationInIntervals(resultLocationIntervals)

	fmt.Printf("\n\nlowest location: %d", lowestLocation)
	if lowestLocation != 37384986 {
		fmt.Printf("\nincorrect answer: %d, want: %d", lowestLocation, 37384986)
		//  342851781700 - base total time
	}
	endTime := time.Now().UnixNano()
	// fmt.Printf("\ntotal time: %v, mapping time: %v", endTime-startTime, endTime-startMappingTime)

	fmt.Printf("\ntotal time: %dms, %dµs", (endTime-startTime)/1000, endTime-startTime)
}

func processSeedsThroughMaps(seedsIntervals [][2]int, mapsInfo theezzMaps) [][2]int {

	locIntervals := make([][2]int, 0, len(seedsIntervals))

	keys := mapsInfo.keys

	for _, seed := range seedsIntervals {
	loopMappers:
		for _, key := range keys {
			mapper := mapsInfo.maps[key]
			for _, mapping := range mapper {
				if mapping[1] < seed[0] {
					continue
				}

				// we found the first interval that overlaps with the seed

				end := min(mapping[1], seed[1]) + mapping[2]
				start := seed[0] + mapping[2]
				locIntervals = append(locIntervals, [2]int{
					start,
					end,
				})

				seed[0] = end + 1

				// interval's start greater than interval's end means that we "extracted" all values
				if seed[0]+1 > seed[1] {
					break loopMappers
				}
			}
		}
	}

	fmt.Printf("locIntervals: %v", locIntervals)

	return locIntervals
}

func findLowestLocationInIntervals(intervals [][2]int) int {
	lowest := intervals[0][0]

	for _, v := range intervals {
		if v[0] < lowest {
			lowest = v[0]
		}
	}

	return lowest
}

func updateMapsInfo(mapsInfo *theezzMaps) {
	keys := mapsInfo.keys
	maps := mapsInfo.maps

	// sort mappings
	for _, key := range keys {
		m := maps[key]

		fmt.Println()
		// for i, mapping := range m {
		for i := 0; i < len(m); i++ {
			mapping := m[i]
			// fmt.Printf("\n%2d. from: %v", i, mapping)
			length := mapping[2]
			source := mapping[1]
			destination := mapping[0]
			end := source + length - 1
			offset := destination - source
			// {start, end, offset}
			mapping[0] = source
			mapping[1] = end
			mapping[2] = offset
			// fmt.Printf("\n%2d.   to: %v", i, mapping)
			m[i] = mapping
		}
		// fmt.Printf("\nto intervals: %v", m)

		slices.SortFunc(m, func(a, b [3]int) int {
			return a[0] - b[0]
		})
		// fmt.Printf("\n      sorted: %v", m)

		// ensure map starts at 0

		if m[0][0] != 0 {
			// insert interval from 0 to start of the current first interval with offset eq 0
			// {start, end, offset}
			newInterval := [3]int{0, m[0][0] - 1, 0}
			m = append(m, newInterval) // just a placeholder
			copy(m[1:], m)
			m[0] = newInterval
		}

		// fmt.Printf("\n  with zeros: %v", m)

		// from last to infinity
		m = append(m, [3]int{m[len(m)-1][1] + 1, math.MaxInt, 0})
		// fmt.Printf("\n    with inf: %v", m)

		// fill hols
		i := 1
		for i < len(m) {

			if m[i-1][1]+1 < m[i][0] {
				// we got a hole
				m = slices.Insert(m, i, [3]int{m[i-1][1] + 1, m[i][0] - 1, 0})
				// probably i++ here
				i++
			}

			i++
		}

		// fmt.Printf("\n  with holes: %v", m)
		maps[key] = m
	}
}

func seedsToIntervals(seeds []int) [][2]int {
	ranges := seedsToSortedRanges(seeds)
	uniqueRanges := removeOverlaps(ranges) // mb unnecessary
	intervals := make([][2]int, len(uniqueRanges))
	for i := 0; i < len(intervals); i++ {
		intervals[i] = [2]int{uniqueRanges[i][0], uniqueRanges[i][0] + uniqueRanges[i][1] - 1}
	}

	return intervals
}

func removeOverlaps(ranges [][2]int) [][2]int {

	anyMerged := false

	rangesCnt := len(ranges)
	intervals := make([][2]int, 0, rangesCnt)

	intervals = append(intervals, ranges[0])

	for i := 0; i < rangesCnt-1; i++ {
		rangeStart := ranges[i][0]
		rangeSize := ranges[i][1]
		rangeEnd := rangeStart + rangeSize - 1

		rangeStart2 := ranges[i+1][0]
		rangeSize2 := ranges[i+1][1]
		rangeEnd2 := rangeStart2 + rangeSize2 - 1

		if rangeStart2 > rangeEnd {
			// range2 is entirely after range 1
			intervals = append(intervals, [2]int{rangeStart2, rangeSize2})
			// anyMerged = true
			continue
		}

		if rangeStart2 >= rangeStart && rangeEnd2 <= rangeEnd {
			// range entirely inside range 1
			continue
		}

		// we are left with ranges that starts inside range 1 and end outside range 1

		rs := rangeEnd + 1
		re := rangeEnd2
		rSize := re - rs + 1

		anyMerged = true
		intervals = append(intervals, [2]int{rs, rSize})
	}

	// fmt.Printf("\nout: %v", out)
	if anyMerged {
		return removeOverlaps(intervals)
	}

	return intervals
}

func seedsToSortedRanges(seeds []int) [][2]int {
	seedsCnt := len(seeds)
	ranges := make([][2]int, seedsCnt/2)

	insertedCnt := 0
	for i := 0; i < seedsCnt; i += 2 {
		r := [2]int{seeds[i], seeds[i+1]}
		for j := 0; j < len(ranges); j++ {
			if ranges[j][0] == 0 {
				ranges[j] = r
				insertedCnt++
				break
			}

			if r[0] > ranges[j][0] {
				continue
			}

			// more next elements by one position
			for k := i / 2; k >= j && k > 0; k-- {
				ranges[k] = ranges[k-1]
			}

			ranges[j] = r
			break
		}

	}

	return ranges
}
