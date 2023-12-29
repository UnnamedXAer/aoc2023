package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
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

	// fmt.Printf("\nseeds: %d, %v", seedsCnt, seeds)

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

	// fmt.Println()
	// for _, seed := range seeds {
	for i := seedsCnt - 1; i > 0; i -= 2 {
		fmt.Printf("\nChecking seeds from: %d to %d, length: %d", seeds[i], seeds[i]+seeds[i-1]-1, seeds[i-1])
		for seed := seeds[i]; seed < seeds[i]+seeds[i-1]; seed++ {
			prevNumber := seed
			for _, k := range mapsInfo.keys {
				// fmt.Printf("\n%23s: %3d -> ", k, prevNumber)
				prevNumber = findCorrespondingNumber(prevNumber, mapsInfo.maps[k])
				// fmt.Printf("%3d", prevNumber)

			}
			// fmt.Printf("\n -- for seed: %d, we got location: %d", seed, prevNumber)
			if prevNumber < lowestLocation {
				lowestLocation = prevNumber
			}
		}
	}

	fmt.Printf("\n\nlowest location: %d", lowestLocation)
}
