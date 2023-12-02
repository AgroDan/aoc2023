package main

import (
	"fmt"
	"strconv"
	"strings"
)

// m := make(map[string]int)
// // m["zero"] = 0
// m["one"]   = 1
// m["two"]   = 2
// m["three"] = 3
// m["four"]  = 4
// m["five"]  = 5
// m["six"]   = 6
// m["seven"] = 7
// m["eight"] = 8
// m["nine"]  = 9

var wordNumbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type WordIndex struct {
	idx  int
	word string
}

func (w WordIndex) GetNumber() int {
	for i, v := range wordNumbers {
		if v == w.word {
			return i
		}
	}
	return -1
}

func GetLowestWordIndex(w []WordIndex) int {
	lowestIdx := -1
	for i, v := range w {
		if lowestIdx < 0 {
			lowestIdx = i
		} else if w[lowestIdx].idx > v.idx {
			lowestIdx = i
		}
	}
	return lowestIdx
}

func GetHighestWordIndex(w []WordIndex) int {
	highestIdx := -1
	for i, v := range w {
		if highestIdx < 0 {
			highestIdx = i
		} else if w[highestIdx].idx < v.idx {
			highestIdx = i
		}
	}
	return highestIdx
}

func findAllOccurrences(s, substr string) []int {
	var indexes []int
	startIndex := 0

	for {
		index := strings.Index(s[startIndex:], substr)
		if index == -1 {
			break
		}
		indexes = append(indexes, startIndex+index)
		startIndex += index + 1
	}
	return indexes
}

func ReturnFirstAndLastNum(line string) (first, last int) {
	// Returns the first and last number in the string, whether it's
	// the word or the number.

	// Find the leftmost and rightmost integers
	leftMostNum, rightMostNum := -1, -1
	for i := 0; i < len(line); i++ {
		myChar := string(line[i])
		_, e := strconv.Atoi(myChar)
		if e == nil {
			// this is a number
			if leftMostNum < 0 {
				leftMostNum = i
			}

			if rightMostNum < 0 {
				rightMostNum = i
			}

			if i > rightMostNum {
				rightMostNum = i
			}
		}
	}
	words := []WordIndex{}
	// now, leftMost and rightMost are indexes of integers only.
	for _, v := range wordNumbers {
		idx := findAllOccurrences(line, v)
		for i := 0; i < len(idx); i++ {
			words = append(words, WordIndex{idx[i], v})
		}
		// indexes
		// idx := strings.Index(line, v)
		// if idx >= 0 {
		// 	words = append(words, WordIndex{idx, v})
		// 	fmt.Printf("Hit -")
		// }
	}
	fmt.Println()

	leftMostWord := GetLowestWordIndex(words)
	rightMostWord := GetHighestWordIndex(words)

	if leftMostWord < 0 {
		// most likely an integer was found
		first, _ = strconv.Atoi(string(line[leftMostNum]))
	} else if leftMostNum < 0 {
		// most likely a word was found
		first = words[leftMostWord].GetNumber()
	} else { // otherwise we have to check which is smallest
		if words[leftMostWord].idx < leftMostNum {
			first = words[leftMostWord].GetNumber()
		} else {
			first, _ = strconv.Atoi(string(line[leftMostNum]))
		}
	}

	if rightMostWord < 0 {
		// most likely an integer was found
		last, _ = strconv.Atoi(string(line[rightMostNum]))
	} else if rightMostNum < 0 {
		last = words[rightMostWord].GetNumber()
	} else {
		if words[rightMostWord].idx > rightMostNum {
			last = words[rightMostWord].GetNumber()
		} else {
			last, _ = strconv.Atoi(string(line[rightMostNum]))
		}
	}

	// fmt.Printf("Struct: %+v\n", words)
	// fmt.Printf("leftNum: %d, rightNum: %d, leftWord: %d, rightWord: %d\n", leftMostNum, rightMostNum, leftMostWord, rightMostWord)
	return
}

// func ReturnFirstNum(line string) int {
// 	// given the string, it will find the first number it finds

// 	firstInt := -1
// 	// first we'll check for integers
// 	for i := 0; i < len(line); i++ {
// 		myChar := string(line[i])
// 		_, e := strconv.Atoi(myChar)
// 		if e == nil {
// 			firstInt = i
// 			break
// 		}
// 	}

// 	var idxNum []int
// 	for i, v := range wordNumbers {
// 		idx := strings.Index(line, v)
// 		if idx >= 0 {
// 			idxNum = append(idxNum, idx)
// 		}
// 	}

// 	// lowestIdx := -1
// 	// whichNum := -1
// 	// // Now find the first word it finds
// 	// for i, v := range wordNumbers {
// 	// 	idx := strings.Index(line, v)
// 	// 	if idx >= 0 {
// 	// 		// fmt.Printf("Found %s at idx %d\n", v, idx)
// 	// 		if lowestIdx >= 0 && idx < lowestIdx {
// 	// 			lowestIdx = idx
// 	// 			whichNum = i
// 	// 		}
// 	// 	}
// 	// }

// 	if firstInt >= 0 && firstInt < lowestIdx {
// 		// actual num is the first up
// 		// convert and return
// 		myChar := string(line[firstInt])
// 		byteToInt, _ := strconv.Atoi(myChar)
// 		return byteToInt
// 	} else if lowestIdx >= 0 && lowestIdx < firstInt {
// 		// a number came first. convert.
// 		return whichNum
// 	} else {
// 		return -1
// 	}
// }
