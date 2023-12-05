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

	reText := "(?i)" + strings.Join(words, "|") + "|" + strings.Join(digits, "|")
	fmt.Printf("\nre text: %q", reText)
	re := regexp.MustCompile(reText)

	ord := make([]string, 0, 1000)

	file, err := os.Open("./day1/day1ex2_test.txt")
	help.IfErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	fmt.Println()
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		fmt.Printf("%50s", line)
		ord = append(ord, fmt.Sprintf("%55s", line))
		match := re.FindAllString(string(line), -1)

		ord[i] += fmt.Sprintf("[ %-12v ]", strings.Join([]string{match[0], match[len(match)-1]}, " "))
		fmt.Printf(" %-30v", fmt.Sprint(match))

		m := match[0]
		idx := slices.Index(words, m)
		if idx == -1 {
			idx = slices.Index(digits, m)
			// b := unsafe.Slice(unsafe.StringData(m), 1)
			// idx = bytes.Index(digits, b)
		}

		if idx == -1 {
			fmt.Fprintf(os.Stderr, "for: %q, m = %q, index was 0", line, m)
			panic(fmt.Sprintf("for: %q, m = %q, index was 0", line, m))

			continue
		}
		idx += 1
		fmt.Printf(" - [%d]", idx)
		ord[i] += fmt.Sprintf(" - [%d]", idx)

		// last number
		lineSum := idx * 10
		m = match[len(match)-1]
		idx = slices.Index(words, m)
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
		ord[i] += fmt.Sprintf(" [%d]", idx)
		lineSum += idx
		sum += lineSum
		fmt.Printf(" = %d\n", lineSum)
		ord[i] += fmt.Sprintf(" = %d\n", lineSum)
		i++
	}
	slices.Sort(ord)
	fmt.Printf("\n\n%v", ord)
	fmt.Printf("\n sum %d", sum)
}
