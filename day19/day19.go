package day19

import (
	"bufio"
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
	lt         ruleCondition = '<'
	gt         ruleCondition = '>'
	alwaysTrue ruleCondition = 'Y'

	reject = byte('R')
	accept = byte('A')
)

type rule struct {
	categ       partCategory
	condition   ruleCondition
	value       int
	destination string
}

func extractData() (any, any) {

	f, err := os.Open("./day19/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	workflows := make(Workflows, 530)
	ratings := make([]Rating, 0, 202)

	readingRatings := false
	for scanner.Scan() {
		line := scanner.Bytes()
		lineSize := len(line)

		if readingRatings || lineSize == 0 {
			readingRatings = true
			r := parseRatingLine(line, lineSize)
			ratings = append(ratings, r)
			continue
		}

		name, rules := parseWorkflow(line, lineSize)
		workflows[name] = rules

	}

	help.IfErr(scanner.Err())

	return nil, nil
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
	for i := lineSize - 1; i >= rulesStartIdx; i-- {
		// TODO: start here, reading values inside {} from the end
	}

	return name, rules
}

func parseRatingLine(line []byte, lineSize int) Rating {
	r := Rating{}

	var num, offset int
	for i := lineSize - 1; i > 0; i-- {
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
	_, _ := extractData()
}

func Ex2() {

}
