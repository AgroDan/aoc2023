package main

import (
	"bufio"
	"day11/cosmos"
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
	myCosmos := cosmos.NewUniverse(lines)
	myCosmos.PrintUniverse()

	fmt.Println("Now expand the universe...")
	myCosmos.ExpandUniverse()
	myCosmos.PrintUniverse()

	fmt.Println("Get the pairs...")
	pairs := cosmos.GeneratePairs(myCosmos)

	fmt.Println("Now add everything up.")
	total := pairs.GetTotal()
	fmt.Printf("Here: %d\n", total)

	// Now let's add that million number
	myHugeCosmos := cosmos.NewUniverse(lines)
	fmt.Println("Expanding the universe by a tremendous amount...")
	myHugeCosmos.ArbitrarilyExpandUniverse(1_000_000)
	// myHugeCosmos.PrintUniverse()

	fmt.Printf("Getting pairs...\n")
	hugePairs := cosmos.GeneratePairs(myHugeCosmos)

	fmt.Println("Added up becomes...")
	hugeTotal := hugePairs.GetTotal()
	fmt.Printf("Grand total: %d\n", hugeTotal)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
