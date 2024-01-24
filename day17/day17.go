package day17

import (
	"bufio"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() (any, any) {

	f, err := os.Open("./day17/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Bytes()
	}

	help.IfErr(scanner.Err())

	return nil, nil
}

func Ex1() {
	_, _ := extractData()
}

func Ex2() {

}

