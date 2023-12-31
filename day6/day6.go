package day6

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/unnamedxaer/aoc2023/help"
)

type race struct {
	time     int
	distance int
}

func extractDataEx1() []race {
	f, err := os.Open("./day6/data_t.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = scanner.Bytes()
	// Time:      7  15   30
	// Distance:  9  40  200

	races := []race{}

	for i := len(line) - 1; line[i] != ':'; i-- {
		if !help.IsNumber(line[i]) {
			continue
		}
		multiplier := 1
		v := 0
		for ; help.IsNumber(line[i]); i-- {
			v += int(line[i]-'0') * multiplier
			multiplier *= 10
		}
		currentRace := race{time: v}
		races = append(races, currentRace)
	}

	scanner.Scan()
	line = scanner.Bytes()

	racesCnt := len(races)
	rIdx := 0
	for i := len(line) - 1; rIdx < racesCnt; i-- {
		if !help.IsNumber(line[i]) {
			continue
		}
		multiplier := 1
		v := 0
		for ; help.IsNumber(line[i]); i-- {
			v += int(line[i]-'0') * multiplier
			multiplier *= 10
		}
		races[rIdx].distance = v
		rIdx++
	}

	// for debuggability
	slices.Reverse(races)

	return races
}

func Day6ex1() {
	races := extractDataEx1()
	fmt.Printf("\n\n races: %#v", races)

	total := calcTotalWinCombinationsV2(races)

	// total := calcTotalWinCombinationsV1(races)

	fmt.Printf("\ntotal: %d", total)
}

func calcTotalWinCombinationsV2(races []race) int {
	total := 1
	for i := 0; i < len(races); i++ {
		currentRace := races[i]
		j := 1
		for ; j < currentRace.time/2; j++ {
			if isWinPossible(j, currentRace) {
				break
			}
		}
		raceWinPossibilitiesCnt := currentRace.time - 2*j + 1
		fmt.Printf("\nfor race:\n%+v\npossibilities:\n%v", currentRace, raceWinPossibilitiesCnt)
		total *= max(1, raceWinPossibilitiesCnt)
	}
	return total
}

func calcTotalWinCombinationsV1(races []race) int {
	total := 1
	for i := 0; i < len(races); i++ {
		currentRace := races[i]
		winPossibilities := make([]race, 0, currentRace.time)

		raceWinPossibilitiesCnt := 0
		j := 1
		for ; j < currentRace.time; j++ {

			winPossibilities, raceWinPossibilitiesCnt = checkIfChargingTimeWinnable(j, currentRace, winPossibilities, raceWinPossibilitiesCnt)

		}

		total *= max(1, raceWinPossibilitiesCnt)
		fmt.Printf("\nfor race: %+v -> possibilities:\n%v", currentRace, raceWinPossibilitiesCnt)
	}
	return total
}

func isWinPossible(chargingTime int, currentRace race) bool {
	dist := calcDistance(chargingTime, currentRace.time)
	return dist > currentRace.distance
}

func checkIfChargingTimeWinnable(chargingTime int, currentRace race, winPossibilities []race, raceWinPossibilitiesCnt int) ([]race, int) {
	dist := calcDistance(chargingTime, currentRace.time)
	if dist > currentRace.distance {

		winPossibilities = append(winPossibilities, race{time: chargingTime, distance: dist})
		raceWinPossibilitiesCnt++
	}
	return winPossibilities, raceWinPossibilitiesCnt
}

func calcSpeed(timeHold int, raceTime int) int {
	// if timeHold == 0 {
	// 	return 0
	// }
	if timeHold >= raceTime {
		return 0
	}

	return timeHold
}

func calcDistance(timeCharging int, raceTime int) int {

	moveTime := raceTime - timeCharging

	if moveTime == 0 {
		return 0
	}

	speed := calcSpeed(timeCharging, raceTime)

	return speed * moveTime
}

func calcMoveTime(timeCharging int, raceTime int) int {
	return raceTime - timeCharging
}

func extractDataEx2() race {
	f, err := os.Open("./day6/data.txt")
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	var line []byte = scanner.Bytes()
	// Time:      7  15   30
	// Distance:  9  40  200

	multiplier := 1
	time := 0
	for i := len(line) - 1; line[i] != ':'; i-- {
		if !help.IsNumber(line[i]) {
			continue
		}
		for ; help.IsNumber(line[i]); i-- {
			time += int(line[i]-'0') * multiplier
			multiplier *= 10
		}
	}

	scanner.Scan()
	line = scanner.Bytes()

	multiplier = 1
	distance := 0
	for i := len(line) - 1; line[i] != ':'; i-- {
		if !help.IsNumber(line[i]) {
			continue
		}
		for ; help.IsNumber(line[i]); i-- {
			distance += int(line[i]-'0') * multiplier
			multiplier *= 10
		}
	}

	mainRace := race{time, distance}

	return mainRace
}

func Day6ex2() {
	mainRace := extractDataEx2()

	total := calcTotalWinCombinationsV2([]race{mainRace})

	fmt.Printf("\n race: %+v, total: %d", mainRace, total)
}
