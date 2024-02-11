package day20

import (
	"fmt"
	"strings"

	"github.com/unnamedxaer/aoc2023/help"
)

type Module interface {
	receivePulse(receivedFrom string, pulse modulePulse) (modulePulse, bool)
	destinations() []string
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

func (mp modulePulse) String() string {
	if mp == low {
		return "low"
	}

	return "high"
}

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
	return fmt.Sprintf("b :{m: %s}", m.m)
}
func (m *flipFlopModule) String() string {
	return fmt.Sprintf("f :{m: %s}", m.m)
}
func (m *conjunctionModule) String() string {
	return fmt.Sprintf("c :{m: %s, history: %v}", m.m, m.history)
}

func (m *broadcasterModule) destinations() []string {
	return m.m.destinations
}

func (m *flipFlopModule) destinations() []string {
	return m.m.destinations
}

func (m *conjunctionModule) destinations() []string {
	return m.m.destinations
}

func (m *broadcasterModule) receivePulse(from string, pulse modulePulse) (modulePulse, bool) {
	m.m.state = pulse
	return pulse, true
}

func (m *flipFlopModule) receivePulse(from string, pulse modulePulse) (modulePulse, bool) {
	if pulse == high {
		// m.m.state = high
		return m.m.state, false
	}

	if m.m.state == low {
		m.m.state = high
		return high, true
	}

	m.m.state = low
	return low, true
}

func (m *conjunctionModule) receivePulse(from string, pulse modulePulse) (modulePulse, bool) {
	m.history[from] = pulse
	for _, p := range m.history {
		if p == low {
			m.m.state = high
			return high, true
		}
	}
	m.m.state = low
	return low, true
}

func Ex1() {
	modules := extractData()
	// fmt.Printf("\nlen: %d", len(modules))
	// fmt.Printf("\n%s", modules)

	// for _, m := range modules {
	// 	fmt.Printf("\n%s", m)
	// }

	totalPushes := [2]int{0, 0}
	for i := 0; i < 1000; i++ {
		// fmt.Printf("\n\n______________RUN %4d", i)
		pushes := process(modules)
		totalPushes[low] += pushes[low]
		totalPushes[high] += pushes[high]
		// fmt.Printf("\n%v", modules)
	}

	fmt.Printf("\n\ntotal pushes: %v", totalPushes)
	total := totalPushes[low] * totalPushes[high]
	fmt.Printf("\ntotal: %d", total)
}

func process(modules Platform) [2]int {
	const startModuleName = "broadcaster"

	// printPlatformStates(modules)

	type stackItem struct {
		dst   string
		src   string
		pulse modulePulse
	}

	q := help.NewQAny[stackItem]()

	q.Push(stackItem{
		dst:   startModuleName,
		src:   "button",
		pulse: low,
	})

	pushes := [2]int{0, 0}

	for !q.IsEmpty() {
		item := q.Pop()

		pushes[item.pulse]++

		// fmt.Printf("\nreceived pulse: %11s -> %4s -> %s", item.src, item.pulse, item.dst)
		currModule, ok := modules[item.dst]
		if !ok {
			continue
		}

		nextPulseState, follow := currModule.receivePulse(item.src, item.pulse)

		if !follow {
			continue
		}

		destinations := currModule.destinations()
		for _, dst := range destinations {
			q.Push(stackItem{
				dst:   dst,
				src:   item.dst,
				pulse: nextPulseState,
			})
		}
	}

	return pushes
}

func printPlatformStates(modules Platform) {
	for _, v := range modules {
		var x module
		switch v.(type) {
		case *broadcasterModule:
			x = v.(*broadcasterModule).m
		case *flipFlopModule:
			x = v.(*flipFlopModule).m
		case *conjunctionModule:
			x = v.(*conjunctionModule).m
		}

		fmt.Printf("\n %11s %s %s", x.name, x.mtype, x.state)
	}
}
