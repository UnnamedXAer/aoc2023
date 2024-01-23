package day16

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]byte {

	f, err := os.Open("./day16/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	tiles := make([][]byte, 0, 110)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		tiles = append(tiles, line)
	}

	help.IfErr(scanner.Err())

	return tiles
}

func Ex1() {
	tiles := extractData()
	fmt.Printf("\n%s", string(bytes.Join(tiles, []byte{'\n'})))

}

func Ex2() {

}
