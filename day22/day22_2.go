package day22

import "fmt"

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
		total = calcChainReactionForBrick(bricks, b1, supports, supportedBy)
	}

	return total
}

func calcChainReactionForBrick(bricks []*brick, b1 *brick, supports map[*brick][]*brick, supportedBy map[*brick][]*brick) int {

	list := supports[b1]

	total := 0

	// bricksOnlySupportedByB1 := []*brick{}

	for _, b2 := range list {
		// len == 0 cannot happen

		// if b2 is only supported by 1 brick, that means that only b1 is supporting it,
		// so disintegrating b1 disintegrate b2, therefore we add 1 plus result of
		// disintegrating b2.
		if len(supportedBy[b2]) == 1 {
			total++
			total += calcChainReactionForBrick(bricks, b2, supports, supportedBy)
			// bricksOnlySupportedByB1 = append(bricksOnlySupportedByB1, b2)
		} else {
			// b2 is supported by some other(s) brick(s) than b1, but those bricks
			// potentially could be disintegrated as they may have a common parent with b2

			// in that case, do we have to keep list off all fallen bricks that have common parent
			// with b2 (which means that we need a list off all fallen bricks from starting from
			// the brick passed to the calcChainReaction function).
		}
	}

	return total
}
