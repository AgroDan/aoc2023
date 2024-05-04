package spring

import (
	"fmt"
	"strings"
)

// This is just some helper functions that don't relate to springs really

func MultiplyRecords(line string, amount int) string {
	// what a bitch
	splitString := strings.Split(line, " ")

	var records, groups string
	var first bool = true
	for i := 0; i < amount; i++ {
		if first {
			first = false
		} else {
			records += "?"
			groups += ","
		}
		records += splitString[0]
		groups += splitString[1]
	}
	return fmt.Sprintf("%s %s", records, groups)
}
