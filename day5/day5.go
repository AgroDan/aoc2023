package main

import (
	"bufio"
	"day5/parser"
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
	a := parser.ParseAlmanac(lines)
	a.PrintAlmanac()
	lowest := -1
	for _, s := range a.Seeds {
		endPoint := a.FindDestination(s)
		fmt.Printf("With Seed %d -> ", s)
		fmt.Printf("%d\n", endPoint)
		if lowest < 0 || endPoint < lowest {
			lowest = endPoint
		}
	}
	fmt.Printf("Lowest is %d\n", lowest)

	seedPairs := a.GenerateSeedRanges()
	// fmt.Printf("Here's a bunch of seed pairs: %+v\n", seedPairs)

	fmt.Printf("Now with the seed ranges, let's check everything.\n")
	newLow := -1
	for _, s := range seedPairs {
		// fmt.Printf("s: %+v, s0 - %d, s1 - %d\n", s, s[0], s[1])
		for i := s[0]; i <= s[0]+s[1]; i++ {
			// fmt.Printf("%d\n", i)
			endPoint := a.FindDestination(i)
			if newLow < 0 || endPoint < newLow {
				newLow = endPoint
				fmt.Printf("New low! %d from seed %d\n", newLow, s[0])
			}
		}
	}
	fmt.Printf("Part 2 new lowest: %d\n", newLow)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
