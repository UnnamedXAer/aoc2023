package day15

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]byte {

	splitFn := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		commaidx := bytes.IndexByte(data, ',')
		if commaidx > 0 {
			// we need to return the next position
			buffer := data[:commaidx]
			return commaidx + 1, buffer, nil
		}

		// if we are at the end of the string, just return the entire buffer
		if atEOF {
			// but only do that when there is some data. If not, this might mean
			// that we've reached the end of our input CSV string
			if len(data) > 0 {
				return len(data), data, nil
			}
		}

		// when 0, nil, nil is returned, this is a signal to the interface to read
		// more data in from the input reader. In this case, this input is our
		// string reader and this pretty much will never occur.
		return 0, nil, nil
	}

	f, err := os.Open("./day15/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)
	scanner.Split(splitFn)

	input := make([][]byte, 0, 4001)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		input = append(input, line)
	}

	help.IfErr(scanner.Err())

	return input
}

func Ex1() {
	total := 0
	input := extractData()
	for _, str := range input {
		total += process_hash([]byte(str))
	}

	fmt.Printf("\n\n Total: %d", total)
}

func process_hash(str []byte) int {
	currValue := 0
	for _, b := range str {
		currValue += int(b)
		currValue *= 17
		currValue %= 256
	}

	return currValue
}

func Ex2() {

}
