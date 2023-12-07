package day2

import (
	"bufio"
	"bytes"
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

	var total int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		var green, blue, red int
		line := scanner.Bytes()
		lineString := string(line)
		fmt.Printf("\n%-85s", lineString)
		count := len(line)
		var b byte
		var number int
	gameLoop:
		for i := 5; i < count; i++ {
			b = line[i]
			// skip game id
			if b != ':' {
				continue
			}

			// to read gameId loop back a few characters checking if they are number
			i++ // skip ':'
			// after 'Game X: '
			for /*i < count */ {
				i++ // skip space

				b = line[i]
				number = int(b - '0')
				// peek next char, if it's digit update number
				b = line[i+1]
				if b != ' ' {
					number = number*10 + int(b-'0')
					i++
				}
				i++ // skip space
				i++ // get first char of the color
				b = line[i]
				switch b {
				case 'r':
					if number > red {
						red = number
					}
				case 'g':
					if number > green {
						green = number
					}
				case 'b':
					if number > blue {
						blue = number
					}
				default:
					colonIdx := bytes.IndexByte(line, ':')
					panic(fmt.Sprintf("\nincorrect color %q, at %d, in %q, here: %q", string(b), i, line[:colonIdx], line[i:]))
				}

				// skip rest of the color name's characters
				for ; ; i++ {
					if i == count {
						break gameLoop // end of game line
					}
					b = line[i]
					if b != ',' && b != ';' {
						continue
					}
					i++ // skip ',' or ';'
					break
				}
			}
		}

		fmt.Printf(" - {r: %2d, g: %2d, b: %2d} = %d", red, green, blue, red*green*blue)
		total += red * green * blue
	}
	fmt.Printf("\n total: %d", total)
}
