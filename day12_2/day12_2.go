package main

import (
	"bufio"
	"day12_2/springs"
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

	var mySprings []springs.SpringSet
	// Insert code here
	for _, v := range lines {
		mySprings = append(mySprings, springs.NewSpringSet(v))
	}

	fmt.Printf("Folding 5x...\n")

	var tot int = 0
	for _, v := range mySprings {
		modSpring := springs.FoldSpringSet(v, 5)
		springs.PrintSpringSet(modSpring)
		combo := springs.GetCombinations(modSpring.Springs, modSpring.Inst)
		fmt.Printf("Combined: %d\n\n", combo)
		tot += combo
	}
	fmt.Printf("Total: %d\n", tot)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
