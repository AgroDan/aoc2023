package hand

import "fmt"

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var cardValue = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

var part2CardValue = map[rune]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type Hand struct {
	cards    []rune
	bid      int
	result   int // this is the result of what kind of hand it is
	pairs    map[rune]int
	nonPairs []rune
}

func NewHand(c string, b int) Hand {
	runes := []rune(c)
	h := Hand{
		cards: runes,
		bid:   b,
	}
	h.setHand()
	h.setNonPairs()
	return h
}

func NewHandPartTwo(c string, b int) Hand {
	runes := []rune(c)
	h := Hand{
		cards: runes,
		bid:   b,
	}
	h.setHandPartTwo()
	h.setNonPairsPartTwo()
	return h
}

func (h Hand) GetBid() int {
	// returns the Bid because I don't want to
	// refactor proper exported variable names
	return h.bid
}

func (h Hand) PrintHand() {
	fmt.Printf("Current Hand: %s, Current Bid: %d\n", string(h.cards), h.bid)
	fmt.Printf("Hand result: ")
	hr := "??"
	switch h.result {
	case HIGH_CARD:
		hr = "High Card"
	case ONE_PAIR:
		hr = "One Pair"
	case TWO_PAIR:
		hr = "Two Pair"
	case THREE_OF_A_KIND:
		hr = "Three of a Kind"
	case FULL_HOUSE:
		hr = "Full House"
	case FOUR_OF_A_KIND:
		hr = "Four of a Kind"
	case FIVE_OF_A_KIND:
		hr = "Five of a Kind"
	}
	var pairSlice []rune
	for k := range h.pairs {
		pairSlice = append(pairSlice, k)
	}
	fmt.Printf("%s, Known Pairs: %s, Remaining cards: %s\n", hr, string(pairSlice), string(h.nonPairs))
}

func (h *Hand) setHand() {
	// this function will determine the hand result

	// first construct the pair maps
	pairs := make(map[rune]int)

	// Now lets look for pairs.
	for i := 0; i < len(h.cards); i++ {
		cmp := h.cards[i] // for ease of reading I guess
		_, created := pairs[cmp]
		if created {
			// we already have the pairs for this card
			continue
		}
		pairs[cmp] = 0
		for j := i + 1; j < len(h.cards); j++ {
			if cmp == h.cards[j] {
				// we have a pair here. Record it.
				pairs[cmp]++
			}
		}
	}
	/* At this point, we have a map of all _unique_ pairs in
	 * the hand. If we add up how many different pairs we have,
	 * we can correlate it to the hand that we have as well.
	 * So two "one pair" hands are enumerated as 1 and 1,
	 * so 1 + 1 = 2, which is enumerated to two pair.
	 * If a "one pair" and a "three of a kind" are there, 1 + 3 = 4
	 * aka a full house. So add all these up, but first we need
	 * to determine the value of the type of pairs. If one card
	 * has 2 pairs, then technically (by the way I checked), that is
	 * a 3 of a kind. One card has 2 pairs, 4 of a kind, etc.
	 */
	// fmt.Printf("pairs: %+v\n", pairs)
	pairList := make(map[rune]int)
	for k, v := range pairs {
		if v > 0 {
			pairList[k] = v
		}
	}
	h.pairs = pairList

	// Now, for each pair, determine the value
	var intVal int = 0
	for _, v := range pairs {
		switch v {
		case 1:
			intVal += ONE_PAIR
		case 2:
			intVal += THREE_OF_A_KIND
		case 3:
			intVal += FOUR_OF_A_KIND
		case 4:
			intVal += FIVE_OF_A_KIND
		}
	}
	h.result = intVal
}

func (h *Hand) setHandPartTwo() {
	// this function will do the same as the original, only
	// now with part 2's new rule, it will account for J
	// being a joker now which will add to the fun. Hooray.
	pairs := make(map[rune]int)

	// lets look for _non joker pairs_
	jokerCount := 0
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] == 'J' {
			// found a joker, don't consider it
			// a pair. Just increment and continue
			jokerCount++
			continue
		}
		cmp := h.cards[i]
		_, created := pairs[cmp]
		if created {
			// already have pairs for this card
			continue
		}
		pairs[cmp] = 0
		for j := i + 1; j < len(h.cards); j++ {
			if h.cards[j] == 'J' {
				// ignore that joker
				continue
			}
			if cmp == h.cards[j] {
				// we have a pair, record it
				pairs[cmp]++
			}
		}
	}

	// Now lets determine what we have without jokers accounted for.
	pairList := make(map[rune]int)
	for k, v := range pairs {
		if v > 0 {
			pairList[k] = v
		}
	}
	h.pairs = pairList

	// Now that we have the pairlist, let's determine the hand without jokers:
	var intVal int = 0
	for _, v := range pairs {
		switch v {
		case 1:
			intVal += ONE_PAIR
		case 2:
			intVal += THREE_OF_A_KIND
		case 3:
			intVal += FOUR_OF_A_KIND
		case 4:
			intVal += FIVE_OF_A_KIND
		}
	}

	// Now do some "Joker math". This is weird.
	switch jokerCount {
	case 1: // only one joker, loads of possibilities
		switch intVal {
		case HIGH_CARD:
			// one joker, high card
			intVal = ONE_PAIR
		case ONE_PAIR:
			// better hand is always 3 of a kind
			intVal = THREE_OF_A_KIND
		case TWO_PAIR:
			intVal = FULL_HOUSE
		case THREE_OF_A_KIND:
			intVal = FOUR_OF_A_KIND
		case FOUR_OF_A_KIND:
			intVal = FIVE_OF_A_KIND
		}
	case 2: // two jokers, not as many
		switch intVal {
		case HIGH_CARD:
			intVal = THREE_OF_A_KIND
		case ONE_PAIR:
			intVal = FOUR_OF_A_KIND
		case THREE_OF_A_KIND:
			intVal = FIVE_OF_A_KIND
		}
	case 3: // three jokers, even fewer
		switch intVal {
		case HIGH_CARD:
			intVal = FOUR_OF_A_KIND
		case ONE_PAIR:
			intVal = FIVE_OF_A_KIND
		}
	case 4: // 4 jokers????
		// if there are 4 jokers then it can only be a
		// high card situation, which means the 4 jokers
		// will complement the remaining and make it
		// five of a kind
		intVal = FIVE_OF_A_KIND
	case 5:
		// obviously this is 5 of a kind
		intVal = FIVE_OF_A_KIND
	}
	h.result = intVal
}

func (h *Hand) setNonPairs() {
	// This function sets the slice of non-pairs ordered
	// from highest val to lowest
	dupeMap := make(map[rune]bool)

	// First, find the duplicates
	for _, v := range h.cards {
		if _, ok := dupeMap[v]; ok {
			dupeMap[v] = true
		} else {
			dupeMap[v] = false
		}
	}

	var result []rune
	for _, v := range h.cards {
		if !dupeMap[v] {
			result = append(result, v)
		}
	}

	// Now bubblesort this because I'm oldschool cool
	h.nonPairs = bubbleSort(result)
}

func (h *Hand) setNonPairsPartTwo() {
	// This is slightly different from the above, just
	// taking into account the new rules for part 2
	dupeMap := make(map[rune]bool)

	// first, find dupes
	for _, v := range h.cards {
		if _, ok := dupeMap[v]; ok {
			dupeMap[v] = true
		} else {
			dupeMap[v] = false
		}
	}

	var result []rune
	for _, v := range h.cards {
		if !dupeMap[v] {
			result = append(result, v)
		}
	}

	// and do a better bubblesort
	h.nonPairs = bubbleSortPart2(result)
}
