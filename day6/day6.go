package main

import (
	"bufio"
	"day6/parser"
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
	r := parser.ParseRace(lines)
	r.PrintRace()
	fmt.Printf("Number required from part 1: %d\n", r.PuzzlePartOne())

	// Now handle the kerning problem
	fmt.Println("Kerning fixed:")
	r2 := parser.KerningParse(lines)
	r2.PrintRace()
	fmt.Printf("Number required for part 2: %d\n", r2.PuzzlePartOne())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
