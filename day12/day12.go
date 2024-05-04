package main

import (
	"bufio"
	"day12/spring"
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

	var springs []spring.SpringRow
	// Insert code here
	for _, line := range lines {
		newListing := spring.MultiplyRecords(line, 5)
		springs = append(springs, spring.NewSpringRow(newListing))
	}

	// springs[0].PrintSpringRow()

	// // let's make a guess
	// x := spring.Guess{'#', '#', '.', '.', '#', '#', '#'}
	// fmt.Printf("Is ##..### compatible with the first springset: %+v\n", spring.IsCompatible(springs[0], x))

	// for i := 0; i < 500; i++ {
	// 	fmt.Printf("Here's the result of %d: %s\n", i, spring.Playtest(i))
	// }
	var sum int = 0
	for _, v := range springs {
		v.PrintSpringRow()
		val := spring.GuessEverything(v)
		fmt.Printf("Amount of compatible guesses: %d\n", val)
		sum += val
	}
	fmt.Printf("Total counts: %d\n", sum)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
