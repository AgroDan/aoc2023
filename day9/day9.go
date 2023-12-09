package main

import (
	"bufio"
	"day9/puzzleparser"
	"flag"
	"fmt"
	"os"
	"time"
)

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
	var totalNumbersPartOne int = 0
	var totalNumbersPartTwo int = 0
	for _, v := range lines {
		s := puzzleparser.ParseSequence(v)
		totalNumbersPartOne += s.GetPredictedValue()
		totalNumbersPartTwo += s.GetPreceedingValue()
		s.PrintSequence()
	}
	fmt.Printf("Total number from predicted values for Part 1: %d\n", totalNumbersPartOne)
	fmt.Printf("Total number from predicted values for Part 2: %d\n", totalNumbersPartTwo)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
