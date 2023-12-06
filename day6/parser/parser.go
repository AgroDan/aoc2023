package parser

import (
	"day6/race"
	"strconv"
	"strings"
)

/*
 * Luckily this isn't much to parse. Kind of nice considering
 * how much of my time the whole parsing problem takes up
 */

func ParseRace(lines []string) race.Race {
	var t, d []int
	for _, v := range lines {
		// I don't care about excessive strings. Just try to
		// convert each and go on about our day
		line := strings.Split(v, " ")
		for _, item := range line[1:] {
			n, e := strconv.Atoi(item)
			if e != nil {
				continue
			}
			if line[0] == "Time:" {
				t = append(t, n)
			} else if line[0] == "Distance:" {
				d = append(d, n)
			}
		}
	}
	return race.NewRace(t, d)
}

func KerningParse(lines []string) race.Race {
	var t, d []int
	for _, v := range lines {
		// I don't care about excessive strings. Just try to
		// convert each and go on about our day
		noSpaces := strings.ReplaceAll(v, " ", "")
		line := strings.Split(noSpaces, ":")
		for _, item := range line {
			n, e := strconv.Atoi(item)
			if e != nil {
				continue
			}
			if line[0] == "Time" {
				t = append(t, n)
			} else if line[0] == "Distance" {
				d = append(d, n)
			}
		}
	}
	return race.NewRace(t, d)
}
