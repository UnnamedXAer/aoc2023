package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/unnamedxaer/aoc2023/help"
)

func calcNextValueOfTheHistoryRecord(values []int) (int, int) {

	valuesCnt := len(values)
	lastValues := make([]int, 0, valuesCnt/2+1)
	firstValues := make([]int, 0, valuesCnt/2+1)

	// fmt.Printf("\n\nvalues:= %v", values[:valuesCnt])
	for valuesCnt > 1 {
		allZeros := true
		firstValues = append(firstValues, values[0])
		for i := 1; i < valuesCnt; i++ {
			values[i-1] = values[i] - values[i-1]
			if allZeros && values[i-1] != 0 {
				allZeros = false
			}
		}
		valuesCnt--
		// fmt.Printf("\nvalues:= %v", values[:valuesCnt])
		lastValues = append(lastValues, values[valuesCnt])
		if allZeros {
			break
		}
	}

	// fmt.Printf("\nlastValues:= %v", lastValues)
	// fmt.Printf("\nfirstValues:= %v", firstValues)
	recordsCnt := len(lastValues)
	tmp := 0
	tmpFront := 0
	for i := recordsCnt - 1; i >= 0; i-- {
		tmp = lastValues[i] + tmp
		tmpFront = firstValues[i] - tmpFront
	}

	fmt.Printf("\n- prediction end: %d, front: %d", tmp, tmpFront)

	return tmp, tmpFront
}

func Ex1() {

	f, err := os.Open("./day9/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	totalNext := 0
	totalFront := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		values := make([]int, 0, 50)
		for i := len(line) - 1; i >= 0; i-- {
			if help.IsNumber(line[i]) {
				v, l := help.ReadNumValueFromEnd(line, i)
				i -= l
				values = append(values, v)
			}
		}

		slices.Reverse(values)

		nextValue, prevValue := calcNextValueOfTheHistoryRecord(values)
		totalNext += nextValue
		totalFront += prevValue
	}

	fmt.Printf("\n\nTotal: next: %d, front: %d", totalNext, totalFront)
}

func Ex2() {
	Ex1()
}
