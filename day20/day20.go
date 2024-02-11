package day20

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

// const inputNameSuffix= ""
const inputNameSuffix = "_t"
const inputName = "./day20/data" + inputNameSuffix + ".txt"

type moduleType byte

const (
	flipFlop    moduleType = '%'
	conjunction moduleType = '&'
	broadcaster moduleType = 'b'
)

func (mt moduleType) String() string {
	return string(mt)
}

type module struct {
	name         string
	mtype        moduleType
	destinations []string
}

func (m module) String() string {
	return fmt.Sprintf("module:{name: %s, mtype: %s destinations: %v}", m.name, m.mtype, m.destinations)
}

func extractData() []module {

	f, err := os.Open(inputName)
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	modules := make([]module, 0, 60)
	for scanner.Scan() {
		line := scanner.Bytes()

		i := 0

		mt := moduleType(line[i])
		if mt != broadcaster {
			i++
		}

		spaceIdx := slices.Index(line, ' ')

		m := module{
			mtype: mt,
			name:  string(line[i:spaceIdx]),
		}

		i = spaceIdx + 4

		d := bytes.Split(line[i:], []byte{',', ' '})
		m.destinations = make([]string, len(d))
		for i, d := range d {
			m.destinations[i] = string(d)
		}

		modules = append(modules, m)
	}

	help.IfErr(scanner.Err())

	return modules
}

func Ex1() {
	modules := extractData()
	fmt.Printf("\nlen: %d", len(modules))

	for _, m := range modules {
		fmt.Printf("\n%s", m)
	}
}
