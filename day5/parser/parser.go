package parser

import (
	"day5/seed"
	"strconv"
	"strings"
)

/*
 *	This will be a monster of a Parser. Will create an "Almanac" object
 * 	which should have all the data necessary.
 */

func ParseAlmanac(lines []string) seed.Almanac {
	a := seed.Almanac{}

	var instSlice []seed.Instruction

	for i := 0; i < len(lines); i++ {
		// First check for the seeds line
		if strings.Contains(lines[i], "seeds:") {
			// This is the seed initializer
			wordSplit := strings.Split(lines[i], ":")
			nums := strings.Split(strings.Trim(wordSplit[1], " "), " ")
			var numStack []int
			for _, v := range nums {
				// For all values
				n, e := strconv.Atoi(v)
				if e != nil {
					// Probably found a blank space or something
					continue
				}
				numStack = append(numStack, n)
			}
			a.Seeds = numStack
		} // at this point, we have the seeds.

		// Now get the maps.
		if strings.Contains(lines[i], "map:") {
			inst := seed.Instruction{}
			types := strings.Split(lines[i], " ")
			srcdst := strings.Split(types[0], "-")
			inst.From = srcdst[0]
			inst.To = srcdst[2]

			// Now let's loop over each ensuing number.
			checkCount := 0
			var rSlice []seed.Ratio
			for {
				checkCount++
				// First, let's check to see if we hit EOF
				if i+checkCount >= len(lines) {
					// yup, EOF
					break
				}
				// Second let's check if we hit a blank line
				wLine := strings.Trim(lines[i+checkCount], " ")
				if len(wLine) <= 0 {
					i += checkCount // move onto the next line
					break
				}
				instNumbers := strings.Split(wLine, " ")
				d, _ := strconv.Atoi(instNumbers[0])
				s, _ := strconv.Atoi(instNumbers[1])
				l, _ := strconv.Atoi(instNumbers[2])
				r := seed.Ratio{
					Dst:    d,
					Src:    s,
					Length: l,
				}
				rSlice = append(rSlice, r)
			}
			inst.R = rSlice
			instSlice = append(instSlice, inst)
		}
	}
	a.Inst = instSlice
	return a
}
