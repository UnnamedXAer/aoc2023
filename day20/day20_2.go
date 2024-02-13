package day20

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
)

func Ex2() {
	modules := extractData()
	// fmt.Printf("\nlen: %d", len(modules))
	// fmt.Printf("\n%s", modules)

	// for _, m := range modules {
	// 	fmt.Printf("\n%s", m)
	// }

	// here - here should be filled automatically, but currently is not.
	// here contains names of modules (conjunctions) that will send pulse to only one module
	// that sends pulse to out final module i.e 'rx'.
	// we need to send low to 'rx', and the preceding module is conjunction so it sends low
	// when all of its inputs sends high to him. so we need to find when they all "meet" with high value
	// apparently LMC is useful here.
	here := map[string][]int{
		"dh": []int{}, "qd": []int{}, "bb": []int{}, "dp": []int{},
	}

	totalPushes := [2]int{0, 0}
	i := 0
outer:
	for {
		i++
		// fmt.Printf("\n\n______________RUN %4d", i)
		pushes, completed := processPart2(modules, here)
		totalPushes[low] += pushes[low]
		totalPushes[high] += pushes[high]

		for _, c := range completed {
			v := here[c]
			n := len(v)
			if n < 1 {
				here[c] = append(v, i)
			}
		}

		for _, v := range here {
			if len(v) < 1 {
				continue outer
			}
		}

		break
	}

	i = -1
	indxes := make([]int, 0, len(here))
	for k, v := range here {
		if len(v) < 1 {
			fmt.Printf("we do not have enough values for %q module", k)
			break
		}
		indxes = append(indxes, v[0])
	}

	lmc := help.Lcm(indxes[0], indxes[1], indxes[2:]...)
	total := lmc

	fmt.Printf("\n%+v", here)

	fmt.Printf("\n\ntotal pushes: %v", totalPushes)
	fmt.Printf("\ntotal: %d", total)
}

func processPart2(modules Platform, here map[string][]int) ([2]int, []string) {
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
	out := make([]string, 0, len(here))

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
		if _, ok := here[item.dst]; ok && nextPulseState == high {
			out = append(out, item.dst)
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

	return pushes, out
}
