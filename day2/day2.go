package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type GameSet struct {
	red, blue, green int
}

type Game struct {
	gameID int
	g      []*GameSet
}

func NewGameSet(red, blue, green int) GameSet {
	g := GameSet{}
	g.red = red
	g.blue = blue
	g.green = green
	return g
}

func NewGame(gameID int) Game {
	// returns a new game object
	g := Game{}
	g.gameID = gameID
	return g
}

func (g *Game) pushGameSet(gs *GameSet) {
	g.g = append(g.g, gs)
}

func IsGamePossible(red, blue, green int, fullGame Game) bool {
	// This function, given the numbers supplied, will check to see
	// if the game is possible to play.
	for _, g := range fullGame.g {
		if g.red > red {
			return false
		}
		if g.blue > blue {
			return false
		}
		if g.green > green {
			return false
		}
	}
	return true
}

func GetPowerOfMinimum(fullGame Game) int {
	// this will get the power of the _minimum_ amount of cubes needed
	// to complete this game.
	var minRed, minGreen, minBlue int = 0, 0, 0
	for _, g := range fullGame.g {
		if minRed < g.red {
			minRed = g.red
		}
		if minGreen < g.green {
			minGreen = g.green
		}
		if minBlue < g.blue {
			minBlue = g.blue
		}
	}
	return minRed * minGreen * minBlue
}

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

	Games := []Game{}
	// Time to parse! Here we goooooo
	for _, line := range lines {
		// First, find the Game number
		gnumline := strings.Split(line, ":")
		GameNum, _ := strconv.Atoi(strings.Split(gnumline[0], " ")[1])
		thisGame := NewGame(GameNum)

		// Now let's split each set
		sets := strings.Split(gnumline[1], ";")
		for _, v := range sets { // for each set per game
			working := strings.Trim(v, " ") // trim the fat
			subsets := strings.Split(working, ",")

			var red, green, blue int = 0, 0, 0
			for _, v1 := range subsets { // for each set of cubes
				choice := strings.Split(strings.Trim(v1, " "), " ")
				choiceNum, _ := strconv.Atoi(choice[0])
				switch choice[1] {
				case "red":
					red = choiceNum
				case "green":
					green = choiceNum
				case "blue":
					blue = choiceNum
				}
			}
			gs := NewGameSet(red, blue, green)
			thisGame.pushGameSet(&gs)
		}
		Games = append(Games, thisGame)
	}

	// parameters: total red: 12, total green: 13, total blue: 14
	var total int = 0
	for _, v := range Games {
		if IsGamePossible(12, 14, 13, v) {
			total += v.gameID
		}
	}
	fmt.Printf("Total possible games: %d\n", total)

	fmt.Printf("Now the powers...\n")

	var totalPart2 int = 0
	for _, v := range Games {
		totalPart2 += GetPowerOfMinimum(v)
	}
	fmt.Printf("The total of the powers: %d\n", totalPart2)

	// fmt.Printf("My games: %+v\n", Games)
	// fmt.Printf("Game 4, first set blue: %d, second set green: %d, third set red: %d\n", Games[3].g[2].blue, Games[3].g[2].green, Games[3].g[2].red)
	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
