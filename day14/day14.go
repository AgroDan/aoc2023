package main

import (
	"bufio"
	"day14/patterns"
	"day14/rocks"
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
	thisMap := rocks.NewRockMap(lines)
	thisMap.PrintRockMap()

	fmt.Printf("Now let's move all the round rocks up ten times...\n")
	for {
		if !thisMap.MoveRocks(rocks.North) {
			fmt.Printf("No movement, breaking!\n")
			break
		}
	}
	// thisMap.MoveRocks(rocks.North)

	thisMap.PrintRockMap()
	fmt.Printf("Total load bearing: %d\n", thisMap.GetTotalLoad())

	// // Lets try to compare rockmaps
	// var myRocks []rocks.RockMap
	// testMap := rocks.NewRockMap(lines)
	// fmt.Printf("Move everything North\n")
	// for {
	// 	if !testMap.MoveRocks(rocks.North) {
	// 		break
	// 	}
	// }
	// fmt.Printf("Deep copy...\n")
	// myRocks = append(myRocks, testMap.DeepCopy())

	// fmt.Printf("Now move everything south.\n")
	// for {
	// 	if !testMap.MoveRocks(rocks.South) {
	// 		break
	// 	}
	// }
	// fmt.Printf("Deep copy...\n")
	// myRocks = append(myRocks, testMap.DeepCopy())

	// fmt.Printf("Now let's make another rockmap from scratch and move north and south.\n")
	// yetAnotherMap := rocks.NewRockMap(lines)
	// for {
	// 	if !yetAnotherMap.MoveRocks(rocks.North) {
	// 		break
	// 	}
	// }
	// hash1, err := hashstructure.Hash(yetAnotherMap, hashstructure.FormatV2, nil)
	// if err != nil {
	// 	panic("could not hash")
	// }
	// hash2, err2 := hashstructure.Hash(myRocks[0], hashstructure.FormatV2, nil)
	// if err2 != nil {
	// 	panic("could not get second hash")
	// }
	// fmt.Printf("Is this the same as idx 0 of the original? %t\n", hash1 == hash2)

	// for {
	// 	if !yetAnotherMap.MoveRocks(rocks.South) {
	// 		break
	// 	}
	// }

	fmt.Printf("Now let's get the next part for step 2, let's build a new map.\n")
	part2Map := rocks.NewRockMap(lines)
	part2Map.PrintRockMap()
	// fmt.Printf("Moving everything west\n")
	// for {
	// 	if !part2Map.MoveRocks(rocks.West) {
	// 		break
	// 	}
	// }
	fmt.Printf("Starting the washer cycle, north, west, south, east 4x\n")

	// var ITER int = 100000
	var ITER int = 1000
	var CHECK int = 305
	var WINDOW int = 10

	rpt := rocks.NewRecordsManager()
	// var stateArray []uint64
	// bCrumb := rocks.NewBreadCrumb()
	// cycleMetrics := rocks.NewLoadMarker()
	washer := [4]int{rocks.North, rocks.West, rocks.South, rocks.East}

	for i := 0; i < ITER; i++ {
		for j := 0; j < len(washer); j++ {
			for {
				if !part2Map.MoveRocks(washer[j]) {
					break
				}
			}
			// fmt.Printf("Finished direction %d...\n", washer[j])
		}
		fmt.Printf("\rCycle #%d...", i)
		rpt.Add(part2Map)
		// if bCrumb.Add(part2Map) {
		// 	fmt.Printf("\nCollision at cycle %d!\n", i)
		// 	fmt.Printf("Load value at cycle %d: %d\n", i, part2Map.GetTotalLoad())
		// }
		// hash, err := hashstructure.Hash(part2Map, hashstructure.FormatV2, nil)
		// if err != nil {
		// 	panic(err)
		// }
		// stateArray = append(stateArray, hash)
		// Now calculate each load value on each wash cycle
		// thisLoad := part2Map.GetTotalLoad()
		// fmt.Printf("thisLoad: %d\n", thisLoad)
		// cycleMetrics.Add(thisLoad)
		// fmt.Printf("Finished cycle %d...\n", i)
	}
	fmt.Printf("\n")
	part2Map.PrintRockMap()

	fmt.Printf("Now let's find all of the repeats...\n")
	foundPatterns := patterns.FindPatternV2(rpt, WINDOW, CHECK)
	var first bool = true
	var prev int = 0
	// fmt.Printf("WTF is this: %+v\n", foundPatterns)
	for _, v := range foundPatterns {
		fmt.Printf("Found pattern at index %d that matches for %d items after.\n", v, WINDOW)

		if first {
			first = false
		} else {
			fmt.Printf("Difference from previous: %d\n", v-prev)
			prev = v
		}

	}
	// I AM GOING TO JUST FIGURE THIS OUT MANUALLY BECAUSE IT'S FASTER
	// Basically, I used the above to find a pattern. I randomly chose the value of
	// iteration 767 because it repeated every 42 times.
	// Here is the math:
	//
	// interesting value: 767 with a repeat every 42 times
	// 1_000_000_000 - 767 ==  999,999,233
	// 999_999_233 / 42 == 23,809,505.54761905 <--- get the floor value of this, so 23_809_505
	// so 238_095_505 * 42 == 999_999_210
	// Add 767 to that and get: 999_999_977
	// Which is only 22 numbers away from one billion.

	// And that means I just need to find the value of LoadBearingVol[767 + 22]
	// ergo, the value of LoadBearingVol[789] is the answer! (remember starting at 0 is normal,
	// but because the AOC dev is a perl programmer everything is wrong and starts at 1)
	answer, _ := rpt.Get(789)
	fmt.Printf("The value of 789 and thus the answer to the puzzle: %d\n", answer)

	// fmt.Printf("Checking idx 501, this is the value: %d\n", thisval)
	// fmt.Printf("Add 9 to that to get the value at 1 billion: %d\n", newval)
	// for i := CHECK; i < ITER; i++ {
	// 	diff := rpt.FindNextMatch(i)
	// 	if diff < 0 {
	// 		// didn't find a match
	// 		continue
	// 	}
	// 	lb, _ := rpt.Get(i)
	// 	fmt.Printf("Pattern repeats at index %d after index %d with a load bearing volume of %d.\n", diff, i, lb)
	// 	fmt.Printf("which means that _potentially_ after %d iterations it will be the same.\n", diff-i)

	// }

	// 507 and 513 should be the same, right?

	// fmt.Printf("507 and 513 should be the same, right? Let's try.\n")
	// fmt.Printf("507: %+v\n", i507)
	// i507.PrintRockMap()
	// h, err1 := i507.Hash()
	// if err1 != nil {
	// 	panic(err1)
	// }
	// fmt.Printf("Hash of 507: %s\n", h)
	// fmt.Printf("513: %+v\n", i513)
	// i513.PrintRockMap()
	// h2, err2 := i513.Hash()
	// if err2 != nil {
	// 	panic(err2)
	// }
	// fmt.Printf("Hash of 513: %s\n", h2)

	// fmt.Printf("Now let's find some patterns...\n")
	// foundPatterns := patterns.FindPattern(stateArray, 10)
	// // sorted := sort.Ints()
	// fmt.Printf("Found repeats at indexes: %+v\n", foundPatterns)

	// fmt.Printf("Top values after %d cycles...\n", ITER)
	// cycleMetrics.PrintTopVals(5)
	// cycleMetrics.PrintVals()

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
