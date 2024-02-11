package day20

import (
	"fmt"
	"strings"
)

type Module interface {
	receivePulse(receivedFrom string, pulse modulePulse) modulePulse
	String() string
}

type Platform map[string]Module

func (p Platform) String() string {
	s := strings.Builder{}

	s.WriteString("\nPlatform:")
	for _, m := range p {
		s.WriteByte('\n')
		s.WriteString(m.String())
	}

	return s.String()
}

type moduleType byte

const (
	flipFlop    moduleType = '%'
	conjunction moduleType = '&'
	broadcaster moduleType = 'b'
)

type modulePulse uint8

const (
	low  modulePulse = 0
	high modulePulse = 1
)

func (mt moduleType) String() string {
	return string(mt)
}

type module struct {
	name         string
	mtype        moduleType
	destinations []string
	state        modulePulse
}

type broadcasterModule struct {
	m module
}

type flipFlopModule struct {
	m module
}

type conjunctionModule struct {
	m module
	// history stores the type of the most recent pulse received from each of their connected input modules;
	history map[string]modulePulse
}

func (m module) String() string {
	return fmt.Sprintf("{name: %11s, mtype: %s destinations: %v}", m.name, m.mtype, m.destinations)
}

func (m *broadcasterModule) String() string {
	return fmt.Sprintf("broadcasterModule :{m: %s}", m.m)
}
func (m *flipFlopModule) String() string {
	return fmt.Sprintf("flipFlopModule    :{m: %s}", m.m)
}
func (m *conjunctionModule) String() string {
	return fmt.Sprintf("conjunctionModule :{m: %s, history: %v}", m.m, m.history)
}

func (m *broadcasterModule) receivePulse(from string, pulse modulePulse) modulePulse {
	return pulse
}

func (m *flipFlopModule) receivePulse(from string, pulse modulePulse) modulePulse {
	if pulse == high {
		return m.m.state
	}

	if m.m.state == low {
		return high
	}
	return low
}

func (m *conjunctionModule) receivePulse(from string, pulse modulePulse) modulePulse {
	m.history[from] = pulse
	for _, p := range m.history {
		if p == low {
			return high
		}
	}
	return low
}

func Ex1() {
	modules := extractData()
	fmt.Printf("\nlen: %d", len(modules))
	fmt.Printf("\n%s", modules)

	// for _, m := range modules {
	// 	fmt.Printf("\n%s", m)
	// }
}

func process(modules Platform) {}
