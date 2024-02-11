package day20

import (
	"bufio"
	"bytes"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

// const inputNameSuffix= ""
const inputNameSuffix = "_t"
const inputName = "./day20/data" + inputNameSuffix + ".txt"

func extractData() Platform {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	modules := make(Platform, 60)
	for scanner.Scan() {
		line := scanner.Bytes()

		i := 0

		mt := moduleType(line[i])
		if mt != broadcaster {
			i++
		}

		spaceIdx := slices.Index(line, ' ')

		name := string(line[i:spaceIdx])

		i = spaceIdx + 4
		d := bytes.Split(line[i:], []byte{',', ' '})

		destinations := make([]string, len(d))
		for i, d := range d {
			destinations[i] = string(d)
		}

		modules[name] = newModule(name, mt, destinations)
	}

	help.IfErr(scanner.Err())

	return modules
}

func newModule(name string, mt moduleType, destinations []string) Module {

	m := module{name, mt, destinations, low}

	switch mt {
	case conjunction:
		return &conjunctionModule{
			m:       m,
			history: make(map[string]modulePulse, 5),
		}
	case flipFlop:

		return &flipFlopModule{
			m: m,
		}
	case broadcaster:
		return &broadcasterModule{
			m: m,
		}
	default:
		panic("unknown module type: " + mt.String())
	}
}
