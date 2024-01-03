package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/unnamedxaer/aoc2023/help"
)

func calcNextValueOfTheHistoryRecord(values []int) int {

	valuesCnt := len(values)
	lastValues := make([]int, 0, valuesCnt/2+1)

	// fmt.Printf("\n\nvalues:= %v", values[:valuesCnt])
	for valuesCnt > 1 {
		allZeros := true
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
	recordsCnt := len(lastValues)
	tmp := 0
	for i := recordsCnt - 1; i >= 0; i-- {
		tmp = lastValues[i] + tmp
	}

	fmt.Printf("\n- prediction: %d", tmp)

	return tmp
}

func Ex1() {

	f, err := os.Open("./day9/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	total := 0
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

		nextValue := calcNextValueOfTheHistoryRecord(values)
		total += nextValue
	}

	fmt.Printf("\n\nTotal: %d", total)
}

func Ex2() {

}
