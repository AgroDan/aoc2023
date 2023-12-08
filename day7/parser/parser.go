package parser

import (
	"day7/hand"
	"strconv"
	"strings"
)

func ParseHand(l string) hand.Hand {
	// given a line, will return a hand object
	// line example:
	// 32T3K 765
	game := strings.Split(strings.Trim(l, " "), " ")
	bid, _ := strconv.Atoi(game[1])
	return hand.NewHand(game[0], bid)
}

func ParseHandPartTwo(l string) hand.Hand {
	// given a line, will return a hand object
	// line example:
	// 32T3K 765
	game := strings.Split(strings.Trim(l, " "), " ")
	bid, _ := strconv.Atoi(game[1])
	return hand.NewHandPartTwo(game[0], bid)
}
