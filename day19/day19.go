package day19

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

type ratingCategory byte

const (
	extremely   ratingCategory = 'x'
	musical     ratingCategory = 'm'
	aerodynamic ratingCategory = 'a'
	shiny       ratingCategory = 's'
)

func (pc ratingCategory) String() string {
	return string(pc)
}

type Workflows map[string][]rule
type PartRatings struct {
	x, m, a, s int
}

const startWorkflowName = "in"

type ruleCondition byte
type definedAction byte

func (rc ruleCondition) String() string {
	return string(rc)
}

func (da definedAction) String() string {
	return string(da)
}

const (
	lt             ruleCondition = '<'
	gt             ruleCondition = '>'
	emptyCondition ruleCondition = 0

	reject = byte('R')
	accept = byte('A')
)

type rule struct {
	reject      bool
	accept      bool
	categ       ratingCategory
	condition   ruleCondition
	value       int
	destination string
}

func (r rule) String() string {
	if r.condition == emptyCondition {
		if r.accept {
			return "{A}"
		}
		if r.reject {
			return "{R}"
		}
		return "{" + r.destination + "}"
	}

	s := fmt.Sprintf("{%s%s%d:", r.categ, r.condition, r.value)
	if r.accept {
		s += "A}"
	} else if r.reject {
		s += "R}"
	} else {
		s += r.destination + "}"
	}

	return s
}

func extractData() (Workflows, []PartRatings) {

	f, err := os.Open("./day19/data.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	workflows := make(Workflows, 530)
	ratings := make([]PartRatings, 0, 202)

	readingRatings := false
	for scanner.Scan() {
		line := scanner.Bytes()
		lineSize := len(line)

		if lineSize == 0 {
			readingRatings = true
			continue
		}

		if readingRatings {
			r := parsePartRatingsLine(line, lineSize)
			ratings = append(ratings, r)
			continue
		}

		name, rules := parseWorkflow(line, lineSize)
		workflows[name] = rules

	}

	help.IfErr(scanner.Err())

	return workflows, ratings
}

func parseWorkflow(line []byte, lineSize int) (string, []rule) {

	i := 1
	for ; i < lineSize; i++ {
		if line[i] == '{' {
			break
		}
	}
	name := string(line[0:i])

	rules := make([]rule, 0, 5)
	rulesStartIdx := i + 1
	rulesBytes := line[rulesStartIdx : lineSize-1]

	rawRules := bytes.Split(rulesBytes, []byte{','})
	for _, rr := range rawRules {
		// this will mean that we have only 'A' or 'R', so no further processing needed.
		// we have destination without condition, and it's real destination, not a 'A' or 'R'
		// move after the ':'
		r := parseRule(rr)
		rules = append(rules, r)
	}

	return name, rules
}

func parseRule(rr []byte) rule {
	// fmt.Printf("\n%s", string(rr))
	rule := rule{}

	if rr[0] == reject {
		rule.reject = true
		return rule
	}
	if rr[0] == accept {
		rule.accept = true
		return rule
	}

	size := len(rr)
	i := size - 1
	for ; i >= 0; i-- {
		if rr[i] == ':' {
			break
		}
	}
	if i == -1 {
		rule.destination = string(rr)
		return rule
	}

	i++

	if rr[i] == reject {
		rule.reject = true
	} else if rr[i] == accept {
		rule.accept = true
	} else {
		rule.destination = string(rr[i:])
	}

	rule.categ = ratingCategory(rr[0])
	rule.condition = ruleCondition(rr[1])

	numVal, _ := help.ReadNumValueFromEnd(rr, i-1-1)
	rule.value = numVal
	return rule
}

func parsePartRatingsLine(line []byte, lineSize int) PartRatings {
	r := PartRatings{}

	var num, offset int
	for i := lineSize - 2; i > 0; i-- {
		num, offset = help.ReadNumValue(line, lineSize, i)
		i -= offset

		i-- // skip comma
		categ := ratingCategory(line[i])
		switch categ {
		case extremely:
			r.x = num
		case musical:
			r.m = num
		case aerodynamic:
			r.a = num
		case shiny:
			r.s = num

		default:
			panic(fmt.Sprintf("\n unknown part category: %+v, %s", categ, string(categ)))
		}

		i--
	}

	return r
}

func Ex1() {
	workflows, ratings := extractData()
	// for destination, v := range workflows {
	// 	fmt.Printf("\n%5s:", destination)
	// 	for _, v := range v {
	// 		fmt.Printf("\n\t%s", v)
	// 	}
	// }
	// fmt.Println()
	// for _, v := range ratings {
	// 	fmt.Printf("\n%+v", v)
	// }
	// fmt.Println()
	total := calcTotalRatingsOfAcceptedParts(workflows, ratings)

	fmt.Printf("\n\n total: %d", total)

}

func calcTotalRatingsOfAcceptedParts(workflows Workflows, partsRatings []PartRatings) int {
	total := 0

	for _, partRatings := range partsRatings {
		ok := pipePartThroughWorkflows(partRatings, workflows)
		if ok {
			total += partRatings.x + partRatings.m + partRatings.a + partRatings.s
		}
	}

	return total
}

func pipePartThroughWorkflows(partRatings PartRatings, workflows Workflows) bool {
	// fmt.Println()
	currWorkflowName := startWorkflowName
	cnt := -1
	for {
		cnt++
		currWorkflow, ok := workflows[currWorkflowName]
		if !ok {
			panic("\nwe ended up without new workflow")
		}
		// fmt.Printf("\n%4d. %s: %+v", cnt, currWorkflowName, currWorkflow)
		currWorkflowName = "" // panic, otherwise we could be trapped in infinite loop

	currWorkflowsLoop:
		for _, r := range currWorkflow {
			// fmt.Printf("\n%s", r)
			if r.condition == emptyCondition {
				if r.accept {
					return true
				} else if r.reject {
					return false
				} else {
					currWorkflowName = r.destination
					break currWorkflowsLoop
				}
			}

			switch r.condition {
			case lt:
				if _, ok := isLt(r, partRatings); ok {
					if r.accept {
						return true
					} else if r.reject {
						return false
					}
					currWorkflowName = r.destination
					break currWorkflowsLoop
				}
			case gt:
				if _, ok := isGt(r, partRatings); ok {
					if r.accept {
						return true
					} else if r.reject {
						return false
					}
					currWorkflowName = r.destination
					break currWorkflowsLoop
				}
			default:
				fmt.Printf("\nunsupported condition: %s in rule: %+v, workflow: %s", r.condition, r, currWorkflowName)
				return false
			}

			continue currWorkflowsLoop
		}
	}

	// return false
}

func isLt(r rule, partRatings PartRatings) (int, bool) {
	switch r.categ {
	case extremely:
		return partRatings.x, partRatings.x < r.value
	case musical:
		return partRatings.m, partRatings.m < r.value
	case aerodynamic:
		return partRatings.a, partRatings.a < r.value
	case shiny:
		return partRatings.s, partRatings.s < r.value
	default:
		fmt.Printf("\n unknown categ: %s, in rule: %+v", r.categ, r)
		return 0, false
	}
}

func isGt(r rule, partRatings PartRatings) (int, bool) {
	switch r.categ {
	case extremely:
		return partRatings.x, partRatings.x > r.value
	case musical:
		return partRatings.m, partRatings.m > r.value
	case aerodynamic:
		return partRatings.a, partRatings.a > r.value
	case shiny:
		return partRatings.s, partRatings.s > r.value
	default:
		fmt.Printf("\n unknown categ: %s, in rule: %+v", r.categ, r)
		return 0, false
	}
}
