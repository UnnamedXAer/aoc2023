package day1

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func Day1ex1() {

	file, err := os.Open("./day1/day1ex1.txt")
	help.IfErr(err)

	defer file.Close()

	sum := 0
	found := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		lineLength := len(line)

		found = 0
		for i := 0; i < lineLength; i++ {
			if found != 1 && line[i] >= 48 && line[i] <= 57 {
				// n[-1] = line[i]
				sum += 10 * int(line[i]-48)
				if found == 2 {
					break
				}
				found = 1

			}

			if found != 2 && line[lineLength-i-1] >= 48 && line[lineLength-i-1] <= 57 {
				sum += int(line[lineLength-i-1] - 48)
				if found == 1 {
					break
				}
				found = 2
			}
		}
	}
	fmt.Println(sum)
}
