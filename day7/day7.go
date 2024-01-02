package day7

import (
	"bufio"
	"fmt"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
	"golang.org/x/exp/slices"
)

const CARDS_IN_HAND = 5

type handCards [CARDS_IN_HAND]byte

func (c handCards) String() string {
	return fmt.Sprintf("%s", c[:])
}

type handType int

const (
	highCard = iota + 1
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var HAND_TYPE_DESC = [...]string{
	"highCard",
	"onePair",
	"twoPair",
	"threeOfAKind",
	"fullHouse",
	"fourOfAKind",
	"fiveOfAKind",
}

type gameHand struct {
	cards    handCards
	handType handType
	bid      int
}

func (h gameHand) String() string {
	return fmt.Sprintf("{cards: %v handType: %d, bid: %4d}", h.cards, h.handType, h.bid)
}

func extractData() []gameHand {
	test := 1

	filename := "./day7/data"
	if test == 0 {
		filename += "_t2"
	}

	filename += ".txt"
	f, err := os.Open(filename)
	help.IfErr(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	hands := make([]gameHand, 0, 100)
	for scanner.Scan() {
		var line []byte = scanner.Bytes()
		lineSize := len(line)

		hand := gameHand{}
		i := lineSize - 1
		numLen := 0
		hand.bid, numLen = help.ReadNumValueFromEnd(line, i)
		i -= numLen
		i-- // space

		for k := CARDS_IN_HAND - 1; k >= 0; k-- {
			hand.cards[k] = line[i]
			i--
		}

		hand.handType = getHandType(hand.cards)

		// fmt.Printf("\n cards: %2v, type: %d - %s", hand.cards, hand.handType, HAND_TYPE_DESC[hand.handType])

		hands = append(hands, hand)
	}

	return hands
}

func getHandType(cards handCards) handType {
	counts := make([]int, CARDS_IN_HAND)

	for i := 0; i < CARDS_IN_HAND; i++ {
		card := cards[i]
		for k := 0; k < CARDS_IN_HAND; k++ {
			if cards[k] == card {
				counts[i]++
			}
		}
	}

	// fmt.Printf("\ncounts: %2v", counts)

	if counts[0] == 5 {
		return fiveOfAKind
	}
	if counts[0] == 4 || counts[1] == 4 {
		return fourOfAKind
	}

	if counts[0] == 3 || counts[1] == 3 || counts[2] == 3 {
		for i := 0; i < CARDS_IN_HAND; i++ {
			if counts[i] == 2 {
				return fullHouse
			}
		}
		return threeOfAKind
	}

	pairs := 0
	for i := 0; i < CARDS_IN_HAND; i++ {
		if counts[i] == 2 {
			pairs++
		}
	}

	if pairs == 4 {
		return twoPair
	}

	if pairs == 2 {
		return onePair
	}

	return highCard
}

func Ex1() {

	hands := extractData()

	sortHands(hands)

	// handsCount := len(hands)
	// for i := 0; i < handsCount; i++ {
	// 	fmt.Printf("%s %d\n", hands[i].cards, hands[i].bid)
	// }
	total := calcTotal(hands)

	// 247878057 - too low
	fmt.Printf("\n\nTotal: %d", total)
}

func calcTotal(hands []gameHand) uint64 {
	total := uint64(0)
	handsCount := len(hands)
	for i := 0; i < handsCount; i++ {
		total += uint64(hands[i].bid * (i + 1))
		fmt.Printf("\n%v / %d, adding: %4d * %4d = %6d, t = %7d", hands[i].cards, hands[i].handType, hands[i].bid, i+1, hands[i].bid*(i+1), total)
	}

	return total
}

var LABELS = [...]byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

const LABELS_COUNT = len(LABELS)

func sortHands(hands []gameHand) {

	slices.SortFunc(hands, func(a, b gameHand) int {
		if a.handType > b.handType {
			return 1
		}

		if a.handType < b.handType {
			return -1
		}

		// equal type

		for i := 0; i < CARDS_IN_HAND; i++ {
			if a.cards[i] == b.cards[i] {
				continue
			}

			ca := a.cards[i]
			cb := b.cards[i]
			for k := 0; k < LABELS_COUNT; k++ {
				// labels won't be eq so if we fist find label a then it has lesser rank
				if LABELS[k] == ca {
					return -1
				}

				if LABELS[k] == cb {
					return 1
				}
			}

			return 0
		}

		return 0
	})
}

func Ex2() {

}
