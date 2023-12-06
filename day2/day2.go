package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
)

func Day2ex1() {
	f, err := os.Open("./day2/day2ex1.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// var r, g, b int
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	total := 0
	cubesCount := 0
	nums := [2]int{}
	for scanner.Scan() {
		gamePossible := true
		line := scanner.Text()
		colonIdx := strings.IndexByte(line, ':')
		gLine := line[colonIdx+1:]
		draws := strings.Split(gLine, "; ")
	loopGame:
		for _, draw := range draws {

			drawResults := strings.Split(draw, ",")

		loopDrawResults:
			for _, result := range drawResults {

				// fmt.Printf("|%v", result)
				nums[0] = 0
				nums[1] = 0
				for _, b := range []byte(result) {
					if b == ' ' {
						continue
					}
					if b >= 48 && b <= 57 {
						if nums[0] == 0 {
							nums[0] = int(b - 48)
						} else {
							nums[1] = int(b - 48)
						}
						continue
					}

					if b == 'r' || b == 'g' || b == 'b' {

						if nums[1] == 0 {
							cubesCount = nums[0]
						} else {
							cubesCount = nums[0]*10 + nums[1]
						}
						if b == 'r' {
							if cubesCount > maxRed {
								fmt.Printf("\n N: %s => %d, %s", []byte{b}, cubesCount, line)
								gamePossible = false
							}
						} else if b == 'g' {
							if cubesCount > maxGreen {
								gamePossible = false
							}
						} else {
							if cubesCount > maxBlue {
								gamePossible = false
							}
						}

						if !gamePossible {
							break loopGame
						}

						continue loopDrawResults
					}
				}

				panic(fmt.Sprintf("unknown bad line: %q", line))
			}
		}

		if !gamePossible {
			fmt.Printf("\n %v", line)
			continue
		}

		gameId := 0
		k := 1
		for i := colonIdx - 1; i >= 5; i-- {
			b := line[i]
			gameId += k * int(b-48)
			k *= 10
		}
		fmt.Printf("\nY: %d, %s", gameId, line)
		total += gameId

	}
	fmt.Printf("\n Total: %d", total)
}

func Day2ex2() {
	f, err := os.Open("./day2/day2ex1.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

	}
}
