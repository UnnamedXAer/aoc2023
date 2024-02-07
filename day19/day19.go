package day19

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

type partCategory byte

const (
	extremely   partCategory = 'x'
	musical     partCategory = 'm'
	aerodynamic partCategory = 'a'
	shiny       partCategory = 's'
)

type Workflows map[string][]rule
type Rating struct {
	x, m, a, s int
}

const startWorkflowName = "in"

type ruleCondition byte
type definedAction byte

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
	categ       partCategory
	condition   ruleCondition
	value       int
	destination string
}

func extractData() (Workflows, []Rating) {

	f, err := os.Open("./day19/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	workflows := make(Workflows, 530)
	ratings := make([]Rating, 0, 202)

	readingRatings := false
	for scanner.Scan() {
		line := scanner.Bytes()
		lineSize := len(line)

		if lineSize == 0 {
			readingRatings = true
			continue
		}

		if readingRatings {
			r := parseRatingLine(line, lineSize)
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

	rule.categ = partCategory(rr[0])
	rule.condition = ruleCondition(rr[1])

	numVal, _ := help.ReadNumValueFromEnd(rr, i-1-1)
	rule.value = numVal
	return rule
}

func parseRatingLine(line []byte, lineSize int) Rating {
	r := Rating{}

	var num, offset int
	for i := lineSize - 2; i > 0; i-- {
		num, offset = help.ReadNumValue(line, lineSize, i)
		i -= offset

		i-- // skip comma
		categ := partCategory(line[i])
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
	for _, v := range workflows {
		fmt.Println()
		for _, v := range v {
			if v.condition == emptyCondition {
				fmt.Printf("\nno Condition: %+v", v)
			}
			fmt.Printf("\n%+v", v)
		}
	}
	fmt.Println()
	for _, v := range ratings {
		fmt.Printf("\n%+v", v)
	}
	// fmt.Printf("\n%+v, \n%+v", a, b)
}

func Ex2() {

}
