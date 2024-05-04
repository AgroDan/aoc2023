package spring

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type SpringRow struct {
	s            []rune
	instructions []int
	qIdx         []int // index of question marks
}

func NewSpringRow(line string) SpringRow {
	thisSpringRow := SpringRow{}
	splitString := strings.Split(line, " ")

	thisSpringRow.s = []rune(splitString[0])

	splitNums := strings.Split(splitString[1], ",")
	for i := 0; i < len(splitNums); i++ {
		snum, err := strconv.Atoi(splitNums[i])
		if err != nil {
			panic("did not get expected number")
		}
		thisSpringRow.instructions = append(thisSpringRow.instructions, snum)
	}

	for i := 0; i < len(thisSpringRow.s); i++ {
		if thisSpringRow.s[i] == '?' {
			thisSpringRow.qIdx = append(thisSpringRow.qIdx, i)
		}
	}
	return thisSpringRow
}

func (s SpringRow) CopySpringRow() []rune {
	// generates a copy of the springrow's rune array
	var c = make([]rune, len(s.s))
	_ = copy(c, s.s)
	return c
}

func (s SpringRow) PrintSpringRow() {
	fmt.Printf("Spring Row: ")
	for _, v := range s.s {
		fmt.Printf("%c", v)
	}

	var first bool = true
	fmt.Printf(" Set Instructions: ")
	for _, v := range s.instructions {
		if first {
			first = false
		} else {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", v)
	}

	fmt.Printf(" Idx of ?'s: ")
	first = true
	for _, v := range s.qIdx {
		if first {
			first = false
		} else {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", v)
	}

	x := float64(len(s.qIdx))
	possibilities := math.Pow(2, x)
	fmt.Printf(" Potential combinations: %0.f", possibilities)
	fmt.Printf("\n")
}
