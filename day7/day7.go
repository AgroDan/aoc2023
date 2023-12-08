package main

import (
	"bufio"
	"day7/hand"
	"day7/parser"
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
	var gameList []hand.Hand
	for _, v := range lines {
		thisGame := parser.ParseHand(v)
		thisGame.PrintHand()
		fmt.Printf("\n")
		gameList = append(gameList, thisGame)
	}

	fmt.Printf("Now let's sort and display\n")
	// newList := hand.MergeSort(gameList)
	newList := hand.MergeSort(gameList)
	var totalBid int = 0
	for i, v := range newList {
		rankVal := i + 1
		sumMoney := v.GetBid() * rankVal
		totalBid += sumMoney
		fmt.Printf("Rank %d:\n", rankVal)
		v.PrintHand()
		fmt.Printf("\n")
	}
	fmt.Printf("Returned total bid value: %d\n", totalBid)

	fmt.Println("And now to accommodate for the new Joker cards, do all this again.")
	var partTwoGameList []hand.Hand
	for _, v := range lines {
		thisGame := parser.ParseHandPartTwo(v)
		partTwoGameList = append(partTwoGameList, thisGame)
	}
	fmt.Printf("Now let's sort and display\n")
	partTwoSortedList := hand.MergeSortPartTwo(partTwoGameList)
	var partTwoTotalBid int = 0
	for i, v := range partTwoSortedList {
		rankVal := i + 1
		sumMoney := v.GetBid() * rankVal
		partTwoTotalBid += sumMoney
		fmt.Printf("Rank %d:\n", rankVal)
		v.PrintHand()
		fmt.Printf("\n")
	}
	fmt.Printf("The value of the new games are: %d\n", partTwoTotalBid)
	// var blah [2]hand.Hand = hand.CompareHands(gameList[2], gameList[3])
	// for _, v := range blah {
	// 	v.PrintHand()
	// }

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
