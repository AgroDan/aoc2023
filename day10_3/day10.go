package main

import (
	"bufio"
	"day10_3/pipes"
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

	// First build the map
	thisMap := pipes.NewPipeMap(lines)

	// Print the map
	thisMap.PrintMap()

	// Now we have a starting point, so let's set up a Queue
	var dirQueue pipes.Queue

	start := thisMap.GetStart()
	// fmt.Printf("Start X: %d, Start Y: %d\n", start.X, start.Y)
	dirQueue.Push(start)

	for {
		curPos, err := dirQueue.Pop()
		if err != nil {
			break
		}

		// Have we been here before?
		if thisMap.PipeCircuit.Contains(curPos) {
			continue
		} else { // if not, put it in the set
			thisMap.PipeCircuit.Add(curPos)
		}

		// Get the neighbors now
		neighbors := thisMap.GetValidNeighbors(curPos)

		for _, v := range neighbors {
			dirQueue.Push(v)
		}
	}

	// fmt.Printf("Known set: %+v\n", visited)
	fmt.Printf("Half of distance travelled: %d\n", thisMap.PipeCircuit.Len()/2)

	fmt.Printf("Cleaning the pipes...")
	thisMap.CleanMyPipes()
	fmt.Printf("Done!\n")
	thisMap.PrintMap()

	// pleaseWorkCoord := pipes.Coord{Y: 7, X: 0}

	// test, err := thisMap.CountLayers(pleaseWorkCoord, pipes.W)
	// if err != nil {
	// 	fmt.Printf("Uh oh!: %s\n", err)
	// }

	// fmt.Printf("Checked coordinate: %+v\n", pleaseWorkCoord)
	// fmt.Printf("Value at location: %c\n", thisMap.GetMapLocation(pleaseWorkCoord))
	// fmt.Printf("Layers found: %d\n", test)
	// fmt.Printf("If odd, would be internal. Is it odd? %t\n", pipes.IsOdd(test))

	// now loop through all the ground pieces and check to see if it's internal.
	// if using the CountLayers() function, if the number is odd, it should be
	// enclosed within the loop.
	finalAnswer := thisMap.GetInternal()
	fmt.Printf("Amount of internal pices: %d\n", finalAnswer)

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
