package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() [][]int {

	f, err := os.Open("./day17/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	blocks := make([][]int, 0, 141)

	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Bytes()
		blocks = append(blocks, make([]int, len(line)))

		for k, v := range line {
			blocks[i][k] = int(v - '0')
		}
	}

	help.IfErr(scanner.Err())

	if len(blocks) != len(blocks[0]) {
		panic("blocks not a square")
	}

	return blocks
}

type direction int

const (
	north direction = iota
	south
	west
	east
)

var offsets = [...][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var results = make([]int, 0, 100)

func Ex1() {
	blocks := extractData()

	size := len(blocks)
	lowestHeatLost := math.MaxInt
	visited := make([]byte, size*size)
	for i := 0; i < len(visited); i++ {
		visited[i] = '~'
	}
	total := traverse(blocks, visited, size, 1, 0, 1, east, 0, lowestHeatLost, 0)

	for i := 0; i < len(visited); i++ {
		visited[i] = '~'
	}
	total2 := traverse(blocks, visited, size, 0, 1, 1, east, 0, lowestHeatLost, 0)

	// total -= blocks[0][0]
	fmt.Printf("\n\n  Total: %d", total)
	fmt.Printf("\n\n  Total2: %d", total2)

	fmt.Printf("\n results: \n%v", results)
}

func traverse(blocks [][]int, visited []byte, size int, row, col int, straightMoves int, direct direction, heatLost, lowestHeatLost int, deep int) int {
	deep++

	for {
		if heatLost >= lowestHeatLost {
			// fmt.Printf("\n1. returns: %3d, %3d => %v,", row, col, lowestHeatLost)
			return lowestHeatLost
		}

		if straightMoves > 3 || outsizeRange(size, row, col) {
			// fmt.Printf("\n2. returns: %3d, %3d => %v,", row, col, lowestHeatLost)
			return lowestHeatLost
		}

		if visited[row*size+col] == 'X' {
			// fmt.Printf("\n3. returns: %3d, %3d => %v,", row, col, lowestHeatLost)
			return lowestHeatLost
		}

		visited[row*size+col] = 'X'
		heatLost += blocks[row][col]

		// // print
		// time.Sleep(10 * time.Millisecond)
		// fmt.Println()
		// for row := 0; row < size; row++ {
		// 	fmt.Println()
		// 	for col := 0; col < size; col++ {
		// 		fmt.Print(string(visited[row*size+col]))
		// 		// if visited[row*size+col] == 'X' {
		// 		// 	fmt.Printf("X")
		// 		// } else {
		// 		// 	fmt.Printf("~")
		// 		// }
		// 	}
		// }

		if row == size-1 && col == size-1 { // finish

			// fmt.Printf("\nfinish: h: %5d, lowest: %d, deep: %d", heatLost, lowestHeatLost, deep)
			// for row := 0; row < size; row++ {
			// 	fmt.Println()
			// 	for col := 0; col < size; col++ {
			// 		fmt.Print(string(visited[row*size+col]))
			// 	}
			// }

			if heatLost < lowestHeatLost {
				lowestHeatLost = heatLost
			}
			results = append(results, lowestHeatLost)
			// fmt.Printf("\n4. returns: %3d, %3d => %v,", row, col, lowestHeatLost)
			return lowestHeatLost
		}

		currHeatLost := math.MaxInt
		var tmpVisited []byte

		// GO:

		if direct != west {
			tmpVisited = make([]byte, size*size)
			copy(tmpVisited, visited)
			moves := 1
			if direct == east {
				moves = straightMoves + 1
			}
			currHeatLost = traverse(blocks, tmpVisited, size, row, col+1, moves, east, heatLost, lowestHeatLost, deep)
			if currHeatLost < lowestHeatLost {
				lowestHeatLost = currHeatLost
			}
		}

		if direct != north {
			tmpVisited = make([]byte, size*size)
			copy(tmpVisited, visited)
			moves := 1
			if direct == south {
				moves = straightMoves + 1
			}
			currHeatLost = traverse(blocks, tmpVisited, size, row+1, col, moves, south, heatLost, lowestHeatLost, deep)
			if currHeatLost < lowestHeatLost {
				lowestHeatLost = currHeatLost
			}
		}

		if direct != east {
			tmpVisited = make([]byte, size*size)
			copy(tmpVisited, visited)
			moves := 1
			if direct == west {
				moves = straightMoves + 1
			}
			currHeatLost = traverse(blocks, tmpVisited, size, row, col-1, moves, west, heatLost, lowestHeatLost, deep)
			if currHeatLost < lowestHeatLost {
				lowestHeatLost = currHeatLost
			}
		}

		// tmpVisited = make([]bool, size*size)
		// if direct != south {
		// 	copy(tmpVisited, visited)
		// 	moves := 1
		// 	if direct == north {
		// 		moves = straightMoves + 1
		// 	}
		// 	currHeatLost = traverse(blocks, tmpVisited, size, row-1, col, moves, north, heatLost, lowestHeatLost, deep)

		// 	if currHeatLost < lowestHeatLost {
		// 		lowestHeatLost = currHeatLost
		// 	}
		// }

		if direct == south {
			// fmt.Printf("\n5. returns: %3d, %3d => %v,", row, col, lowestHeatLost)
			return lowestHeatLost
		}

		if direct == north {
			straightMoves++
		} else {
			straightMoves = 1
		}

		direct = north
		row--
	}
}

func Ex2() {

}

func outsizeRange(size, row, col int) bool {

	if row < 0 || col < 0 || row == size || col == size {
		return true
	}

	return false
}
