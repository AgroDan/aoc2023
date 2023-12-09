package puzzleparser

import (
	"day9/sequence"
	"strconv"
	"strings"
)

func ParseSequence(line string) sequence.Sequence {
	var numStack []int
	numSplit := strings.Split(strings.Trim(line, " "), " ")
	for i := 0; i < len(numSplit); i++ {
		n, err := strconv.Atoi(numSplit[i])
		if err != nil {
			continue
		}
		numStack = append(numStack, n)
	}
	return sequence.NewSequence(numStack)
}
