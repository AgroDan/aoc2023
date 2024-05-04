package spring

import (
	"math"
)

// This will have to do with attempts at a configuration.

type Guess []rune

func (g Guess) AmountBroken() int {
	// returns the amount of broken springs in the
	// guess itself
	var counter int = 0
	for _, v := range g {
		if v == '#' {
			counter++
		}
	}
	return counter
}

func IsCompatible(s SpringRow, g Guess) bool {
	// This simply returns if the provided guess
	// is compatible for basic checks.

	// fmt.Printf("Checking length...\n")
	// fmt.Printf("s.s: %d, g: %d\n", len(s.s), len(g))
	if len(s.s) != len(g) {
		return false
	}

	// fmt.Printf("Checking if sum adds up...\n")
	var sum int = 0
	for _, v := range s.instructions {
		sum += v
	}

	// fmt.Printf("Checking if amount broken adds up...\n")
	if sum != g.AmountBroken() {
		return false
	}

	// Now that we're done with the simple stuff, let's see if
	// this guess is compatible from the set of instructions.

	var checkInst int = 0   // index of instruction we are on
	var record bool = false // flag to confirm when we start recording successive string
	var buf int = 0         // amount of broken springs we found so far
	for i, v := range g {
		// fmt.Printf("Checking %c at idx %d\n", v, i)
		if record {

			if v == '#' && i < len(g)-1 {
				buf++
			} else {
				// we might be a the end of the line here
				// so process this next if statement
				if v == '#' {
					buf++
				}
				record = false

				if buf != s.instructions[checkInst] {
					// fmt.Printf("Failed while recording. cIdx: %d Buf: %d\n", checkInst, buf)
					return false
				} else {
					buf = 0
					checkInst++
				}
			}
		} else {
			if v == '#' {
				record = true
				buf = 1
				if i >= len(g)-1 {
					if buf != s.instructions[checkInst] {
						// fmt.Printf("Failed while not recording. cIdx: %d Buf: %d\n", checkInst, buf)
						return false
					}
				}
			}
		}
	}
	return true
}

func GuessEverything(s SpringRow) int {
	// This will attempt to guess everything and return the amount of
	// potential configurations.
	var counter int = 0
	possibilities := math.Pow(2, float64(len(s.qIdx)))

	// remember, the below function only generates possibilities for the
	// ? characters, you have to fill in those to generate a true guess!
	guesses := generatePossibilities(len(s.qIdx), int(possibilities))
	// fmt.Printf("Guesses: %+v\n", guesses)
	for _, v := range guesses {
		thisGuess := s.CopySpringRow()
		for i := 0; i < len(v); i++ {
			thisGuess[s.qIdx[i]] = v[i]
		}
		// fmt.Printf("Trying %+v...\n", thisGuess)
		if IsCompatible(s, thisGuess) {
			counter++
		}
	}
	return counter
}

func generatePossibilities(width, possibilities int) [][]rune {
	// This is an internal function which will generate a 2 dimensional
	// array of runes given the padding of the width function.
	var retval [][]rune
	for i := 0; i < possibilities; i++ {
		iter := i
		var val []rune
	innerLoop:
		for {
			if iter == 0 {
				break innerLoop // just for visual help i guess
			}

			testVal := iter & 1
			if testVal == 1 {
				val = append(val, '#')
			} else {
				val = append(val, '.')
			}
			iter >>= 1
		}

		// check to see if we need to crop or pad
		if len(val) > width {
			val = val[:width]
		} else if len(val) < width {
		padLoop:
			for {
				if len(val) >= width {
					break padLoop
				} else {
					val = append(val, '.')
				}
			}
		}
		retval = append(retval, val)
	}
	return retval
}

func Playtest(n int) string {
	retval := ""
	for {
		if n == 0 {
			break
		}

		testVal := n & 1
		if testVal == 1 {
			retval += "#"
		} else {
			retval += "."
		}
		n >>= 1
	}
	return retval
}
