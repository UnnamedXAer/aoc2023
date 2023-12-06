package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
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

func Day1ex2() {
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	// digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	words2 := make([]string, len(words))
	for i, v := range words {
		b := []byte(v)
		slices.Reverse(b)
		words2[i] = string(b)
	}

	reText := strings.Join(words, "|") + "|" + strings.Join(digits, "|")
	fmt.Printf("\nre text: %q", reText)
	re := regexp.MustCompile(reText)
	reTextB := []byte(reText)
	slices.Reverse(reTextB)
	re2 := regexp.MustCompile(string(reTextB))

	file, err := os.Open("./day1/day1ex2.txt")
	help.IfErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	fmt.Println()
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		fmt.Printf("%50s", line)
		match := re.FindString(string(line))

		fmt.Printf(" %-10v", fmt.Sprint(match))

		m := match
		idx := slices.Index(words, m)
		if idx == -1 {
			idx = slices.Index(digits, m)
			// b := unsafe.Slice(unsafe.StringData(m), 1)
			// idx = bytes.Index(digits, b)
		}

		if idx == -1 {
			fmt.Fprintf(os.Stderr, "for: %q, m = %q, index was 0", line, m)
			panic(fmt.Sprintf("for: %q, m = %q, index was 0", line, m))
		}
		idx += 1
		fmt.Printf(" - [%d]", idx)

		// last number
		lineSum := idx * 10
		// m = match[len(match)-1]
		// idx = slices.Index(words, m)
		// if idx == -1 {
		// 	idx = slices.Index(digits, m)
		// 	// b := unsafe.Slice(unsafe.StringData(m), 1)
		// 	// idx = bytes.Index(digits, b)
		// }

		slices.Reverse(line)
		match = re2.FindString(string(line))
		m = match
		fmt.Printf(" %-10v", fmt.Sprint(match))
		idx = slices.Index(words2, m)
		if idx == -1 {
			idx = slices.Index(digits, m)
			// b := unsafe.Slice(unsafe.StringData(m), 1)
			// idx = bytes.Index(digits, b)
		}

		if idx == -1 {
			fmt.Fprintf(os.Stderr, "for: %q, m = %q, index was 0", line, m)
			panic(fmt.Sprintf("for: %q, m = %q, index was 0", line, m))
		}

		idx += 1
		fmt.Printf(" [%d]", idx)
		lineSum += idx
		sum += lineSum
		fmt.Printf(" = %d\n", lineSum)
		i++
	}
	fmt.Printf("\n sum %d", sum)
}
