package springs

import (
	"fmt"
	"strconv"
	"strings"
)

func NewSpringSet(line string) SpringSet {
	thisSpringSet := SpringSet{}
	work := strings.Split(line, " ")
	thisSpringSet.Springs = []rune(work[0])

	splitNums := strings.Split(work[1], ",")
	for i := 0; i < len(splitNums); i++ {
		snum, err := strconv.Atoi(splitNums[i])
		if err != nil {
			panic("did not get expected number")
		}
		thisSpringSet.Inst = append(thisSpringSet.Inst, snum)
	}
	return thisSpringSet
}

func PrintSpringSet(s SpringSet) {
	fmt.Printf("Springs: ")
	for _, v := range s.Springs {
		fmt.Printf("%c", v)
	}
	fmt.Printf(" Instructions: ")
	first := true
	for _, v := range s.Inst {
		if first {
			first = false
		} else {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n")
}
