package main

import (
	"bufio"
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
	s := NewSchematic(lines)
	s.PrintSchematic()

	var allNumbers []Block

	var total int = 0
	// Loop through the Y coords
	for i, v := range s.mapping {
		nums := FindNumbersOnY(i, v)
		for _, num := range nums {
			// loop for each number on the line
			if s.IsPartNumber(num) {
				// fmt.Printf("Part number: %d\n", s.GetNumber(num))
				total += s.GetNumber(num)
			}
			// now push the number onto the allNumbers slice
			allNumbers = append(allNumbers, num)
		}
	}

	fmt.Printf("Total part numbers: %d\n", total)

	// Now let's get the gear ratios. First, find the asterisks.

	stars := s.FindStars()

	// Now let's check which of these stars are gears

	gears := s.FindGears(stars, allNumbers)

	// Now lets loop through the gears and find the gear ratios
	var totalGearRatios int = 0
	for _, g := range gears {
		ratio := s.GetGearRatio(g)
		fmt.Printf("Ratio: %d\n", ratio)
		totalGearRatios += ratio
	}

	fmt.Printf("Adding up all gear ratios returns: %d\n", totalGearRatios)

	// Just care about the 25th row
	// b := FindNumbersOnY(25, s.mapping[25])

	// for _, v := range b {
	// 	v.PrintBlock()
	// 	fmt.Printf("And the number representing the above: %d\n", s.GetNumber(v))
	// 	fmt.Printf("Is it a part number? %s\n", s.IsPartNumber(v))
	// }
	// fmt.Printf("%+v\n", s)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
