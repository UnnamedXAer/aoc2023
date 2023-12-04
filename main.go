package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInt2(s []byte) int {
	var res int

	return res
}

func main() {
	// fmt.Printf("\n %v", int(byte(51)-48))
	// trebuchet()
	b := []byte{51, 57}
	fmt.Fprintf(os.Stdout, "%s", b)
	nn, err := strconv.Atoi(string(b))
	ifErr(err)

	fmt.Printf("\n %v", nn)
	var n int = getInt2([]byte{51, 57})
	fmt.Printf("\n %v", n)
}

func trebuchet() {

	file, err := os.Open("./day1ex1.txt")
	ifErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		cnt := len(line)
		n := make([]byte, 2)
		for i := 0; i < cnt; i++ {
			if line[i] >= 48 && line[i] <= 57 {
				n[0] = line[i]
			}

			if line[cnt-i-1] >= 48 && line[cnt-i-1] <= 57 {
				n[1] = line[cnt-i-1]
			}
		}
		fmt.Print(n)
	}
}

func ifErr(err error) {
	if err != nil {
		panic(err)
	}
}
