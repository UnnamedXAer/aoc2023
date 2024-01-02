package day8

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/constraints"
)

type instructions []int
type network map[string][2]string

func extractData() (instr instructions, net network, starts []string) {

	f, err := os.Open("./day8/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	instr = make(instructions, len(line))
	for i, v := range line {
		if v == 'L' {
			instr[i] = 0
		} else {
			instr[i] = 1
		}
	}

	net = make(network)
	scanner.Scan()
	for scanner.Scan() {
		line = scanner.Text()
		key := line[:3]
		net[key] = [2]string{line[7:10], line[12:15]}

		if key[2] == 'A' {
			starts = append(starts, key)
		}

	}

	return instr, net, starts
}

func Ex1() {

	instr, net, _ := extractData()

	fmt.Printf("\ninstr: %v", instr)
	fmt.Printf("\nnet: \n%v", net)

	total := calcStepsEx1(instr, net)
	fmt.Printf("\nTotal steps: %d", total)
}

func calcStepsEx1(instrs instructions, net network) int {
	steps := 0
	currentPos := "AAA"
	dest := "ZZZ"

	for {
		for _, instr := range instrs {
			currentPos = net[currentPos][instr]
			steps++
			if currentPos == dest {
				return steps
			}
		}
	}
}

func Ex2() {

	instr, net, starts := extractData()

	fmt.Printf("\ninstr: %v", instr)
	fmt.Printf("\nnet: \n%v", net)
	fmt.Printf("\nstart: %v", starts)
	startAt := time.Now().UnixNano()
	total := calcStepsEx2(instr, net, starts)
	endAt := time.Now().UnixNano()
	calcTime := (endAt - startAt) / 1000
	fmt.Printf("\nTotal steps: %d, in: %d ms", total, calcTime)
}

func calcStepsEx2(instrs instructions, net network, starts []string) int {
	steps := 0
	startsCnt := len(starts)

	startsCopy := make([]string, startsCnt)
	copy(startsCopy, starts)

	firstZAt := make([]int, startsCnt)

	instrsCnt := len(instrs)
	if instrsCnt == 0 {
		panic("cannot move without instructions!")
	}
	zNode := ""
	i := 0
	// we fing should base the solution base on the assumption that there is a cycle in every path
	// and use LCM with lengths of each path....
	for k := 0; k < startsCnt; k++ {
		for ; ; i++ {
			instr := instrs[i%instrsCnt]
			// for _, instr := range instrs {
			steps++
			// for k := 0; k < startsCnt; k++ {
			starts[k] = net[starts[k]][instr]
			// fmt.Printf(" -> %v", starts[k])
			if starts[k][2] == 'Z' {
				zNode = starts[k]
				break
			}
		}
		stepsToFirstZ := steps
		steps = 0

		i++

		for ; ; i++ {
			instr := instrs[i%instrsCnt]
			steps++
			starts[k] = net[starts[k]][instr]
			// fmt.Printf(" -> %v", starts[k])
			if starts[k] == zNode {
				break
			}
		}

		stepsInLoop := steps
		firstZAt[k] = stepsInLoop

		// for c := 0; c < 4; c++ {
		// 	i++
		// 	steps = 0
		// 	for ; ; i++ {
		// 		instr := instrs[i%instrsCnt]
		// 		steps++
		// 		starts[k] = net[starts[k]][instr]
		// 		// fmt.Printf(" -> %v", starts[k])
		// 		if starts[k] == zNode {
		// 			break
		// 		}
		// 	}

		// 	if steps != stepsInLoop {
		// 		fmt.Printf("\nsteps in loop: %d, steps: %d", stepsInLoop, steps)
		// 		return -33
		// 	}
		// }
		fmt.Printf("\nfirst Z after: %d steps, loop from Z to Z: %v", stepsToFirstZ, stepsInLoop)
	}

	fmt.Printf("\nsteps z to z: %v", firstZAt)

	t := lcm(firstZAt[0], firstZAt[1], firstZAt...)

	return t
}

func gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm[T constraints.Integer](a, b T, integers ...T) T {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
