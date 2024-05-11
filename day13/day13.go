package main

import (
	"bufio"
	"day13/mirrors"
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
	ms := mirrors.GenerateMirrorSets(lines)
	for i, v := range ms {
		fmt.Printf("Iter %d:\n", i)
		v.PrintMirror()
		fmt.Printf("\n")
	}

	var MULT int = 100              // the multiplier for the first challenge
	var totRows, totCols int = 0, 0 // total rows, total cols

	// fmt.Printf("Rows first:\n")
	for _, v := range ms {
		// fmt.Printf("Map %d\n", i)
		for j := range v.M {
			// fmt.Printf("Checking %d and %d...\n", j, j+1)
			if v.CompareRows(j, j+1) {
				// fmt.Printf("Comparing refraction for rows %d and %d...\n", j, j+1)
				if v.IsRowRefraction(j) {
					// We have a perfect mirror for rows, add 1 to the row
					// because the challenge makers start counting rows at 1
					// and it's most likely because he's a perl developer
					// so of course things don't make sense
					totRows += ((j + 1) * MULT)
				}
			}
		}
	}

	// fmt.Printf("Now columns:\n")
	for _, v := range ms {
		// fmt.Printf("Map %d\n", i)
		for j := range v.M[0] {
			if v.CompareCols(j, j+1) {
				// fmt.Printf("Comparing refraction for cols %d and %d...\n", j, j+1)
				if v.IsColRefraction(j) {
					totCols += (j + 1)
				}
			}
		}
	}

	fmt.Printf("Total val of rows: %d\n", totRows)
	fmt.Printf("Total val of cols: %d\n", totCols)
	fmt.Printf("Total score: %d\n", totCols+totRows)

	// reset the counter
	totRows, totCols = 0, 0
	// Now for part 2, we only care if there is an off-by-one scenario.
	for _, v := range ms {
		for j := range v.M {
			if v.IsRowRefractionOffByOne(j) {
				totRows += ((j + 1) * MULT)
			}
		}
	}

	for _, v := range ms {
		for j := range v.M[0] {
			if v.IsColRefractionOffByOne(j) {
				totCols += (j + 1)
			}
		}
	}

	fmt.Printf("Total val of rows for part 2: %d\n", totRows)
	fmt.Printf("Total val of cols for part 2: %d\n", totCols)
	fmt.Printf("Total score for part 2: %d\n", totCols+totRows)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
