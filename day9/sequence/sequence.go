package sequence

import "fmt"

type Sequence struct {
	num [][]int
}

func NewSequence(n []int) Sequence {
	// constructor for sequence object.
	var temp [][]int
	temp = append(temp, n)
	s := Sequence{
		num: temp,
	}
	s.genDeltas()
	s.guessNext()
	// now for part 2...
	s.prepareSequences() // pad with zeroes
	s.goBackInTime()     // guess the numbers
	return s
}

func (s Sequence) PrintSequence() {
	fmt.Printf("Initial:\n")
	for i := 0; i < len(s.num[0]); i++ {
		fmt.Printf("%2d ", s.num[0][i])
	}
	if len(s.num) > 1 {
		fmt.Printf("\nDeltas:\n")
		for i := 1; i < len(s.num); i++ {
			for j := 0; j < len(s.num[i]); j++ {
				fmt.Printf("%2d ", s.num[i][j])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func (s *Sequence) genDeltas() {
	// This function will generate the deltas for each
	// sequence given up to the all zeroes sequence and
	// appends them to the origin sequence
	var iter int = 0
	for {
		if allZeroes(s.num[iter]) {
			break
		}

		var newNumStack []int
		for i := 0; i < len(s.num[iter]); i++ {
			if i+1 == len(s.num[iter]) {
				break
			}
			diff := s.num[iter][i+1] - s.num[iter][i]
			newNumStack = append(newNumStack, diff)
		}
		s.num = append(s.num, newNumStack)
		iter++
	}
}

func (s *Sequence) guessNext() {
	// This function will append the next guess of
	// numbers based on the given formula and add it
	// to each line in the sequence
	// for i := len(s.num) - 1; i >= 0; i-- {
	for i := 0; i < len(s.num); i++ {
		appNumber := addAllTopValues(*s, i, 0) // this may be wrong
		s.num[i] = append(s.num[i], appNumber)
	}
}

func allZeroes(num []int) bool {
	// checks to see if a given number stack
	// is all zeroes
	for _, v := range num {
		if v != 0 {
			return false
		}
	}
	return true
}

func addAllTopValues(s Sequence, place, total int) int {
	// recursive function to add all values in a sequence
	// and return the total
	if allZeroes(s.num[place]) {
		return 0
	}

	lastNum := s.num[place][len(s.num[place])-1]
	return lastNum + addAllTopValues(s, place+1, total)
}

func (s Sequence) GetPredictedValue() int {
	// Returns the predicted value. If you used the constructor
	// to create a sequence like you should, then this will work
	// without a problem.
	return s.num[0][len(s.num[0])-1]
}

func (s Sequence) GetPreceedingValue() int {
	return s.num[0][0]
}

func (s *Sequence) prepareSequences() {
	// This function does nothing more than prepare the sequences
	// by padding the first index with a 0.
	for i := 0; i < len(s.num); i++ {
		// for each row of numbers
		var padded []int
		padded = append(padded, 0)
		padded = append(padded, s.num[i]...) // tack on the rest of the slice
		s.num[i] = padded
	}
}

func (s *Sequence) goBackInTime() {
	// This function fills in the number _before_ the provided sequence.
	// It assumes that the slices have been pre-padded with a leading 0 first.
	// The formula it will follow is num[i][1] - num[i+1][0].
	for i := len(s.num) - 1; i >= 0; i-- {
		// I could start at the second to last, but just in case...
		if allZeroes(s.num[i]) {
			continue
		}
		s.num[i][0] = s.num[i][1] - s.num[i+1][0]
	}
}
