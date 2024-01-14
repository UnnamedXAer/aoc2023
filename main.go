package main

import (
	"fmt"
	"time"

	"github.com/unnamedxaer/aoc2023/day12"
)

func main() {
	// day1.Day1ex1()
	// day1.Day1ex2()
	// day2.Day2ex1()
	// day2.Day2ex2()
	// day3.Day3ex2()
	// day4.Day4ex1()
	// day4.Day4ex2()
	// day5.Day5ex1()
	// day5.Day5ex2()
	// day5.Day5ex2_v2()

	// day6.Day6ex1()
	// day6.Day6ex2()
	// day7.Ex1()
	// day7.Ex2()

	// day8.Ex1()
	// day8.Ex2()

	// day9.Ex1()
	// day9.Ex2()

	// day10.Ex1()
	// day10.Ex2()

	// day11.Ex1()
	// day11.Ex2()

	measure(day12.Ex1)
	fmt.Printf("\n--------------------------\n")
	measure(day12.Ex2)
}

func measure(fn func()) {
	start := time.Now().UnixNano()

	fn()

	end := time.Now().UnixNano()
	printTime(start, end)

}

func printTime(start, end int64) {
	timeTotal := end - start
	// fmt.Printf("\nnoano: %d", timeTotal)
	if timeTotal < 1000000 {
		fmt.Printf("\ntime: %d ns", timeTotal)
	} else if timeTotal < 10*1000000000 {
		fmt.Printf("\ntime: %d ms", timeTotal/1000000)
	} else if timeTotal < 120*1000000000 {
		fmt.Printf("\ntime: %d s", timeTotal/1000000000)
	} else {
		fmt.Printf("\ntime: %d min", timeTotal/60*1000000000)
	}
}
