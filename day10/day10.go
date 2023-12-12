package main

import (
	"bufio"
	"day10/pipes"
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
	// I don't need no stinkin parer
	pMap := pipes.NewPipesMap(lines)
	pMap.PrintMap()
	player := pipes.NewPlayer(&pMap)
	player.StartWalking()
	fmt.Printf("We ran a full circuit in %d steps.\n", player.GetSteps())
	fmt.Printf("The answer to part 1 is: %d\n", player.GetSteps()/2)

	// Now by the time we got here, we should have obtained a breadcrumb list of
	// every pipe in the circuit. From here, I can loop through every coord in
	// that path and find all 4 directions from each point, determine which (if any)
	// of those directions is part of the actual pipe, and to the BFS work on it
	// to determine how much area could be a nest.

	// potentialNest := pipes.NewCircuit()
	// deezPipes := player.GetBreadCrumbs()
	// for _, v := range deezPipes.GetListOfPipes() {
	// 	// i know, shut up.
	// 	dTiles := [4]pipes.Coord{v.Peek(pipes.N), v.Peek(pipes.S), v.Peek(pipes.W), v.Peek(pipes.E)}
	// 	for _, vv := range dTiles {
	// 		if !deezPipes.Contains(vv) && !pMap.IsEdge(vv) && pMap.IsInBounds(vv) {
	// 			potentialNest.Add(vv)
	// 		}
	// 	}
	// }

	// fmt.Printf("Possible nest locations: %d\n", len(potentialNest.GetListOfPipes()))
	// // Now, loop on through and hold onto yer butts...
	// var total int = 0
	// for _, v := range potentialNest.GetListOfPipes() {
	// 	fmt.Printf("Checking %+v\n", v)
	// 	nest := pMap.DetectSpaces(v, deezPipes)
	// 	total += len(nest)
	// }
	// fmt.Printf("Part 2 -=- Total tiles enclosed in circuit: %d\n", total)

	fmt.Printf("Map with pipes blocked out:\n")
	pMap.PrintVisitedMap(player.GetBreadCrumbs())
	potentialNests := pipes.GetNotInCircuit(pMap, player.GetBreadCrumbs())
	fmt.Printf("Amount of tiles not in the puzzle: %d\n", len(potentialNests))

	// Now let's DFS this crap
	var total int = 0
	var totNests []pipes.Coord
	for _, v := range potentialNests {
		nest := pMap.DetectSpaces(v, player.GetBreadCrumbs())
		totNests = append(totNests, nest...)
		total += len(nest)
	}
	fmt.Printf("Part 2: %d\n", total)
	pMap.PrintMarkedSpaces(totNests, player.GetBreadCrumbs())

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
