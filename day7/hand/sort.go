package hand

import "fmt"

// This file will handle sorting of stuff.
func bubbleSort(cardHand []rune) []rune {
	// I'm going to use bubblesort because why not, it's not a huge data pool
	for i := 0; i < len(cardHand)-1; i++ {
		for j := 0; j < len(cardHand)-1; j++ {
			if cardValue[cardHand[j]] < cardValue[cardHand[j+1]] {
				cardHand[j], cardHand[j+1] = cardHand[j+1], cardHand[j]
			}
		}
	}
	return cardHand
}

// This file will handle sorting of stuff.
func bubbleSortPart2(cardHand []rune) []rune {
	// I'm going to use bubblesort because why not, it's not a huge data pool
	for i := 0; i < len(cardHand)-1; i++ {
		for j := 0; j < len(cardHand)-1; j++ {
			if part2CardValue[cardHand[j]] < part2CardValue[cardHand[j+1]] {
				cardHand[j], cardHand[j+1] = cardHand[j+1], cardHand[j]
			}
		}
	}
	return cardHand
}

func MergeSortPartTwo(m []Hand) []Hand {
	// GOOD OL' MERGE SORT BABY
	if len(m) == 1 {
		return m
	} else {
		half := len(m) / 2
		left := m[:half]
		right := m[half:]
		lsort := MergeSortPartTwo(left)
		rsort := MergeSortPartTwo(right)
		return mergePartTwo(lsort, rsort)
	}
}

func mergePartTwo(left, right []Hand) []Hand {
	var res []Hand
	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}
		if CompareHandsBinaryPartTwo(left[0], right[0]) {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		for {
			if len(left) == 0 {
				break
			}
			res = append(res, left[0])
			left = left[1:]
		}
	}
	if len(right) > 0 {
		for {
			if len(right) == 0 {
				break
			}
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func MergeSort(m []Hand) []Hand {
	// GOOD OL' MERGE SORT BABY
	if len(m) == 1 {
		return m
	} else {
		half := len(m) / 2
		left := m[:half]
		right := m[half:]
		lsort := MergeSort(left)
		rsort := MergeSort(right)
		return merge(lsort, rsort)
	}
}

func merge(left, right []Hand) []Hand {
	var res []Hand
	for {
		if len(left) == 0 || len(right) == 0 {
			break
		}
		if CompareHandsBinary(left[0], right[0]) {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		for {
			if len(left) == 0 {
				break
			}
			res = append(res, left[0])
			left = left[1:]
		}
	}
	if len(right) > 0 {
		for {
			if len(right) == 0 {
				break
			}
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func CompareHandsBinaryPartTwo(h1, h2 Hand) bool {
	/*
	 * This is similar to the below, but doesn't return an ordered
	 * set. Rather, it returns a false if h1 is greater than h2,
	 * true if it isn't. This is mostly so I can use mergesort
	 * like a big boy.
	 *
	 * Remember, [0] if 1 is greater, [1] if 2 is greater
	 */
	if h1.result > h2.result {
		return false
	}

	if h1.result < h2.result {
		return true
	}

	// Otherwise it's the same hand. Compare each card then I guess.
	for i := 0; i < len(h1.cards); i++ {
		if part2CardValue[h1.cards[i]] > part2CardValue[h2.cards[i]] {
			return false
		} else if part2CardValue[h1.cards[i]] < part2CardValue[h2.cards[i]] {
			return true
		} else {
			continue
		}
	}

	// if you got here, then this is the same hand value. Otherwise I would
	// just return either in any direction, but since order matters I'm
	// going to just panic here so I can debug.
	fmt.Printf("Card1: %+v, Card2: %+v\n", h1, h2)
	h1.PrintHand()
	h2.PrintHand()
	panic("Weird logic!")
}

func CompareHandsBinary(h1, h2 Hand) bool {
	/*
	 * This is similar to the below, but doesn't return an ordered
	 * set. Rather, it returns a false if h1 is greater than h2,
	 * true if it isn't. This is mostly so I can use mergesort
	 * like a big boy.
	 *
	 * Remember, [0] if 1 is greater, [1] if 2 is greater
	 */
	if h1.result > h2.result {
		return false
	}

	if h1.result < h2.result {
		return true
	}

	// Otherwise it's the same hand. Compare each card then I guess.
	for i := 0; i < len(h1.cards); i++ {
		if cardValue[h1.cards[i]] > cardValue[h2.cards[i]] {
			return false
		} else if cardValue[h1.cards[i]] < cardValue[h2.cards[i]] {
			return true
		} else {
			continue
		}
	}

	// if you got here, then this is the same hand value. Otherwise I would
	// just return either in any direction, but since order matters I'm
	// going to just panic here so I can debug.
	fmt.Printf("Card1: %+v, Card2: %+v\n", h1, h2)
	h1.PrintHand()
	h2.PrintHand()
	panic("Weird logic!")
}

func CompareHandsCamelStyle(h1, h2 Hand) [2]Hand {
	/*
	 * Behold the result of my wonton disregard for actually reading the
	 * instructions as defined in the challenge. I spent a large amount of
	 * time writing the "CompareHandsPokerStyle()" function thinking that
	 * it was to be ordered based on actual Poker rules of highest card
	 * wins within a stalemate. But CamelPoker or whatever works differntly,
	 * and significantly easier to work with. If two of the same _type_ of
	 * hands are dealt, regardless of the value of those hands, you simply
	 * check each card in the hand one by one for the highest value. That
	 * one wins. So if one hand is JJJTT and the second is QQ333, the second
	 * hand wins because the first card is compared with the first in the
	 * second hand, which is J < Q, so QQ333 wins. So weird, but that's the
	 * rule. The worst part is the sample data gave me the same result with
	 * the poker rules defined below.
	 *
	 * Anyway, this returns lesser hand [0], greater hand [1].
	 */
	if h1.result > h2.result {
		return [2]Hand{h2, h1}
	}

	if h1.result < h2.result {
		return [2]Hand{h1, h2}
	}

	// Otherwise it's the same hand. Compare each card then I guess.
	for i := 0; i < len(h1.cards); i++ {
		if cardValue[h1.cards[i]] > cardValue[h2.cards[i]] {
			return [2]Hand{h2, h1}
		} else if cardValue[h1.cards[i]] < cardValue[h2.cards[i]] {
			return [2]Hand{h1, h2}
		} else {
			continue
		}
	}

	// if you got here, then this is the same hand value. Otherwise I would
	// just return either in any direction, but since order matters I'm
	// going to just panic here so I can debug.
	fmt.Printf("Card1: %+v, Card2: %+v\n", h1, h2)
	h1.PrintHand()
	h2.PrintHand()
	panic("Weird logic!")
}

func CompareHandsPokerStyle(h1, h2 Hand) [2]Hand {
	// This will take two hands and determine which is the greater hand.
	// The greater hand will be in result[1], and the lesser will be result[0]

	// Lets do some quick wins.
	if h1.result > h2.result {
		return [2]Hand{h2, h1}
	}

	if h1.result < h2.result {
		return [2]Hand{h1, h2}
	}

	if h1.result == HIGH_CARD && h2.result == HIGH_CARD {
		// find out which is the high card
		// we know there aren't any pairs, so just loop on down.
		for i := 0; i < len(h1.nonPairs); i++ {
			if cardValue[h1.nonPairs[i]] > cardValue[h2.nonPairs[i]] {
				return [2]Hand{h2, h1}
			} else if cardValue[h1.nonPairs[i]] < cardValue[h2.nonPairs[i]] {
				return [2]Hand{h1, h2}
			} else {
				continue
			}
		}

		// if we somehow got here, the hands are equal. So just return them.
		return [2]Hand{h1, h2}
	}

	if h1.result == ONE_PAIR && h2.result == ONE_PAIR {
		// one pair, get the pairs and decide quick
		var h1Pair, h2Pair rune
		for k := range h1.pairs {
			h1Pair = k
		}
		for k := range h2.pairs {
			h2Pair = k
		}
		if cardValue[h1Pair] > cardValue[h2Pair] {
			return [2]Hand{h2, h1}
		} else if cardValue[h1Pair] < cardValue[h2Pair] {
			return [2]Hand{h1, h2}
		} else {
			// now we'll just go down the list in order here
			for i := 0; i < len(h1.nonPairs); i++ {
				if cardValue[h1.nonPairs[i]] > cardValue[h2.nonPairs[i]] {
					return [2]Hand{h2, h1}
				} else if cardValue[h1.nonPairs[i]] < cardValue[h2.nonPairs[i]] {
					return [2]Hand{h1, h2}
				} else {
					continue
				}
			}
		}
	}

	// Now let's deal with the issue of 2 pairs being equal
	if h1.result == TWO_PAIR && h2.result == TWO_PAIR {
		// I don't care to make this syntactically nice, just do it
		var h1PairVals []int
		var h2PairVals []int
		for k := range h1.pairs {
			h1PairVals = append(h1PairVals, cardValue[k])
		}
		for k := range h2.pairs {
			h2PairVals = append(h2PairVals, cardValue[k])
		}

		// Now here's some slapped together logic, deal with it
		if h1PairVals[0] > h2PairVals[0] && h1PairVals[0] > h2PairVals[1] ||
			h1PairVals[1] > h2PairVals[0] && h1PairVals[1] > h2PairVals[1] {
			// We did it, h1 is greater
			return [2]Hand{h2, h1}
		} else if h1PairVals[0] < h2PairVals[0] && h1PairVals[0] < h2PairVals[1] ||
			h1PairVals[1] < h2PairVals[0] && h1PairVals[1] < h2PairVals[1] {
			return [2]Hand{h1, h2}
		} else {
			// Otherwise they've gotta be equal. Close the loop with the high card
			if cardValue[h1.nonPairs[0]] >= cardValue[h2.nonPairs[0]] {
				// if they're both equal then it doesn't matter
				return [2]Hand{h2, h1}
			} else {
				return [2]Hand{h1, h2}
			}
		}
	} // yikes

	// Now lets do hands that are 3 of a kind
	if h1.result == THREE_OF_A_KIND && h2.result == THREE_OF_A_KIND {
		// Compare the 3-of-a-kind hands first
		var h1Threes, h2Threes rune
		for k := range h1.pairs {
			h1Threes = k
		}
		for k := range h2.pairs {
			h2Threes = k
		}
		if cardValue[h1Threes] > cardValue[h2Threes] {
			return [2]Hand{h2, h1}
		} else if cardValue[h1Threes] < cardValue[h2Threes] {
			return [2]Hand{h1, h2}
		} else {
			// Otherwise the TOAK values are the same. Go by high card.
			if (cardValue[h1.nonPairs[0]] > cardValue[h2.nonPairs[0]] &&
				cardValue[h1.nonPairs[0]] > cardValue[h2.nonPairs[1]]) ||
				(cardValue[h1.nonPairs[1]] > cardValue[h2.nonPairs[1]] &&
					cardValue[h1.nonPairs[1]] > cardValue[h2.nonPairs[1]]) {
				return [2]Hand{h2, h1}
			} else {
				return [2]Hand{h1, h2}
			}
		}
	}

	// what a crap show. now let's worry about full houses.

	if h1.result == FULL_HOUSE && h2.result == FULL_HOUSE {
		// Find which cards have the 3-of-a-kind

		var h13kind, h1otherPair rune
		var h23kind, h2otherPair rune

		for k, v := range h1.pairs {
			if v > 1 {
				h13kind = k
			} else {
				h1otherPair = k
			}
		}
		for k, v := range h1.pairs {
			if v > 1 {
				h23kind = k
			} else {
				h2otherPair = k
			}
		}

		// first find the high pair vals
		if cardValue[h13kind] > cardValue[h23kind] {
			return [2]Hand{h2, h1}
		} else if cardValue[h13kind] < cardValue[h23kind] {
			return [2]Hand{h1, h2}
		} else {
			// otherwise there is a tie. FIX THE TIE
			if cardValue[h1otherPair] >= cardValue[h2otherPair] {
				return [2]Hand{h2, h1}
			} else {
				return [2]Hand{h1, h2}
			}
		}
	}

	// The rest shouldn't be so difficult.
	if h1.result == FOUR_OF_A_KIND && h2.result == FOUR_OF_A_KIND {
		var h1ThisHand, h2ThisHand rune
		for k := range h1.pairs {
			h1ThisHand = k
		}
		for k := range h2.pairs {
			h2ThisHand = k
		}

		if cardValue[h1ThisHand] > cardValue[h2ThisHand] {
			return [2]Hand{h2, h1}
		} else if cardValue[h1ThisHand] < cardValue[h2ThisHand] {
			return [2]Hand{h1, h2}
		} else {
			// equal 4-of-a-kinds, differentiate on the outlier
			if cardValue[h1.nonPairs[0]] >= cardValue[h2.nonPairs[0]] {
				return [2]Hand{h2, h1}
			} else {
				return [2]Hand{h1, h2}
			}
		}
	}

	// finally, 5 of a kind
	if h1.result == FIVE_OF_A_KIND && h2.result == FIVE_OF_A_KIND {
		var h1HighCard, h2HighCard rune
		for k := range h1.pairs {
			h1HighCard = k
		}
		for k := range h2.pairs {
			h2HighCard = k
		}

		if cardValue[h1HighCard] >= cardValue[h2HighCard] {
			return [2]Hand{h2, h1}
		} else {
			return [2]Hand{h1, h2}
		}
	}

	// If you got here, something went horribly wrong.
	fmt.Printf("Card1: %+v, Card2: %+v\n", h1, h2)
	h1.PrintHand()
	h2.PrintHand()
	panic("Weird logic!")
}

func SortSlice(myHands []Hand) []Hand {
	// This is thanks to chatgpt I guess, thanks for
	// the bubblesort again
	n := len(myHands)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			sortedPair := CompareHandsCamelStyle(myHands[j], myHands[j+1])
			myHands[j], myHands[j+1] = sortedPair[0], sortedPair[1]
		}
	}
	return myHands
}
