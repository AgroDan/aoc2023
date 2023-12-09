package main

import (
	"bufio"
	"day8/network"
	"day8/puzzleparser"
	"flag"
	"fmt"
	"os"
	"time"
)

// Helpers to find the lowest common multiple

// Greatest Common Denominator
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least Common Multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func getMostCommon(arr []int) int {
	// returns the most common item from a slice of ints
	m := make(map[int]int)
	for _, i := range arr {
		m[i]++
	}
	highest := 0
	highVal := 0
	for k, v := range m {
		if v > highest {
			highest = v
			highVal = k
		}
	}
	return highVal
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

	// Insert code here
	m := puzzleparser.Parse(lines)
	m.PrintMap()

	fmt.Printf("Now let's traverse the map starting at AAA...\n")
	ptr := m.Map["AAA"]
	var steps int = 0
	for {
		if ptr.Name == "ZZZ" {
			break
		}
		switch m.NextDirection() {
		case 'L':
			ptr = ptr.Left
			steps++
		case 'R':
			ptr = ptr.Right
			steps++
		}
	}
	fmt.Printf("Reached ZZZ in %d steps!\n", steps)

	/*
	 * At this point, we are now working concurrently. Sort of. I could
	 * attempt to do this with concurrent threads but it's really not
	 * necessary when I can just step through one at a time.
	 */

	fmt.Printf("Now let's do some calculations to determine how several travelers can have a common multiple.\n")

	var ghostPtr []*network.Node
	for k := range m.Map {
		if k[len(k)-1] == 'A' {
			ghostPtr = append(ghostPtr, m.Map[k])
		}
	}

	// The ghostPtr pointers will tell us the valid starting points.
	// Lets create separate maps for each as a []Traveller object slice

	var tr []network.Traveller

	for _, v := range ghostPtr {
		tr = append(tr, network.Traveller{
			Ptr: v,
			FM:  m.DeepCopy(),
		})
	}

	var deltaStack []int
	for _, thisTrav := range tr {
		fmt.Printf("Traveller %s: ", thisTrav.Ptr.Name)
		var recSteps int = 0
		var stepsToEnd []int
		for i := 0; i < 100_000; i++ {
			if thisTrav.Ptr.Name[len(thisTrav.Ptr.Name)-1] == 'Z' {
				// hit an end, record steps taken
				stepsToEnd = append(stepsToEnd, recSteps)
			}

			recSteps++
			switch thisTrav.FM.NextDirection() {
			case 'R':
				thisTrav.Ptr = thisTrav.Ptr.Right
			case 'L':
				thisTrav.Ptr = thisTrav.Ptr.Left
			}
		}

		for _, v := range stepsToEnd {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("-=- Deltas: ")
		var deezDeltas []int
		for i := 1; i < len(stepsToEnd); i++ {
			fmt.Printf("%d ", stepsToEnd[i]-stepsToEnd[i-1])
			deezDeltas = append(deezDeltas, stepsToEnd[i]-stepsToEnd[i-1])
		}
		deltaStack = append(deltaStack, getMostCommon(deezDeltas))
		fmt.Printf("\n")
	}

	fmt.Printf("We have a list of most common deltas between each traveler: %+v\n", deltaStack)
	fmt.Printf("Now find the lowest common multiple to get the total steps needed!\n")

	// Now find the lowest comment multiple.
	if len(deltaStack) > 2 {
		lcmResult := lcm(deltaStack[0], deltaStack[1])
		for i := 2; i < len(deltaStack); i++ {
			lcmResult = lcm(lcmResult, deltaStack[i])
		}
		fmt.Printf("It will take %d steps to all reach the end.\n", lcmResult)
	} else if len(deltaStack) == 2 {
		lcmResult := lcm(deltaStack[0], deltaStack[1])
		fmt.Printf("It will take %d steps to all reach the end.\n", lcmResult)
	} else {
		lcmResult := deltaStack[0]
		fmt.Printf("It will take %d steps to all reach the end.\n", lcmResult)
	}

	/*
	 * This next section is when I tried to find commonalities between all of
	 * the travellers. It's here that I discovered the pattern by looking at
	 * the deltas of all the times a traveller would hit an endpoint.
	 */

	// potEnd := make(map[string]*network.Node)
	// endTracks := make(map[string][]int)
	// var recSteps int = 0
	// for i := 0; i < 100_000; i++ {
	// 	if tr[0].Ptr.Name[len(tr[0].Ptr.Name)-1] == 'Z' {
	// 		if _, ok := potEnd[tr[0].Ptr.Name]; ok {
	// 			// Been here before, record this
	// 			potEnd[tr[0].Ptr.Name] = tr[0].Ptr
	// 			endTracks[tr[0].Ptr.Name] = append(endTracks[tr[0].Ptr.Name], recSteps)
	// 		} else {
	// 			var s []int
	// 			s = append(s, recSteps)
	// 			endTracks[tr[0].Ptr.Name] = s
	// 			potEnd[tr[0].Ptr.Name] = tr[0].Ptr
	// 		}
	// 	}

	// 	recSteps++
	// 	switch tr[0].FM.NextDirection() {
	// 	case 'L':
	// 		tr[0].Ptr = tr[0].Ptr.Left
	// 	case 'R':
	// 		tr[0].Ptr = tr[0].Ptr.Right
	// 	}
	// }

	// fmt.Printf("endTracks: %+v\n", endTracks)
	// fmt.Printf("potEnd: %+v\n", potEnd)

	// // now let's get the deltas
	// var deltas []int
	// for k := range endTracks {
	// 	fmt.Printf("Deltas for %s: ", k)
	// 	for i := 0; i < len(endTracks[k]); i++ {
	// 		if i+1 == len(endTracks[k]) {
	// 			break
	// 		}
	// 		deltas = append(deltas, endTracks[k][i+1]-endTracks[k][i])
	// 		fmt.Printf("%d, ", endTracks[k][i+1]-endTracks[k][i])
	// 	}
	// 	fmt.Printf("\n")
	// }

	/*
	 * Below here is my first try. This is when I made the disheartening discovery
	 * that this challenge will be one of those "Use math to find the number rather
	 * than brute-forcing your way to the solution yourself." The below would work...
	 * eventually. But it will take just a tremendous amount of time to compute.
	 * I don't have the heart to delete this so here it is in comment form.
	 */

	// fmt.Printf("Found %d ghosts to follow!\n", len(tr))
	// fmt.Printf("Names of: ")
	// for _, v := range tr {
	// 	fmt.Printf("%s ", v.Ptr.Name)
	// }
	// fmt.Printf("\n")
	// var partTwoSteps int = 0
	// for {
	// 	// First, are we at an end?
	// 	end := true
	// 	for i := 0; i < len(tr); i++ {
	// 		// fmt.Printf("Name is: %s\n", tr[i].Ptr.Name)
	// 		if !end {
	// 			break
	// 		}
	// 		if tr[i].Ptr.Name[len(tr[i].Ptr.Name)-1] != 'Z' {
	// 			end = false
	// 		}
	// 	}
	// 	if end {
	// 		// We reached the end!
	// 		break
	// 	}

	// 	partTwoSteps++
	// 	// fmt.Printf("%d steps so far...\n", partTwoSteps)
	// 	for i := 0; i < len(tr); i++ {
	// 		// Now we progress each one forward once
	// 		thisDir := tr[i].FM.NextDirection()
	// 		switch thisDir {
	// 		case 'L':
	// 			tr[i].Ptr = tr[i].Ptr.Left
	// 		case 'R':
	// 			tr[i].Ptr = tr[i].Ptr.Right
	// 		}
	// 	}
	// }
	// fmt.Printf("Found all possible directions, took %d steps!\n", partTwoSteps)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
