package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	ID            int
	Winning, Game []int
}

func NewGame(GameNum int, w, g []int) Card {
	// This will return a new game
	return Card{ID: GameNum, Winning: w, Game: g}
}

func PrintGame(g Card) {
	fmt.Printf("Game %3d: ", g.ID)
	for _, n := range g.Winning {
		fmt.Printf("%2d ", n)
	}
	fmt.Printf("| ")
	for _, n := range g.Game {
		fmt.Printf("%2d ", n)
	}
	fmt.Printf("\n")
}

func (c Card) GetScore() int {
	// This function will loop through each Winning number
	// and look for matches in the Game array. If it finds
	// one, it adds one to the score, then returns a score
	// based on the algorithm, which is starting at 1 and
	// doubled for each match

	var winCount int = 0
	for _, v := range c.Winning {
		for _, thisGame := range c.Game {
			if v == thisGame {
				winCount++
			}
		}
	}

	var score int = 0
	for i := 0; i < winCount; i++ {
		if score == 0 {
			score++
		} else {
			score *= 2
		}
	}
	return score
}

func (c Card) GetCopies() []int {
	// This will return a slice of numbers that should be
	// game IDs that should have copies. Hard to explain
	// without knowing what the challenge wants (for part 2)
	var winCount int = 0
	for _, v := range c.Winning {
		for _, thisGame := range c.Game {
			if v == thisGame {
				winCount++
			}
		}
	}

	var retVal []int
	for i := 0; i < winCount; i++ {
		retVal = append(retVal, c.ID+(i+1))
	}

	// fmt.Printf("For Game number %d, we have %d winners.\n", c.ID, len(retVal))
	return retVal
}

func TruncateIfNeeded(gameList []Card, copyList []int) []int {
	// This is also hard to explain. Takes the output of the
	// getCopies and determines if any of the numbers listed
	// exceed the amount of cards we have. if so, remove them.
	highestCard := gameList[len(gameList)-1].ID

	var retVal []int
	for _, v := range copyList {
		if v <= highestCard {
			retVal = append(retVal, v)
		}
	}
	// fmt.Printf("%+v\n", retVal)
	return retVal
}

func ParseGame(line string) Card {
	// This will parse a line and return a game
	// a line will look like this:
	//
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	splitGame := strings.Split(line, ":")

	// First, let's get the game number
	gameMeta := strings.Split(strings.Trim(splitGame[0], " "), " ")
	gameNum, _ := strconv.Atoi(gameMeta[len(gameMeta)-1])

	// Now let's get the Winning games
	GameData := strings.Split(splitGame[1], "|")

	winners := strings.Split(strings.Trim(GameData[0], " "), " ")
	checkScores := strings.Split(strings.Trim(GameData[1], " "), " ")

	var wNum []int
	var cNum []int

	for _, w := range winners {
		n, err := strconv.Atoi(w)
		// ignore errors, but print anyway
		if err != nil {
			// toss out the junk data
			// fmt.Printf("Error!: %s\n", err)
			continue
		}
		wNum = append(wNum, n)
	}

	for _, c := range checkScores {
		n, err := strconv.Atoi(c)
		if err != nil {
			// toss out the junk data
			// fmt.Printf("Error!: %s\n", err)
			continue
		}
		cNum = append(cNum, n)
	}

	return NewGame(gameNum, wNum, cNum)
}

// Now let's create an unordered Set
type CardGames struct {
	ID int
}

// Use this to generate a set with copies
// mySet := make(map[CardGames]int)

// Then you can initialize it with
// mySet[CardGames{ID: 50}] = 1 // or whatever
