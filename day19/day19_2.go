package day19

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
)

func Ex2() {

	workflows, _ := extractData()
	for destination, v := range workflows {
		fmt.Printf("\n%5s:", destination)
		for _, v := range v {
			fmt.Printf("\n\t%s", v)
		}
	}
	fmt.Println()
	// for _, v := range ratings {
	// 	fmt.Printf("\n%+v", v)
	// }
	// fmt.Println()
	total := calcTotalPossibilities(workflows)

	fmt.Printf("\n\n total: %d", total)
}

type ratingMinMax struct {
	min, max int
}

type tesseract struct {
	x, m, a, s ratingMinMax
}

func (t tesseract) clone() *tesseract {
	return &t
}

type ruleFilterResult struct {
	matched     *tesseract
	unmatched   *tesseract
	destination []rule
}

func calcTotalPossibilities(workflows Workflows) int {
	//

	tesseracts := make([]*tesseract, 0, 100)

	q := help.NewQAny[ruleFilterResult]()
	q.Push(ruleFilterResult{
		matched: &tesseract{
			x: ratingMinMax{1, 4000},
			m: ratingMinMax{1, 4000},
			a: ratingMinMax{1, 4000},
			s: ratingMinMax{1, 4000},
		},
		unmatched:   nil,
		destination: workflows[startWorkflowName],
	})

	for !q.IsEmpty() {
		item := q.Pop()
		currTesseract := item.matched
		workflow := item.destination

		for _, rule := range workflow {
			result := filterRule(rule, currTesseract, workflows)
			if rule.accept {
				if result.matched != nil {
					tesseracts = append(tesseracts, result.matched)
				}
				if result.unmatched != nil {
					currTesseract = result.unmatched
				}
				continue
			} else if rule.reject {
				if result.unmatched != nil {
					currTesseract = result.unmatched
				}
				continue
			}

			if result.unmatched != nil {
				currTesseract = result.unmatched
			}

			q.Push(ruleFilterResult{
				matched:     result.matched,
				destination: result.destination,
			})
		}
	}

	total := 0
	for _, t := range tesseracts {
		tesseractTotal := 1
		tesseractTotal *= t.x.max - t.x.min + 1
		tesseractTotal *= t.m.max - t.m.min + 1
		tesseractTotal *= t.a.max - t.a.min + 1
		tesseractTotal *= t.s.max - t.s.min + 1
		total += tesseractTotal
	}

	return total
}

func filterRule(rule rule, t *tesseract, workflows Workflows) ruleFilterResult {
	if rule.condition == emptyCondition {
		if rule.accept {
			return ruleFilterResult{
				matched:     t,
				unmatched:   nil,
				destination: nil,
			}
		}
		if rule.reject {
			return ruleFilterResult{
				matched:     nil,
				unmatched:   nil,
				destination: nil,
			}
		}

		return ruleFilterResult{
			matched:     t,
			unmatched:   nil,
			destination: workflows[rule.destination],
		}
	}

	var matched *tesseract
	var unmatched *tesseract
	if rule.condition == lt {
		switch rule.categ {
		case extremely:
			if rule.value < t.x.min {
				matched = nil
				unmatched = t
			} else if rule.value >= t.x.max {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.x.min = rule.value
				matched = t.clone()
				matched.x.max = rule.value - 1
			}
		case musical:
			if rule.value < t.m.min {
				matched = nil
				unmatched = t
			} else if rule.value >= t.m.max {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.m.min = rule.value
				matched = t.clone()
				matched.m.max = rule.value - 1
			}
		case aerodynamic:
			if rule.value < t.a.min {
				matched = nil
				unmatched = t
			} else if rule.value >= t.a.max {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.a.min = rule.value
				matched = t.clone()
				matched.a.max = rule.value - 1
			}
		case shiny:
			if rule.value < t.s.min {
				matched = nil
				unmatched = t
			} else if rule.value >= t.s.max {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.s.min = rule.value
				matched = t.clone()
				matched.s.max = rule.value - 1
			}
		}
	} else if rule.condition == gt {
		switch rule.categ {
		case extremely:
			if rule.value >= t.x.max {
				matched = nil
				unmatched = t
			} else if rule.value < t.x.min {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.x.max = rule.value
				matched = t.clone()
				matched.x.min = rule.value + 1
			}
		case musical:
			if rule.value >= t.m.max {
				matched = nil
				unmatched = t
			} else if rule.value < t.m.min {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.m.max = rule.value
				matched = t.clone()
				matched.m.min = rule.value + 1
			}
		case aerodynamic:
			if rule.value >= t.a.max {
				matched = nil
				unmatched = t
			} else if rule.value < t.a.min {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.a.max = rule.value
				matched = t.clone()
				matched.a.min = rule.value + 1
			}
		case shiny:
			if rule.value >= t.s.max {
				matched = nil
				unmatched = t
			} else if rule.value < t.s.min {
				matched = t
				unmatched = nil
			} else {
				unmatched = t.clone()
				unmatched.s.max = rule.value
				matched = t.clone()
				matched.s.min = rule.value + 1
			}
		}
	} else {
		panic("unknown condition: " + string(rule.condition))
	}

	resultLt := ruleFilterResult{
		matched:     matched,
		unmatched:   unmatched,
		destination: workflows[rule.destination],
	}

	return resultLt

}
