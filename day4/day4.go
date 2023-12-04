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
	// For each line, we'll pass it into the game interpreter to spit out a game object

	var myGames []Card

	for _, line := range lines {
		g := ParseGame(line)
		PrintGame(g)
		myGames = append(myGames, g)
	}

	cardSet := make(map[CardGames]int)

	// Now let's calculate
	var score int = 0
	for _, g := range myGames {
		score += g.GetScore()

		// first, let's see how many times we have
		// to loop over ourselves.
		loop, ok := cardSet[CardGames{ID: g.ID}]
		if !ok {
			loop = 1
		} else {
			loop++
		}
		// fmt.Printf("We have %d copies of this game #%d.\n", loop, g.ID)
		for i := 0; i < loop; i++ {

			// Now let's get the amount of games we won
			workingWinners := g.GetCopies()
			tWinners := TruncateIfNeeded(myGames, workingWinners)

			for _, v := range tWinners {
				_, ok := cardSet[CardGames{ID: v}]
				if ok {
					cardSet[CardGames{ID: v}]++
				} else {
					cardSet[CardGames{ID: v}] = 1
				}
			}
		}
	}
	// fmt.Printf("CardSet: %+v\n", cardSet)
	fmt.Printf("The total score for Part 1 is: %d\n", score)

	var totalCount int = 0
	for _, v := range cardSet {
		totalCount += v
	}
	// Now the originals too, this may not work
	totalCount += len(myGames)

	fmt.Printf("Total count of games we have for Part 2 is: %d\n", totalCount)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
