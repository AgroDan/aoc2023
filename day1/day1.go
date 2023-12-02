package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func flIntegers(line string) (first, last int) {
	// This will take a string and return the first digit
	// number it finds in the string as well as the last
	// digit it finds.

	// first digit
	for i := 0; i < len(line); i++ {
		myChar := string(line[i])
		c, e := strconv.Atoi(myChar)
		if e == nil {
			first = c
			break
		}
	}

	// last digit
	for i := len(line); i > 0; i-- {
		myChar := string(line[i-1])
		c, e := strconv.Atoi(myChar)
		if e == nil {
			fmt.Printf("Got: %d\n", c)
			last = c
			break
		}
	}
	return
}

func mergeNumbers(numLeft, numRight int) int {
	// this function will take two numbers and merge
	// them into one double-digit number.
	return (numLeft * 10) + numRight
}

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here
	// var total int = 0
	// for _, v := range lines {
	// 	fmt.Printf("Working with %s\n", v)
	// 	first, last := flIntegers(v)
	// 	fmt.Printf("First num: %d, last num: %d\n", first, last)
	// 	mNumber := mergeNumbers(first, last)
	// 	total += mNumber
	// }
	// fmt.Printf("Grand total: %d\n", total)
	total := 0
	for _, v := range lines {
		fmt.Printf("Working with %s -- ", v)
		firstNum, lastNum := ReturnFirstAndLastNum(v)
		fmt.Printf("%d%d\n", firstNum, lastNum)
		// break
		mNumber := mergeNumbers(firstNum, lastNum)
		fmt.Printf("%d + %d = %d\n", total, mNumber, total+mNumber)
		total += mNumber
	}
	fmt.Printf("Grand total: %d\n", total)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
