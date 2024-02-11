package day20

import (
	"bufio"
	"bytes"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

const inputNameSuffix = ""

// const inputNameSuffix = "_t"
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

	for _, m := range modules {
		c, ok := m.(*conjunctionModule)
		if !ok {
			continue
		}

		for _, m := range modules {
			if !slices.Contains(m.destinations(), c.m.name) {
				continue
			}
			name := ""
			switch tmp := m.(type) {
			case *broadcasterModule:
				name = tmp.m.name
			case *flipFlopModule:
				name = tmp.m.name
			case *conjunctionModule:
				name = tmp.m.name
			default:
				panic("unknown type")
			}
			c.history[name] = low
		}
	}

	help.IfErr(scanner.Err())

	return modules
}

func newModule(name string, mt moduleType, destinations []string) Module {

	m := module{name, mt, destinations, low}

	switch mt {
	case conjunction:
		cm := &conjunctionModule{
			m:       m,
			history: make(map[string]modulePulse, 5),
		}

		return cm
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
