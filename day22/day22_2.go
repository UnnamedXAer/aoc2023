package day22

import (
	"fmt"

	"github.com/unnamedxaer/aoc2023/help"
)

func Ex2() {
	bricks := extractData()
	// for _, b := range bricks {
	// 	fmt.Printf("\n%v", b)
	// }

	// fmt.Println()
	// printBricksFromPerspective(bricks, pX, pZ)
	// printBricksFromPerspective(bricks, pY, pZ)
	// printBricksFromPerspective(bricks, pZ, pX)
	// printBricksFromPerspective(bricks, pZ, pY)
	// printBricksFromPerspective(bricks, pX, pY)

	bricks = fallBricks(bricks)

	// printBricksFromPerspective(bricks, pX, pZ)
	// printBricksFromPerspective(bricks, pY, pZ)
	// printBricksFromPerspective(bricks, pX, pY)

	supports, supportedBy := determineWhatSupportWhat(bricks)
	// total := calcDisintegrable(bricks, supports, supportedBy)

	total := calcChainReactions(bricks, supports, supportedBy)

	fmt.Printf("\n\n Total: %d", total)
}

func calcChainReactions(bricks []*brick, supports map[*brick][]*brick, supportedBy map[*brick][]*brick) int {

	total := 0
	for _, b1 := range bricks {
		// fmt.Println(b1)
		total += calcChainReactionForBrick(b1, supports, supportedBy)
	}
	// total = calcChainReactionForBrick(bricks, bricks[0], supports, supportedBy)

	return total
}

func calcChainReactionForBrick(b *brick, supports map[*brick][]*brick, supportedBy map[*brick][]*brick) int {

	q := help.NewQAny[*brick]()
	q.Push(b)

	depends := map[*brick]bool{
		b: true,
	}

	for !q.IsEmpty() {
		b1 := q.Pop()

		list := supports[b1]
		for _, b2 := range list {

			if !isDepending(b, b2, supportedBy, depends) {
				continue
			}

			depends[b2] = true

			q.Push(b2)
		}
	}

	return len(depends) - 1
}

func isDepending(base *brick, b *brick, supportedBy map[*brick][]*brick, depends map[*brick]bool) bool {

	for _, supporter := range supportedBy[b] {
		// if supporter.id == base.id {
		// 	continue
		// }

		if depends[supporter] {
			continue
		}

		return false
	}

	return true
}
